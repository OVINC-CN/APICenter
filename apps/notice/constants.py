from django.utils.translation import gettext_lazy
from ovinc_client.core.models import TextChoices


class NoticeWayChoices(TextChoices):
    MSG = "msg", gettext_lazy("msg")
    MAIL = "mail", gettext_lazy("mail")
    ROBOT = "robot", gettext_lazy("robot")
