from dataclasses import fields

from apps.home.constants import BuildInKeys
from apps.home.models import MetaConfig


def registry_build_in_config_key(*args, **kwargs):
    """
    Registry Build In Config Key To DB
    """

    for field in fields(BuildInKeys):
        key, val = getattr(BuildInKeys, field.name)
        MetaConfig.objects.get_or_create(key=key, defaults={"val": val})
