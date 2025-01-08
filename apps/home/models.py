from typing import Dict

from django.db import models
from django.db.models import Q
from django.utils.translation import gettext_lazy
from ovinc_client.core.models import BaseModel, UniqIDField


class MetaConfig(BaseModel):
    """
    Meta Config
    """

    id = UniqIDField(gettext_lazy("ID"), primary_key=True)
    key = models.CharField(gettext_lazy("Config Key"), unique=True, max_length=255)
    val = models.TextField(gettext_lazy("Config Value"), null=True, blank=True)
    is_public = models.BooleanField(gettext_lazy("Is Public"), default=True, db_index=True)

    class Meta:
        verbose_name = gettext_lazy("Meta Config")
        verbose_name_plural = gettext_lazy("Meta Configs")
        ordering = ["key"]

    def __str__(self):
        return str(self.key)

    @classmethod
    def as_map(cls, condition: Q) -> Dict[str, str]:
        return {config.key: config.val for config in cls.objects.filter(condition)}
