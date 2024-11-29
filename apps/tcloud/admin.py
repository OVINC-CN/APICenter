from django.contrib import admin

from apps.tcloud.models import AuditCallback


@admin.register(AuditCallback)
class AuditCallbackAdmin(admin.ModelAdmin):
    list_display = [
        "id",
        "bucket_region",
        "bucket_id",
        "event_name",
        "audit_id",
        "is_sensitive",
        "creation_time",
        "callback_time",
        "updated_time",
        "is_handled",
    ]
    list_filter = ["bucket_region", "bucket_id", "event_name", "is_sensitive", "is_handled"]
    list_editable = ["is_handled"]
