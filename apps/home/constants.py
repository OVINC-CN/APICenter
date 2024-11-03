from dataclasses import dataclass
from typing import Tuple

# pylint: disable=C0103
CONFIG_DEFAULT_T = Tuple[str, str]


# pylint: disable=C0103,R0902
@dataclass
class BuildInKeys:
    """
    Build in config keys
    """

    LOGO_URL: CONFIG_DEFAULT_T = "logo_url", "/favicon.ico"
    WEBSITE_TITLE: CONFIG_DEFAULT_T = "website_title", ""
    BRAND_NAME: CONFIG_DEFAULT_T = "brand_name", "OVINC"
    BRAND_TITLE: CONFIG_DEFAULT_T = "brand_title", "OVINC CN"
    BRAND_DESC: CONFIG_DEFAULT_T = "brand_desc", ""
    BRAND_VISION: CONFIG_DEFAULT_T = "brand_vision", ""
    CONTACT_PICTURE_URL: CONFIG_DEFAULT_T = "contact_picture_url", ""
    CONTACT_MAIL: CONFIG_DEFAULT_T = "contact_email", ""
    CONTACT_PHONE: CONFIG_DEFAULT_T = "contact_phone", ""
    CONTACT_PLACE: CONFIG_DEFAULT_T = "contact_place", ""
    COPYRIGHT: CONFIG_DEFAULT_T = "copyright", "OVINC CN"
    MIIT_FILLING_CODE: CONFIG_DEFAULT_T = "miit_filling_code", ""
    MIIT_FILLING_URL: CONFIG_DEFAULT_T = "miit_filling_url", ""
    GONGAN_FILLING_ID: CONFIG_DEFAULT_T = "gongan_filling_id", ""
    GONGAN_FILLING_URL: CONFIG_DEFAULT_T = "gongan_filling_url", ""
    USER_AGREEMENT: CONFIG_DEFAULT_T = "user_agreement", ""
    PRIVACY_AGREEMENT: CONFIG_DEFAULT_T = "privacy_agreement", ""
    BACKGROUND_IMAGE: CONFIG_DEFAULT_T = "background_image", ""
