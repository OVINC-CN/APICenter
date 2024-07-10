from rest_framework.routers import DefaultRouter

from apps.home.views import HomeView, I18nViewSet, MetaConfigViewSet

router = DefaultRouter()
router.register("", HomeView)
router.register("i18n", I18nViewSet, basename="i18n")
router.register("meta", MetaConfigViewSet)

urlpatterns = router.urls
