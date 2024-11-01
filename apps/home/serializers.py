from adrf.serializers import Serializer
from django.utils.translation import check_for_language, gettext_lazy
from rest_framework import serializers

from apps.home.exceptions import LanguageCodeInvalid


class I18nRequestSerializer(Serializer):
    """
    I18n
    """

    language = serializers.CharField(label=gettext_lazy("Language Code"))

    def validate_language(self, language: str) -> str:
        if check_for_language(language):
            return language
        raise LanguageCodeInvalid()


class MetaConfigQuerySerializer(Serializer):
    """
    MetaConfig
    """

    key = serializers.CharField(label=gettext_lazy("Config Key"), required=False)
