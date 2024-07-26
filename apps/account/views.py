import datetime
import json
from json import JSONDecodeError

import httpx
from channels.db import database_sync_to_async
from django.conf import settings
from django.contrib import auth
from django.contrib.auth import get_user_model
from django.core.cache import cache
from ovinc_client.core.auth import SessionAuthenticate
from ovinc_client.core.logger import logger
from ovinc_client.core.utils import get_ip, uniq_id
from ovinc_client.core.viewsets import MainViewSet
from rest_framework.decorators import action
from rest_framework.response import Response

from apps.account.constants import (
    WECHAT_LOGIN_STATE_KEY,
    WECHAT_USER_INFO_KEY,
    PhoneNumberAreas,
    WeChatAuthType,
)
from apps.account.exceptions import (
    StateInvalid,
    UserNotExist,
    WeChatLoginFailed,
    WrongSignInParam,
    WrongToken,
)
from apps.account.models import User
from apps.account.rates import IPRateThrottle, SMSRateThrottle
from apps.account.serializers import (
    ResetPasswordRequestSerializer,
    SendVerifyCodeRequestSerializer,
    SignInSerializer,
    UserInfoSerializer,
    UserRegistrySerializer,
    VerifyCodeRequestSerializer,
    WeChatLoginReqSerializer,
)
from core.auth import ApplicationAuthenticate

USER_MODEL: User = get_user_model()


class UserInfoViewSet(MainViewSet):
    """
    User Info
    """

    queryset = USER_MODEL.get_queryset()
    serializer_class = UserInfoSerializer

    async def list(self, request, *args, **kwargs):
        """
        Get User Info
        """

        serializer = self.serializer_class(instance=request.user)
        return Response(await serializer.adata)


