# pylint: disable=R0801,C0103
# Generated by Django 4.2 on 2023-04-21 09:38

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("account", "0001_initial"),
    ]

    operations = [
        migrations.RemoveField(
            model_name="user",
            name="wechat_id",
        ),
        migrations.AddField(
            model_name="user",
            name="wechat_open_id",
            field=models.CharField(blank=True, max_length=255, null=True, verbose_name="Wechat Open ID"),
        ),
        migrations.AddField(
            model_name="user",
            name="wechat_union_id",
            field=models.CharField(blank=True, max_length=255, null=True, verbose_name="Wechat Union ID"),
        ),
    ]
