from rest_framework.throttling import SimpleRateThrottle


class AppRateThrottle(SimpleRateThrottle):
    """
    App Throttle
    """

    scope = "app"

    def get_cache_key(self, request, view):
        return f'rate:{self.scope}:{getattr(request.user, "app_code", "")}'
