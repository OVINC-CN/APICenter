package home

import (
	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/pkg/ginUtils"
)

func HealthZ(c *gin.Context) {
	ginUtils.ResponseSuccess(c, nil)
}