class UserSignViewSet(MainViewSet):
    """
    User Login and Logout
    """

    queryset = USER_MODEL.get_queryset()

    @action(methods=["POST"], detail=False, authentication_classes=[SessionAuthenticate])
    async def sign_in(self, request, *args, **kwargs):
        """
        Sign in
        """

        # Validate Request Data
        request_serializer = SignInSerializer(data=request.data, context={"user_ip": get_ip(request)})
        request_serializer.is_valid(raise_exception=True)
        request_data = request_serializer.validated_data

        # login
        user: User = await database_sync_to_async(auth.authenticate)(request, **request_data)
        if not user:
            raise WrongSignInParam()

        # bind wechat
        if request_data.get("wechat_code"):
            await self.update_user_by_wechat(user, request_data["wechat_code"])

        # auth session
        await database_sync_to_async(auth.login)(request, user)

        # oauth
        if request_data["is_oauth"]:
            return Response({"code": user.generate_oauth_code()})

        return Response()

    @action(methods=["GET"], detail=False)
    async def sign_out(self, request, *args, **kwargs):
        """
        Sign out
        """

        await database_sync_to_async(auth.logout)(request)
        return Response()

    @action(methods=["POST"], detail=False, authentication_classes=[SessionAuthenticate])
    async def sign_up(self, request, *args, **kwargs):
        """
        sign up
        """

        # validate request
        request_serializer = UserRegistrySerializer(data=request.data, context={"user_ip": get_ip(request)})
        request_serializer.is_valid(raise_exception=True)
        request_data = request_serializer.validated_data

        # save
        user = await database_sync_to_async(USER_MODEL.objects.create_user)(
            last_login=datetime.datetime.now(),
            username=request_data["username"],
            password=request_data["password"],
            nick_name=request_data["nick_name"],
            phone_number=request_data["phone_number"],
        )

        # bind wechat
        if request_data.get("wechat_code"):
            await self.update_user_by_wechat(user, request_data["wechat_code"])

        # login session
        await database_sync_to_async(auth.login)(request, user)

        # oauth
        if request_data["is_oauth"]:
            return Response({"code": user.generate_oauth_code()})

        # response
        return Response()

    @action(
        methods=["POST"],
        detail=False,
        authentication_classes=[SessionAuthenticate],
        throttle_classes=[SMSRateThrottle, IPRateThrottle],
    )
    async def phone_verify_code(self, request, *args, **kwargs):
        """
        send verify code
        """

        # validate request
        request_serializer = SendVerifyCodeRequestSerializer(data=request.data, context={"user_ip": get_ip(request)})
        request_serializer.is_valid(raise_exception=True)
        request_data = request_serializer.validated_data

        # send
        USER_MODEL.send_phone_verify_code(area=request_data["phone_area"], phone_number=request_data["phone_number"])

        # response
        return Response()

    @action(methods=["GET"], detail=False)
    async def oauth_code(self, request, *args, **kwargs):
        """
        oauth code
        """

        return Response({"code": request.user.generate_oauth_code()})

    @action(methods=["POST"], detail=False, authentication_classes=[ApplicationAuthenticate])
    async def verify_code(self, request, *args, **kwargs):
        """
        verify oauth code
        """

        # validate request
        request_serializer = VerifyCodeRequestSerializer(data=request.data)
        request_serializer.is_valid(raise_exception=True)
        request_data = request_serializer.validated_data

        # load user
        is_success, user = await database_sync_to_async(USER_MODEL.check_oauth_code)(request_data["code"])
        if is_success:
            return Response(await UserInfoSerializer(instance=user).adata)
        raise WrongToken()

    @action(methods=["GET"], detail=False, authentication_classes=[SessionAuthenticate])
    async def wechat_config(self, request, *args, **kwargs):
        """
        WeChat Config
        """

        state = uniq_id()
        cache_key = WECHAT_LOGIN_STATE_KEY.format(state=state)
        cache.set(cache_key, True, timeout=settings.WECHAT_SCOPE_TIMEOUT)
        return Response({"app_id": settings.WECHAT_APP_ID, "state": state})

    @action(methods=["POST"], detail=False, authentication_classes=[SessionAuthenticate])
    async def wechat_login(self, request, *args, **kwargs):
        """
        WeChat Login
        """

        # validate request
        request_serializer = WeChatLoginReqSerializer(data=request.data)
        request_serializer.is_valid(raise_exception=True)
        request_data = request_serializer.validated_data

        # check state
        cache_key = WECHAT_LOGIN_STATE_KEY.format(state=request_data["state"])
        if not cache.get(cache_key):
            raise StateInvalid()
        cache.delete(cache_key)

        # load access token
        url = (
            f"{settings.WECHAT_OAUTH_API}/access_token"
            f"?appid={settings.WECHAT_APP_ID}"
            f"&secret={settings.WECHAT_APP_KEY}"
            f"&code={request_data['code']}"
            f"&grant_type={WeChatAuthType.CODE}"
        )
        client = httpx.AsyncClient()
        try:
            resp = await client.get(url)
            access_info = resp.json()
        except Exception as err:
            logger.exception("[CallWeChatAPIFailed] %s %s", url, err)
            raise WeChatLoginFailed() from err
        finally:
            await client.aclose()

        if "openid" not in access_info:
            logger.error("[WeChatLoginFailed] %s", access_info)
            raise WeChatLoginFailed()

        # get user info
        url = (
            f"{settings.WECHAT_USER_INFO_API}?access_token={access_info['access_token']}&openid={access_info['openid']}"
        )
        client = httpx.AsyncClient()
        try:
            resp = await client.get(url)
            user_info = resp.json()
        except Exception as err:
            logger.exception("[CallWeChatAPIFailed] %s", url, err)
            raise WeChatLoginFailed() from err
        finally:
            await client.aclose()

        code = uniq_id()
        cache_key = WECHAT_USER_INFO_KEY.format(code=code)
        cache.set(cache_key, json.dumps(user_info, ensure_ascii=False), timeout=settings.WECHAT_SCOPE_TIMEOUT)

        # load user
        user: User = await database_sync_to_async(USER_MODEL.load_user_by_union_id)(union_id=user_info["unionid"])
        if user:
            await self.update_user_by_wechat(user, code)
            await database_sync_to_async(auth.login)(request, user)
            return Response({"code": user.generate_oauth_code() if request_data["is_oauth"] else ""})

        # need registry
        return Response({"wechat_code": code})

    async def update_user_by_wechat(self, user: User, wechat_code: str) -> None:
        """
        Update User Info By WeChat
        """

        cache_key = WECHAT_USER_INFO_KEY.format(code=wechat_code)
        try:
            user_info = json.loads(cache.get(cache_key, default="{}"))
        except JSONDecodeError:
            return

        cache.delete(cache_key)

        user.wechat_union_id = user_info["unionid"]
        user.wechat_open_id = user_info["openid"]
        user.avatar = user_info["headimgurl"]
        await database_sync_to_async(user.save)(update_fields=["wechat_union_id", "wechat_open_id", "avatar"])

    @action(methods=["POST"], detail=False, authentication_classes=[SessionAuthenticate])
    async def reset_password(self, request, *args, **kwargs) -> Response:
        """
        Reset Password
        """

        # validate request
        request_serializer = ResetPasswordRequestSerializer(data=request.data, context={"user_ip": get_ip(request)})
        request_serializer.is_valid(raise_exception=True)
        request_data = request_serializer.validated_data

        # load user
        try:
            user: User = await database_sync_to_async(USER_MODEL.objects.get)(
                username=request_data["username"], phone_number=request_data["phone_number"]
            )
        except USER_MODEL.DoesNotExist as err:
            raise UserNotExist() from err

        # set new password
        await database_sync_to_async(user.reset_password)(request_data["password"])

        return Response()

    @action(methods=["GET"], detail=False, authentication_classes=[SessionAuthenticate])
    async def phone_areas(self, request, *args, **kwargs) -> Response:
        """
        Phone Number Areas
        """

        return Response(data=[{"value": value, "label": str(label)} for value, label in PhoneNumberAreas.choices])
