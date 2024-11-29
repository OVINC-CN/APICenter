from django.conf import settings
from django.urls import include, path
from rest_framework.routers import SimpleRouter

from apps.tcloud.views import AuditCallbackViewSet

router = SimpleRouter()
router.register("", AuditCallbackViewSet, basename="audit_callback")

urlpatterns = [path(f"ci/callback/{settings.TCI_AUDIT_CALLBACK_PREFIX}", include(router.urls))]
