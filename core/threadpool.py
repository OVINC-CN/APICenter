from concurrent.futures.thread import ThreadPoolExecutor

from django.conf import settings

db_executor = ThreadPoolExecutor(max_workers=settings.DB_EXECUTOR_SIZE)
