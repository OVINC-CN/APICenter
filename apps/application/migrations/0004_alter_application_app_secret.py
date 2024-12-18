# pylint: disable=C0103
# Generated by Django 4.2.17 on 2024-12-18 07:18

from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("application", "0003_application_is_hidden"),
    ]

    operations = [
        migrations.AlterField(
            model_name="application",
            name="app_secret",
            field=models.CharField(blank=True, max_length=255, null=True, verbose_name="App Secret"),
        ),
    ]
