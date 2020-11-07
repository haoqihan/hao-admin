package router

import (
	"github.com/gin-gonic/gin"
	v1 "hao-admin/internal/routers/api/v1"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	base := v1.NewBase()
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", base.Login)
		BaseRouter.POST("captcha", base.Captcha)
	}
	return BaseRouter
}
