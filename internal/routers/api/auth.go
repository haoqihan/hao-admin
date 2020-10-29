package api

import (
	"hao-admin/global"
	"hao-admin/internal/service"
	"hao-admin/pkg/app"
	"hao-admin/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid : %v", errors)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{"token": token})
}
