package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "hao-admin/docs"
	"hao-admin/global"
	"hao-admin/internal/middleware"
	"hao-admin/internal/routers/api"
	"hao-admin/internal/routers/router"
	"hao-admin/pkg/limiter"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeOut(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())

	upload := api.NewUpload()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.POST("/upload/file", upload.UploadFile)

	ApiGroup := r.Group("/api/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	//apiv1.Use(middleware.JWT())
	//{
	//	apiv1.POST("/tags", tag.Create)
	//	apiv1.DELETE("/tags/:id", tag.Delete)
	//	apiv1.PUT("/tags/:id", tag.Update)
	//	apiv1.PATCH("/tags/:id/state", tag.Update)
	//	apiv1.GET("/tags", tag.List)
	//
	//	apiv1.POST("/articles", article.Create)
	//	apiv1.DELETE("/articles/:id", article.Delete)
	//	apiv1.PUT("/articles/:id", article.Update)
	//	apiv1.PATCH("/articles/:id/state", article.Update)
	//	apiv1.GET("/articles/:id", article.Get)
	//	apiv1.GET("/articles", article.List)
	//
	//}

	return r
}
