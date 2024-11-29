from django.contrib.auth import get_user_model
from django.db import models
from django.utils.translation import gettext_lazy
from ovinc_client.core.constants import MAX_CHAR_LENGTH, SHORT_CHAR_LENGTH
from ovinc_client.core.models import BaseModel

from apps.tcloud.constants import CICallbackEventName
from apps.tcloud.parsers import TCICallbackParser, TCIDocumentCallbackParser

USER_MODEL = get_user_model()


class AuditCallback(BaseModel):
    """
    Audit Callback
    """

    id = models.BigAutoField(verbose_name=gettext_lazy("ID"), primary_key=True)
    bucket_region = models.CharField(verbose_name=gettext_lazy("Bucket Region"), max_length=SHORT_CHAR_LENGTH)
    bucket_id = models.CharField(verbose_name=gettext_lazy("Bucket ID"), max_length=MAX_CHAR_LENGTH, db_index=True)
    event_name = models.CharField(
        verbose_name=gettext_lazy("Event Name"),
        max_length=SHORT_CHAR_LENGTH,
        db_index=True,
        choices=CICallbackEventName.choices,
    )
    audit_id = models.CharField(verbose_name=gettext_lazy("Audit ID"), max_length=MAX_CHAR_LENGTH, db_index=True)
    is_sensitive = models.BooleanField(verbose_name=gettext_lazy("Is Sensitive"), db_index=True)
    detail = models.JSONField(verbose_name=gettext_lazy("Detail"))
    creation_time = models.CharField(
        verbose_name=gettext_lazy("Audit Time"), max_length=SHORT_CHAR_LENGTH, db_index=True, null=True, blank=True
    )
    callback_time = models.DateTimeField(verbose_name=gettext_lazy("Callback Time"), db_index=True, auto_now_add=True)
    updated_time = models.DateTimeField(verbose_name=gettext_lazy("Updated Time"), db_index=True, auto_now=True)
    is_handled = models.BooleanField(verbose_name=gettext_lazy("Handled"), db_index=True, default=False)

    class Meta:
        verbose_name = gettext_lazy("Audit Callback")
        verbose_name_plural = verbose_name
        ordering = ["-id"]
        unique_together = [["bucket_region", "bucket_id", "event_name", "audit_id"]]

    def __str__(self):
        return f"{self.bucket_region}:{self.bucket_id}:{self.event_name}:{self.audit_id}"

    @classmethod
    def add_callback(cls, data: dict) -> "AuditCallback":
        event_name = data["EventName"]
        match event_name:
            case CICallbackEventName.REVIEW_DOCUMENT:
                parser = TCIDocumentCallbackParser(data)
            case _:
                parser = TCICallbackParser(data)
        return cls.objects.create(
            bucket_region=parser.bucket_region,
            bucket_id=parser.bucket_id,
            event_name=event_name,
            audit_id=parser.audit_id,
            is_sensitive=parser.is_sensitive,
            detail=parser.detail,
            creation_time=parser.creation_time,
        )
