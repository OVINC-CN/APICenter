package httpd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/internal/apps/account"
	"github.com/ovinc-cn/apicenter/v2/internal/apps/account/accountRouters"
	"github.com/ovinc-cn/apicenter/v2/internal/apps/home"
	"github.com/ovinc-cn/apicenter/v2/internal/config"
	"github.com/ovinc-cn/apicenter/v2/pkg/apiMixin"
	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
	_ "github.com/ovinc-cn/apicenter/v2/support_files/api_docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Serve() {
	// init gin
	r := gin.New()

	// middlewares
	r.Use(gin.Recovery(), apiMixin.ObserveMiddleware())

	// routers
	apiV1Group := r.Group("/v1")
	{
		// home
		homeGroup := apiV1Group.Group("")
		{
			homeGroup.GET("/healthz", home.HealthZ)
		}
		// account
		accountGroup := apiV1Group.Group("/account")
		{
			accountGroup.POST("/sign-up", accountRouters.SignUp)
			accountGroup.POST("/sign-in", accountRouters.SignIn)
			accountGroup.GET("/user-info", account.LoginRequiredMiddleware(), accountRouters.UserInfo)
		}
	}

	// release mode
	if cfg.AppDebug() {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// start server
	if err := r.Run(config.Config.API.Addr); err != nil {
		log.Fatalf("[Http] server failed; %s", err)
	}
}
