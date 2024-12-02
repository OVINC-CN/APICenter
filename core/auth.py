import json

from channels.db import database_sync_to_async
from django.contrib.auth import get_user_model
from django.utils.translation import gettext
from ovinc_client.constants import APP_AUTH_ID_KEY, APP_AUTH_SECRET_KEY
from rest_framework.authentication import BaseAuthentication

from apps.application.models import Application
from core.constants import APP_AUTH_HEADER_KEY
from core.exceptions import AppAuthFailed
from core.threadpool import db_executor

USER_MODEL = get_user_model()


class ApplicationAuthenticate(BaseAuthentication):
    """
    Application Authenticate
    """

    # pylint: disable=W0236
    async def authenticate(self, request) -> (Application, None):
        # get from query
        if APP_AUTH_ID_KEY in request.query_params:
            app_params = request.query_params
        # get app params
        else:
            app_params = json.loads(request.META.get(APP_AUTH_HEADER_KEY, "{}"))
        if not isinstance(app_params, dict):
            raise AppAuthFailed(gettext("App Auth Params is not Json"))
        app_code = app_params.get(APP_AUTH_ID_KEY)
        app_secret = app_params.get(APP_AUTH_SECRET_KEY)
        if not app_code or not app_secret:
            raise AppAuthFailed(gettext("App Auth Params Not Exist"))
        # varify app
        try:
            app = await database_sync_to_async(Application.objects.get, executor=db_executor)(pk=app_code)
        except Application.DoesNotExist as err:  # pylint: disable=E1101
            raise AppAuthFailed(gettext("App Not Exist")) from err
        # verify secret
        if app.check_secret(app_secret):
            return app, None
        raise AppAuthFailed(gettext("App Code or Secret Incorrect"))
