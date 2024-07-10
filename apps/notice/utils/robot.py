from typing import List, Union

import httpx
from channels.db import database_sync_to_async
from ovinc_client.core.logger import logger

from apps.notice.models import Robot
from apps.notice.utils.base import NoticeBase


class RobotHandler(NoticeBase):
    """
    Robot
    """

    property_key = None

    # pylint: disable=W0246
    def __init__(self, robots: List[str], content: Union[dict, str], **kwargs) -> None:
        super().__init__(robots, content, **kwargs)

    async def _send(self) -> list:
        result = []
        client = httpx.AsyncClient()
        try:
            for robot in self.receivers:
                resp = await client.post(robot, json=self.content)
                result.append(resp.json())
        finally:
            await client.aclose()
        return result

    # pylint: disable=W0237
    @database_sync_to_async
    def _load_receivers(self, robot_ids: List[str]) -> List[str]:
        """
        robot id
        """

        logger.info("[%s LoadReceivers] Robots => %s", self.__class__.__name__, robot_ids)
        robots = Robot.objects.filter(id__in=robot_ids)
        logger.info(
            "[%s LoadReceivers] Receivers => %s",
            self.__class__.__name__,
            [f"{r.id}:{r.webhook}" for r in robots],
        )
        return [r.webhook for r in robots]