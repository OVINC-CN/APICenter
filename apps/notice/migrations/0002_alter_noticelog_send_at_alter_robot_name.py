# pylint: disable=R0801,C0103
# Generated by Django 4.2.11 on 2024-05-08 03:02

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("notice", "0001_initial"),
    ]

    operations = [
        migrations.AlterField(
            model_name="noticelog",
            name="send_at",
            field=models.DateTimeField(auto_now_add=True, db_index=True, verbose_name="Send At"),
        ),
        migrations.AlterField(
            model_name="robot",
            name="name",
            field=models.CharField(db_index=True, max_length=32, verbose_name="Name"),
        ),
    ]
