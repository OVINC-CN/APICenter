FROM python:3.10.15-slim

WORKDIR /usr/src/app

COPY requirements.txt /usr/src/app/
RUN pip3 install -U pip && pip3 install -r requirements.txt

COPY . /usr/src/app

RUN apt-get update && apt-get install -y gettext

RUN cp env.example .env \
    && python3 manage.py compilemessages -l zh_Hans \
    && rm -rf .env

RUN mkdir -p /usr/src/app/celery-logs /usr/src/app/logs  /usr/src/app/daphne-logs
