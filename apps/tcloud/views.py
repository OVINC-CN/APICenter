from django.conf import settings
from django.utils.translation import gettext
from ovinc_client.core.auth import SessionAuthenticate
from ovinc_client.core.viewsets import MainViewSet
from rest_framework.request import Request
from rest_framework.response import Response

from apps.cel.tasks import send_notice
from apps.notice.constants import NoticeWayChoices
from apps.tcloud.models import AuditCallback
from apps.tcloud.serializers import TCICallbackSerializer


class AuditCallbackViewSet(MainViewSet):
    """
    callback view set
    """

    authentication_classes = [SessionAuthenticate]

    def create(self, request: Request, *args, **kwargs) -> Response:
        """
        callback
        """
        # validate
        req_slz = TCICallbackSerializer(data=request.data)
        req_slz.is_valid(raise_exception=True)
        # save
        callback = AuditCallback.add_callback(req_slz.validated_data)
        if callback.is_sensitive:
            send_notice.delay(
                notice_type=NoticeWayChoices.ROBOT,
                robots=[settings.NOTICE_AUDIT_SENSITIVE_ROBOT],
                content={
                    "msgtype": "markdown",
                    "markdown": {
                        "content": gettext(
                            '<font color="warning">Sensitive Object Notification</font>\n'
                            "Bucket Region: %(region)s\n"
                            "Bucket ID: %(bucket)s\n"
                            "Event Name: %(event)s\n"
                            "Audit ID: %(audit_id)s\n"
                            "Creation Time: %(creation_time)s\n"
                        )
                        % {
                            "region": callback.bucket_region,
                            "bucket": callback.bucket_id,
                            "event": callback.get_event_name_display(),
                            "audit_id": callback.audit_id,
                            "creation_time": callback.creation_time,
                        }
                    },
                },
            )
        # response
        return Response()
