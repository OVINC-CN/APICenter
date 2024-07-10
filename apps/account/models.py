import abc
import hashlib
from typing import Union

from django.conf import settings
from django.contrib.auth.base_user import AbstractBaseUser
from django.contrib.auth.hashers import make_password
from django.contrib.auth.models import AbstractUser, AnonymousUser, PermissionsMixin
from django.contrib.auth.models import UserManager as _UserManager
from django.core.cache import cache
from django.db import models
from django.utils.translation import gettext, gettext_lazy
from django_redis.client import DefaultClient
from ovinc_client.core.constants import MAX_CHAR_LENGTH, SHORT_CHAR_LENGTH
from ovinc_client.core.models import SoftDeletedManager, SoftDeletedModel
from ovinc_client.core.utils import num_code, uniq_id

from apps.account.constants import (
    PHONE_VERIFY_CODE_KEY,
    PHONE_VERIFY_CODE_LENGTH,
    PHONE_VERIFY_CODE_TIMEOUT,
    UserTypeChoices,
)
from apps.cel.tasks import send_notice
from apps.notice.constants import NoticeWayChoices

cache: DefaultClient


class UserManager(SoftDeletedManager, _UserManager):
    """
    User Manager
    """

    # pylint: disable=W0237
    def create_user(self, username, nick_name=None, password=None, **extra_fields):
        if not username:
            raise ValueError(gettext("Username Cannot be Empty"))
        user = self.model(username=username, nick_name=nick_name, **extra_fields)
        user.password = make_password(password)
        user.save(using=self._db)
        return user

    # pylint: disable=W0237
    def create_superuser(self, username, nick_name=None, password=None, **extra_fields):
        extra_fields["is_superuser"] = True
        self.create_user(username, nick_name, password, **extra_fields)


class User(SoftDeletedModel, AbstractBaseUser, PermissionsMixin):
    """
    User
    """

    username = models.CharField(
        gettext_lazy("username"),
        max_length=SHORT_CHAR_LENGTH,
        primary_key=True,
        validators=[AbstractUser.username_validator],
        error_messages={"unique": gettext_lazy("already in use")},
    )
    nick_name = models.CharField(gettext_lazy("Nick Name"), max_length=SHORT_CHAR_LENGTH, blank=True, null=True)
    user_type = models.CharField(
        gettext_lazy("User Type"),
        choices=UserTypeChoices.choices,
        max_length=SHORT_CHAR_LENGTH,
        default=UserTypeChoices.PERSONAL.value,
    )
    date_joined = models.DateTimeField(gettext_lazy("Date Joined"), auto_now_add=True)
    phone_number = models.CharField(
        gettext_lazy("Phone Number"), max_length=SHORT_CHAR_LENGTH, null=True, blank=True, unique=True
    )
    email_address = models.EmailField(gettext_lazy("Email Address"), null=True, blank=True, unique=True)
    wechat_open_id = models.CharField(
        gettext_lazy("Wechat Open ID"), max_length=MAX_CHAR_LENGTH, null=True, blank=True, unique=True
    )
    wechat_union_id = models.CharField(
        gettext_lazy("Wechat Union ID"), max_length=MAX_CHAR_LENGTH, null=True, blank=True, unique=True
    )
    avatar = models.TextField(gettext_lazy("Avatar"), null=True, blank=True)
    is_staff = models.BooleanField(gettext_lazy("Is Staff"), default=False)

    USERNAME_FIELD = "username"
    objects = UserManager()
    _objects = _UserManager()

    class Meta:
        verbose_name = gettext_lazy("User")
        verbose_name_plural = verbose_name
        ordering = ["username"]

    def generate_oauth_code(self) -> str:
        """
        Generate OAuth User Code
        """

        code = uniq_id()
        cache.set(code, self.username, timeout=settings.OAUTH_CODE_TIMEOUT)
        return code

    @classmethod
    def check_oauth_code(cls, code: str) -> (bool, Union[models.Model, None]):
        """
        Check OAuth User Code
        """

        username = cache.get(code)
        cache.delete(code)
        try:
            return True, cls.objects.get(username=username)
        except cls.DoesNotExist:  # pylint: disable=E1101
            return False, None

    def reset_password(self, password: str, is_raw: bool = False) -> None:
        """
        Reset User Password
        """

        if is_raw:
            # pylint: disable=E1101
            password = hashlib.sha256(f"{password}{self.username.lower()}".encode()).hexdigest()

        self.set_password(password)
        self.save(update_fields=["password"])

    @classmethod
    def send_phone_verify_code(cls, area: str, phone_number: str) -> None:
        """
        Send Phone Verification Code
        """

        phone_number = f"{area}{phone_number}"
        key = PHONE_VERIFY_CODE_KEY.format(phone_number=phone_number)
        code = num_code(PHONE_VERIFY_CODE_LENGTH)
        cache.set(key=key, value=code, timeout=PHONE_VERIFY_CODE_TIMEOUT)
        send_notice.delay(
            notice_type=NoticeWayChoices.MSG,
            usernames=[],
            receivers=[phone_number],
            content={"tid": settings.NOTICE_SMS_ID_VERIFY_CODE, "params": [code, str(PHONE_VERIFY_CODE_TIMEOUT // 60)]},
        )

    @classmethod
    def check_phone_verify_code(cls, area: str, phone_number: str, code: str) -> bool:
        """
        Verify Phone Number
        """

        phone_number = f"{area}{phone_number}"
        key = PHONE_VERIFY_CODE_KEY.format(phone_number=phone_number)
        val = cache.get(key=key)
        result = str(val) == str(code)
        if result:
            cache.delete(key=key)
        return result

    @classmethod
    def load_user_by_union_id(cls, union_id: str) -> "User":
        return cls.objects.filter(wechat_union_id=union_id).first()


class CustomAnonymousUser(AnonymousUser, abc.ABC):
    """
    Anonymous User
    """

    nick_name = "AnonymousUser"
    user_type = UserTypeChoices.PLATFORM.value