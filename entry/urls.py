from urllib.parse import quote

from django.conf import settings
from django.contrib import admin
from django.contrib.staticfiles.views import serve
from django.urls import include, path, re_path
from django.views.generic import RedirectView
from ovinc_client.core import exceptions


# pylint: disable=W0621
def serve_static(request, path, insecure=True, **kwargs):
    return serve(request, path, insecure=True, **kwargs)


ADMIN_PAGE_URL = f"{settings.BACKEND_URL}/admin/"
ADMIN_PAGE_LOGIN_URL = f"{settings.FRONTEND_URL}/login/?next={quote(ADMIN_PAGE_URL)}"

urlpatterns = [
    path("favicon.ico", RedirectView.as_view(url=f"{settings.FRONTEND_URL}/favicon.ico")),
    re_path(r"^static/(?P<path>.*)$", serve_static, name="static"),
    path("admin/login/", RedirectView.as_view(url=ADMIN_PAGE_LOGIN_URL.replace("%", "%%"))),
    path("admin/", admin.site.urls),
    path("openid/", include("oidc_provider.urls", namespace="oidc_provider")),
    path("", include("apps.home.urls")),
    path("account/", include("apps.account.urls")),
    path("application/", include("apps.application.urls")),
    path("notice/", include("apps.notice.urls")),
    path("tcloud/", include("apps.tcloud.urls")),
    path("trace/", include("ovinc_client.trace.urls")),
    path("tcaptcha/", include("ovinc_client.tcaptcha.urls")),
]

handler400 = exceptions.bad_request
handler403 = exceptions.permission_denied
handler404 = exceptions.page_not_found
handler500 = exceptions.server_error
