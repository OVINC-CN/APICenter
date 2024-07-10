from django.apps import AppConfig
from django.db.models.signals import post_migrate
from django.utils.translation import gettext_lazy


class HomeConfig(AppConfig):
    default_auto_field = "django.db.models.BigAutoField"
    name = "apps.home"
    verbose_name = gettext_lazy("Home Module")

    def ready(self):
        # pylint: disable=C0415
        from apps.home.utils import registry_build_in_config_key

        post_migrate.connect(registry_build_in_config_key, self)
