from rest_framework.routers import DefaultRouter

from apps.home.views import HealthViewSet, HomeView, I18nViewSet, MetaConfigViewSet

router = DefaultRouter()
router.register("", HomeView)
router.register("", HealthViewSet, basename="health")
router.register("i18n", I18nViewSet, basename="i18n")
router.register("meta", MetaConfigViewSet)

urlpatterns = router.urls
