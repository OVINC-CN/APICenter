from asgiref.sync import async_to_sync
from ovinc_client.core.logger import celery_logger

from apps.cel import app


@app.task(bind=True)
def send_notice(self, notice_type: str, **kwargs):
    # pylint: disable=C0415,R0401
    from apps.notice.utils import NoticeBase

    celery_logger.info(f"[SendNotice] Start {self.request.id}")
    async_to_sync(NoticeBase.get_instance(notice_type, **kwargs).send)()
    celery_logger.info(f"[SendNotice] End {self.request.id}")
