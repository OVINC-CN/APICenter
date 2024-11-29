# pylint: disable=C0301,C0103
# Generated by Django 4.2.16 on 2024-11-29 12:00

from django.db import migrations, models


class Migration(migrations.Migration):
    initial = True

    dependencies = []

    operations = [
        migrations.CreateModel(
            name="AuditCallback",
            fields=[
                ("id", models.BigAutoField(primary_key=True, serialize=False, verbose_name="ID")),
                ("bucket_region", models.CharField(max_length=32, verbose_name="Bucket Region")),
                ("bucket_id", models.CharField(db_index=True, max_length=255, verbose_name="Bucket ID")),
                (
                    "event_name",
                    models.CharField(
                        choices=[
                            ("ReviewImage", "Review Image"),
                            ("ReviewDocument", "Review Document"),
                            ("ReviewVideo", "Review Video"),
                            ("ReviewAudio", "Review Audio"),
                            ("ReviewText", "Review Text"),
                            ("ReviewHtml", "Review HTML"),
                        ],
                        db_index=True,
                        max_length=32,
                        verbose_name="Event Name",
                    ),
                ),
                ("audit_id", models.CharField(db_index=True, max_length=255, verbose_name="Audit ID")),
                ("is_sensitive", models.BooleanField(db_index=True, verbose_name="Result")),
                ("detail", models.JSONField(verbose_name="Detail")),
                ("creation_time", models.DateTimeField(db_index=True, verbose_name="Audit Time")),
                ("callback_time", models.DateTimeField(auto_now_add=True, db_index=True, verbose_name="Callback Time")),
                ("updated_time", models.DateTimeField(auto_now=True, db_index=True, verbose_name="Updated Time")),
                ("is_handled", models.BooleanField(db_index=True, default=False, verbose_name="Handled")),
            ],
            options={
                "verbose_name": "Audit Callback",
                "verbose_name_plural": "Audit Callback",
                "ordering": ["-id"],
                "unique_together": {("bucket_region", "bucket_id", "event_name", "audit_id")},
            },
        ),
    ]
