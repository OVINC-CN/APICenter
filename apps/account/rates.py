from ovinc_client.core.utils import get_ip
from rest_framework.throttling import SimpleRateThrottle


class SMSRateThrottle(SimpleRateThrottle):
    """
    SMS Throttle
    """

    scope = "phone_verify_code"

    def get_cache_key(self, request, view) -> str:
        return f'rate:{self.scope}:{request.data.get("phone_area")}{request.data.get("phone_number")}'


class IPRateThrottle(SimpleRateThrottle):
    """
    IP Throttle
    """

    scope = "ip"

    def get_cache_key(self, request, view) -> str:
        return f"rate:{self.scope}:{get_ip(request)}"
