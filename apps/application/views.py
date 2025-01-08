from django.contrib.auth import get_user_model
from ovinc_client.core.auth import SessionAuthenticate
from ovinc_client.core.viewsets import MainViewSet
from rest_framework.decorators import action
from rest_framework.response import Response

from apps.account.models import User
from apps.application.models import Application
from apps.application.serializers import (
    ApplicationListSerializer,
    ApplicationSerializer,
)

USER_MODEL: User = get_user_model()


class ApplicationViewSet(MainViewSet):
    """
    Application
    """

    queryset = Application.get_queryset()
    serializer_class = ApplicationSerializer

    @action(methods=["GET"], detail=False, authentication_classes=[SessionAuthenticate])
    def all(self, request, *args, **kwargs):
        """
        list all applications
        """

        # filter
        queryset = Application.get_queryset().filter(is_hidden=False)

        # response
        serializer = ApplicationListSerializer(queryset, many=True)
        return Response(serializer.data)
