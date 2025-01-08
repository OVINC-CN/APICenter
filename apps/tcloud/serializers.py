from rest_framework import serializers

from apps.tcloud.constants import CICallbackEventName


class TCICallbackSerializer(serializers.Serializer):
    EventName = serializers.ChoiceField(choices=CICallbackEventName.choices)
    JobsDetail = serializers.JSONField()
