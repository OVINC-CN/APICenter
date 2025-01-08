import abc
import traceback
from typing import List, Union

from django.contrib.auth import get_user_model
from django.utils.module_loading import import_string
from django.utils.translation import gettext
from ovinc_client.core.logger import logger

from apps.account.models import User
from apps.cel.tasks import send_notice
from apps.notice.constants import NoticeWayChoices
from apps.notice.models import NoticeLog

USER_MODEL: User = get_user_model()


class NoticeBase:
    """
    Notice Base
    """

    # pylint: disable=R1710
    @classmethod
    def send_notice(cls, notice_type: str, is_async: bool = True, **kwargs) -> None:
        if is_async:
            send_notice.delay(notice_type, **kwargs)
            return None
        send_notice(notice_type, **kwargs)

    @classmethod
    def get_instance(cls, notice_type: str, **kwargs) -> "NoticeBase":
        instance_map = {_value: f"apps.notice.utils.{_value.capitalize()}Handler" for _value in NoticeWayChoices.values}
        try:
            notice_type_class = import_string(instance_map[notice_type])
            return notice_type_class(**kwargs)
        except KeyError as err:
            raise KeyError(gettext("Notice Type not Exists => %s") % notice_type) from err

    def __init__(self, usernames: List[str], content: Union[dict, str], **kwargs) -> None:
        self.kwargs = kwargs
        self.receivers = self._load_receivers(usernames)
        self.content = self._build_content(content)

    @property
    @abc.abstractmethod
    def property_key(self) -> str:
        raise NotImplementedError

    def send(self) -> None:
        """
        send notice
        """

        logger.info("[%s SendNotice] Content => %s", self.__class__.__name__, self.content)
        try:
            result = self._send()
            logger.info("[%s SendNoticeSuccess] Result => %s", self.__class__.__name__, result)
        except Exception as err:  # pylint: disable=W0718
            msg = traceback.format_exc()
            logger.error("[%s SendNoticeFailed] Err => %s; Detail => %s", self.__class__.__name__, err, msg)
            result = {"err": str(err)}
        NoticeLog.objects.create(
            receivers=self.receivers, content=self.content, extra_params=self.kwargs, result=str(result)
        )
        return result

    @abc.abstractmethod
    def _send(self) -> None:
        """
        send notice
        """

        raise NotImplementedError

    def _load_receivers(self, usernames: List[str]) -> List[str]:
        """
        trans username to receiver
        """

        logger.info("[%s LoadReceivers] Usernames => %s", self.__class__.__name__, usernames)
        # load user
        users = list(User.objects.filter(username__in=usernames))
        receivers = list(getattr(u, self.property_key) for u in users if getattr(u, self.property_key, None))
        if not receivers and "receivers" in self.kwargs:
            receivers = self.kwargs["receivers"]
        logger.info("[%s LoadReceivers] Receivers => %s", self.__class__.__name__, receivers)
        # return
        return receivers

    def _build_content(self, content: Union[dict, str]) -> any:
        """
        build content for notice
        """

        return content
