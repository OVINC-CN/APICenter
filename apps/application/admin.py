from django.contrib import admin

from apps.application.models import Application


@admin.register(Application)
class ApplicationAdmin(admin.ModelAdmin):
    list_display = ["app_code", "app_name", "is_deleted"]
    list_filter = ["is_deleted"]
    search_fields = ["app_code", "app_name"]
