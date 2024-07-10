from adrf.serializers import ModelSerializer
from django.contrib.auth import get_user_model

from apps.account.models import User
from apps.application.models import Application

USER_MODEL: User = get_user_model()


class ApplicationSerializer(ModelSerializer):
    """
    Application Info
    """

    class Meta:
        model = Application
        fields = ["app_name", "app_code"]


class ApplicationListSerializer(ModelSerializer):
    """
    Application List
    """

    class Meta:
        model = Application
        fields = ["app_name", "app_code", "app_logo", "app_url", "app_desc"]
