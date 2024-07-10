from django.contrib import admin

from apps.notice.models import NoticeLog, Robot


@admin.register(Robot)
class RobotAdmin(admin.ModelAdmin):
    list_display = ["id", "name", "webhook"]
    ordering = ["name"]


@admin.register(NoticeLog)
class NoticeLogAdmin(admin.ModelAdmin):
    list_display = ["id", "receivers", "content", "result", "send_at"]
    ordering = ["-send_at"]
