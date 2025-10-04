package httpd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/internal/apps/home"
	"github.com/ovinc-cn/apicenter/v2/internal/config"
	"github.com/ovinc-cn/apicenter/v2/pkg/configUtils"
	"github.com/ovinc-cn/apicenter/v2/pkg/ginUtils"
)

func Serve() {
	// release mode
	if configUtils.AppDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// init gin
	r := gin.New()

	// middlewares
	r.Use(gin.Recovery(), ginUtils.ObserveMiddleware())

	// routers
	apiV1Group := r.Group("/v1")
	{
		// home
		homeGroup := apiV1Group.Group("")
		{
			homeGroup.GET("/healthz", home.HealthZ)
		}
	}

	// start server
	if err := r.Run(config.Config.API.Addr); err != nil {
		log.Fatalf("[Http] server failed; %s", err)
	}
}
