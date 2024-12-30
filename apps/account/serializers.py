import re

from adrf.serializers import ModelSerializer, Serializer
from django.contrib.auth import get_user_model
from django.utils.translation import gettext, gettext_lazy
from ovinc_client.core.async_tools import SyncRunner
from ovinc_client.tcaptcha.exceptions import TCaptchaInvalid
from ovinc_client.tcaptcha.utils import TCaptchaVerify
from rest_framework import serializers

from apps.account.constants import (
    PHONE_VERIFY_CODE_LENGTH,
    USERNAME_REGEX,
    PhoneNumberAreas,
)
from apps.account.exceptions import PhoneVerifyCodeInvalid
from apps.account.models import User
from apps.home.constants import BuildInKeys
from apps.home.models import MetaConfig

USER_MODEL: User = get_user_model()


class UserInfoSerializer(ModelSerializer):
    """
    User Info
    """

    class Meta:
        model = USER_MODEL
        fields = ["username", "nick_name", "last_login", "user_type"]


class SignInSerializer(Serializer):
    """
    Sign in
    """

    username = serializers.CharField(label=gettext_lazy("Username"))
    password = serializers.CharField(label=gettext_lazy("Password"))
    is_oauth = serializers.BooleanField(label=gettext_lazy("Is OAuth"), default=False)
    wechat_code = serializers.CharField(label=gettext_lazy("WeChat Code"), required=False)
    tcaptcha = serializers.JSONField(label=gettext_lazy("Tencent Captcha"), required=False, default=dict)

    def validate(self, attrs: dict) -> dict:
        data = super().validate(attrs)
        if not SyncRunner().run(
            TCaptchaVerify(user_ip=self.context.get("user_ip", ""), **data.get("tcaptcha", {})).verify()
        ):
            raise TCaptchaInvalid()
        return data


class UserRegistrySerializer(ModelSerializer):
    """
    User Registry
    """

    username = serializers.RegexField(label=gettext_lazy("Username"), regex=USERNAME_REGEX)
    is_oauth = serializers.BooleanField(label=gettext_lazy("Is OAuth"), default=False)
    wechat_code = serializers.CharField(label=gettext_lazy("WeChat Code"), required=False)
    phone_area = serializers.ChoiceField(label=gettext_lazy("Phone Area"), choices=PhoneNumberAreas.choices)
    phone_number = serializers.CharField(label=gettext_lazy("Phone Number"))
    phone_verify = serializers.CharField(
        label=gettext_lazy("Phone Verify Code"),
        min_length=PHONE_VERIFY_CODE_LENGTH,
        max_length=PHONE_VERIFY_CODE_LENGTH,
    )
    tcaptcha = serializers.JSONField(label=gettext_lazy("Tencent Captcha"), required=False, default=dict)

    class Meta:
        model = USER_MODEL
        fields = [
            "username",
            "nick_name",
            "password",
            "is_oauth",
            "wechat_code",
            "phone_area",
            "phone_number",
            "phone_verify",
            "tcaptcha",
        ]

    def validate(self, attrs: dict) -> dict:
        data = super().validate(attrs)
        if not SyncRunner().run(
            TCaptchaVerify(user_ip=self.context.get("user_ip", ""), **data.get("tcaptcha", {})).verify()
        ):
            raise TCaptchaInvalid()
        if not USER_MODEL.check_phone_verify_code(
            area=data["phone_area"], phone_number=data["phone_number"], code=data["phone_verify"]
        ):
            raise PhoneVerifyCodeInvalid()
        return data

    def validate_username(self, username: str) -> str:
        username_extra_regex = MetaConfig.objects.filter(key=BuildInKeys.USERNAME_EXTRA_REGEX[0]).first()
        if not username_extra_regex:
            return username
        if re.compile(username_extra_regex.val).match(username):
            return username
        raise serializers.ValidationError(gettext("Username Invalid"))

    def validate_phone_number(self, phone_number: str) -> str:
        if USER_MODEL.objects.filter(phone_number=phone_number).exists():
            raise serializers.ValidationError(gettext("Phone Number Already Exists"))
        return phone_number


class VerifyCodeRequestSerializer(Serializer):
    """
    Verify Code
    """

    code = serializers.CharField(label=gettext_lazy("Code"))


class VerifyTicketRequestSerializer(Serializer):
    """
    Verify Ticket
    """

    ticket = serializers.CharField(label=gettext_lazy("Ticket"))


class WeChatLoginReqSerializer(Serializer):
    """
    WeChat Login
    """

    code = serializers.CharField(label=gettext_lazy("Code"))
    state = serializers.CharField(label=gettext_lazy("State"))
    is_oauth = serializers.BooleanField(label=gettext_lazy("Is OAuth"), default=False)


class ResetPasswordRequestSerializer(Serializer):
    """
    Reset Password
    """

    username = serializers.RegexField(label=gettext_lazy("Username"), regex=USERNAME_REGEX)
    password = serializers.CharField(label=gettext_lazy("Password"), max_length=128)
    phone_area = serializers.ChoiceField(label=gettext_lazy("Phone Area"), choices=PhoneNumberAreas.choices)
    phone_number = serializers.CharField(label=gettext_lazy("Phone Number"))
    phone_verify = serializers.CharField(
        label=gettext_lazy("Phone Verify Code"),
        min_length=PHONE_VERIFY_CODE_LENGTH,
        max_length=PHONE_VERIFY_CODE_LENGTH,
    )
    tcaptcha = serializers.JSONField(label=gettext_lazy("Tencent Captcha"), required=False, default=dict)

    def validate(self, attrs: dict) -> dict:
        data = super().validate(attrs)
        if not SyncRunner().run(
            TCaptchaVerify(user_ip=self.context.get("user_ip", ""), **data.get("tcaptcha", {})).verify()
        ):
            raise TCaptchaInvalid()
        if not USER_MODEL.check_phone_verify_code(
            area=data["phone_area"], phone_number=data["phone_number"], code=data["phone_verify"]
        ):
            raise PhoneVerifyCodeInvalid()
        return data


class SendVerifyCodeRequestSerializer(Serializer):
    """
    Verify Code
    """

    phone_area = serializers.ChoiceField(label=gettext_lazy("Phone Area"), choices=PhoneNumberAreas.choices)
    phone_number = serializers.CharField(label=gettext_lazy("Phone Number"))
    tcaptcha = serializers.JSONField(label=gettext_lazy("Tencent Captcha"), required=False, default=dict)

    def validate(self, attrs: dict) -> dict:
        data = super().validate(attrs)
        if not SyncRunner().run(
            TCaptchaVerify(user_ip=self.context.get("user_ip", ""), **data.get("tcaptcha", {})).verify()
        ):
            raise TCaptchaInvalid()
        return data
