from django.conf import settings
from django.contrib import admin
from django.contrib.staticfiles.views import serve
from django.urls import include, path, re_path
from django.views.generic import RedirectView
from ovinc_client.core import exceptions


# pylint: disable=W0621
def serve_static(request, path, insecure=True, **kwargs):
    return serve(request, path, insecure=True, **kwargs)


urlpatterns = [
    path("favicon.ico", RedirectView.as_view(url=f"{settings.FRONTEND_URL}/favicon.ico")),
    re_path(r"^static/(?P<path>.*)$", serve_static, name="static"),
    path("admin/", admin.site.urls),
    path("", include("apps.home.urls")),
    path("account/", include("apps.account.urls")),
    path("application/", include("apps.application.urls")),
    path("notice/", include("apps.notice.urls")),
    path("trace/", include("ovinc_client.trace.urls")),
    path("tcaptcha/", include("ovinc_client.tcaptcha.urls")),
]

handler400 = exceptions.bad_request
handler403 = exceptions.permission_denied
handler404 = exceptions.page_not_found
handler500 = exceptions.server_error