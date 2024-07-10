from django.utils.translation import gettext_lazy
from rest_framework import status
from rest_framework.exceptions import APIException


class AuthTokenInvalid(APIException):
    status_code = status.HTTP_400_BAD_REQUEST
    default_detail = gettext_lazy("Auth Token Invalid")


class AppAuthFailed(APIException):
    status_code = status.HTTP_403_FORBIDDEN
    default_detail = gettext_lazy("App Auth Failed")
