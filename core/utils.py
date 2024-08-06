from rest_framework.request import Request

from core.constants import WECHAT_USER_AGENT_KEY


def is_wechat(request: Request) -> bool:
    return str(request.headers.get("User-Agent") or "").lower().find(WECHAT_USER_AGENT_KEY) != -1
