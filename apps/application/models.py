import hashlib
import secrets
import time
from typing import List

from django.conf import settings
from django.core.cache import cache
from django.db import models, transaction
from django.utils.translation import gettext_lazy
from django_redis.client import DefaultClient
from ovinc_client.core.constants import MAX_CHAR_LENGTH, SHORT_CHAR_LENGTH
from ovinc_client.core.models import (
    BaseModel,
    ForeignKey,
    SoftDeletedManager,
    SoftDeletedModel,
    UniqIDField,
)

from apps.account.models import User
from core.constants import APP_AUTH_NONCE_CACHE_KEY

cache: DefaultClient


class ApplicationObjects(SoftDeletedManager):
    """
    Application Manager Objects
    """

    @transaction.atomic()
    def create(self, app_name: str, app_code: str, app_secret: str, managers: List[User], **kwargs) -> "Application":
        """
        Create Application
        """

        # create application
        application = super().create(app_name=app_name, app_code=app_code, **kwargs)

        # set secret
        application.set_secret(app_secret)

        # set managers
        app_managers = [ApplicationManager(application=application, manager=manager) for manager in managers]
        ApplicationManager.objects.bulk_create(app_managers)

        return application


class Application(SoftDeletedModel):
    """
    Application
    """

    app_name = models.CharField(gettext_lazy("App Name"), max_length=SHORT_CHAR_LENGTH)
    app_code = models.CharField(gettext_lazy("App Code"), max_length=SHORT_CHAR_LENGTH, primary_key=True)
    app_secret = models.CharField(gettext_lazy("App Secret"), max_length=MAX_CHAR_LENGTH, blank=True, null=True)
    app_url = models.URLField(gettext_lazy("App Url"), null=True, blank=True)
    app_logo = models.URLField(gettext_lazy("App Logo"), null=True, blank=True)
    app_desc = models.TextField(gettext_lazy("App Desc"), null=True, blank=True)
    is_hidden = models.BooleanField(gettext_lazy("Is Hidden"), default=False)

    objects = ApplicationObjects()

    class Meta:
        verbose_name = gettext_lazy("Application")
        verbose_name_plural = verbose_name
        ordering = ["app_code"]

    def check_sign(self, signature: str, timestamp: str, nonce: str) -> bool:
        """
        check signature
        """

        # check timestamp
        if not timestamp.isdigit() or int(timestamp) + settings.APP_AUTH_SIGN_EXPIRE < time.time():
            return False

        # check nonce deplicate
        if not cache.add(
            key=APP_AUTH_NONCE_CACHE_KEY.format(appid=self.app_code, nonce=nonce),
            value=signature,
            timeout=settings.APP_AUTH_SIGN_EXPIRE,
        ):
            return False

        # check signature
        raw_content = f"{timestamp}-{nonce}-{self.app_secret}"
        expect_signature = hashlib.sha256(raw_content.encode()).hexdigest()
        return secrets.compare_digest(signature, expect_signature)

    def set_secret(self, raw_secret: str) -> None:
        """
        set secret
        """

        self.app_secret = raw_secret
        self.save(update_fields=["app_secret"])


class ApplicationManager(BaseModel):
    """
    Application Manager
    """

    id = UniqIDField(gettext_lazy("ID"))
    application = ForeignKey(gettext_lazy("Application"), to="application.Application", on_delete=models.CASCADE)
    manager = ForeignKey(gettext_lazy("Manager"), to="account.User", on_delete=models.CASCADE)

    class Meta:
        verbose_name = gettext_lazy("Application Manager")
        verbose_name_plural = verbose_name
        unique_together = ["application", "manager"]
        ordering = ["id"]

    @classmethod
    def load_managers(cls, applications: List[Application]) -> List[User]:
        return ApplicationManager.objects.filter(application__in=applications).select_related("manager")

    @classmethod
    def load_manage_apps(cls, manager):
        return ApplicationManager.objects.filter(manager=manager).values("application")
