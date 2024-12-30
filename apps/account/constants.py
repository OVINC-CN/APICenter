from django.utils.translation import gettext_lazy
from ovinc_client.core.models import TextChoices

USERNAME_REGEX = r"[a-zA-Z]{1}[a-zA-Z0-9-_]{3,}"

PHONE_VERIFY_CODE_LENGTH = 6
PHONE_VERIFY_CODE_TIMEOUT = 60 * 60  # second
PHONE_VERIFY_CODE_KEY = "phone_verify_code:{phone_number}"

WECHAT_LOGIN_STATE_KEY = "wechat_login_state:{state}"
WECHAT_USER_INFO_KEY = "wechat_user_info:{code}"


class UserTypeChoices(TextChoices):
    """
    User Type Choices
    """

    PERSONAL = "personal", gettext_lazy("Personal")
    PLATFORM = "platform", gettext_lazy("Platform")


class WeChatAuthType(TextChoices):
    """
    WeChat Auth Type
    """

    CODE = "authorization_code", gettext_lazy("Code")


class PhoneNumberAreas(TextChoices):
    """
    Phone Number Areas
    """

    CHINA = "+86", gettext_lazy("Mainland China (+86)")
    HONG_KONG = "+852", gettext_lazy("Hong Kong, China (+852)")
    MACAO = "+853", gettext_lazy("Macao, China (+853)")
    TAIWAN = "+886", gettext_lazy("Taiwan, China (+886)")
