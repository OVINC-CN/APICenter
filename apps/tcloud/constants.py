from django.utils.translation import gettext_lazy
from ovinc_client.core.models import TextChoices


class CICallbackEventName(TextChoices):
    REVIEW_IMAGE = "ReviewImage", gettext_lazy("Review Image")
    REVIEW_DOCUMENT = "ReviewDocument", gettext_lazy("Review Document")
    REVIEW_VIDEO = "ReviewVideo", gettext_lazy("Review Video")
    REVIEW_AUDIO = "ReviewAudio", gettext_lazy("Review Audio")
    REVIEW_TEXT = "ReviewText", gettext_lazy("Review Text")
    REVIEW_HTML = "ReviewHtml", gettext_lazy("Review HTML")
