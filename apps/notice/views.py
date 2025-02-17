from django.shortcuts import get_object_or_404
from ovinc_client.core.viewsets import MainViewSet
from rest_framework.decorators import action
from rest_framework.response import Response

from apps.notice.constants import NoticeWayChoices
from apps.notice.models import NoticeLog, Robot
from apps.notice.rates import AppRateThrottle
from apps.notice.serializers import (
    MailRequestSerialzier,
    RegistryRobotSerializer,
    RobotRequestSerializer,
    SmsRequestSerializer,
)
from apps.notice.utils import NoticeBase
from core.auth import ApplicationAuthenticate


class NoticeViewSet(MainViewSet):
    """
    Notice
    """

    queryset = NoticeLog.get_queryset()
    authentication_classes = [ApplicationAuthenticate]
    throttle_classes = [AppRateThrottle]

    @action(methods=["POST"], detail=False)
    def mail(self, request, *args, **kwargs):
        """
        mail
        """

        # validate request
        request_serializer = MailRequestSerialzier(data=request.data)
        request_serializer.is_valid(raise_exception=True)
        request_data = request_serializer.validated_data

        # send
        NoticeBase.send_notice(NoticeWayChoices.MAIL, **request_data)
        return Response()

    @action(methods=["POST"], detail=False)
    def sms(self, request, *args, **kwargs):
        """
        sms
        """

        # validate request
        request_serializer = SmsRequestSerializer(data=request.data)
        request_serializer.is_valid(raise_exception=True)
        request_data = request_serializer.validated_data

        # send
        NoticeBase.send_notice(NoticeWayChoices.MSG, **request_data)
        return Response()

    @action(methods=["POST"], detail=False)
    def robot(self, request, *args, **kwargs):
        """
        robot of Wecom or Feishu
        """

        # validate request
        request_serializer = RobotRequestSerializer(data=request.data)
        request_serializer.is_valid(raise_exception=True)
        request_data = request_serializer.validated_data

        # send
        NoticeBase.send_notice(NoticeWayChoices.ROBOT, **request_data)
        return Response()

    @action(methods=["POST"], detail=False)
    def registry_robot(self, request, *args, **kwargs):
        """
        registry robot
        """

        # get instance
        instance = None
        if request.data.get("id"):
            instance = get_object_or_404(Robot, pk=request.data.pop("id"))

        # validate request
        request_serializer = RegistryRobotSerializer(instance=instance, data=request.data, partial=bool(instance))
        request_serializer.is_valid(raise_exception=True)

        # save
        request_serializer.save()

        return Response(request_serializer.data)
