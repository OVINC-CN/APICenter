from django.apps import AppConfig
from django.utils.translation import gettext_lazy


class TCloudConfig(AppConfig):
    default_auto_field = "django.db.models.BigAutoField"
    name = "apps.tcloud"
    verbose_name = gettext_lazy("TCloud")
