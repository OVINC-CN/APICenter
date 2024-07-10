from django.contrib import admin

from apps.home.models import MetaConfig


@admin.register(MetaConfig)
class MetaConfigAdmin(admin.ModelAdmin):
    list_display = ["key", "is_public", "val"]
    ordering = ["key"]
