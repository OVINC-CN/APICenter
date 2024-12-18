from channels.db import database_sync_to_async
from django.contrib.auth import get_user_model
from django.utils.translation import gettext
from rest_framework.authentication import BaseAuthentication
from rest_framework.request import Request

from apps.application.models import Application
from core.constants import (
    APP_AUTH_HEADER_APPID_KEY,
    APP_AUTH_HEADER_APPID_NONCE,
    APP_AUTH_HEADER_APPID_SIGN,
    APP_AUTH_HEADER_APPID_TIMESTAMP,
)
from core.exceptions import AppAuthFailed

USER_MODEL = get_user_model()


class ApplicationAuthenticate(BaseAuthentication):
    """
    Application Authenticate
    """

    # pylint: disable=W0236
    async def authenticate(self, request: Request) -> (Application, None):
        # load params
        app_code = request.headers.get(APP_AUTH_HEADER_APPID_KEY)
        signature = request.headers.get(APP_AUTH_HEADER_APPID_SIGN)
        timestamp = request.headers.get(APP_AUTH_HEADER_APPID_TIMESTAMP)
        nonce = request.headers.get(APP_AUTH_HEADER_APPID_NONCE)
        if not app_code or not signature or not timestamp or not nonce:
            raise AppAuthFailed(gettext("App Auth Headers Not Exist"))
        # varify app
        try:
            app = await database_sync_to_async(Application.objects.get)(pk=app_code)
        except Application.DoesNotExist as err:  # pylint: disable=E1101
            raise AppAuthFailed(gettext("App Not Exist")) from err
        # verify secret
        if app.check_sign(signature=signature, timestamp=timestamp, nonce=nonce):
            return app, None
        raise AppAuthFailed(gettext("Signature Invalid"))
