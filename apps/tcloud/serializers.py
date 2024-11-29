from adrf.serializers import Serializer
from rest_framework import serializers

from apps.tcloud.constants import CICallbackEventName


class TCICallbackSerializer(Serializer):
    EventName = serializers.ChoiceField(choices=CICallbackEventName.choices)
    JobsDetail = serializers.JSONField()
