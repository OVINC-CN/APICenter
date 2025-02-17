from rest_framework.settings import api_settings


def userinfo(claims, user):
    claims.update(
        {
            "name": user.username,
            "nickname": user.nick_name,
            "updated_at": user.last_login.strftime(api_settings.DATETIME_FORMAT),
        }
    )
    return claims


def default_sub_generator(user):
    return f"{user.username}"
