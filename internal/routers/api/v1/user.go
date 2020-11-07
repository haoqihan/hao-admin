package v1

import (
	"github.com/gin-gonic/gin"
	"hao-admin/global"
	"hao-admin/internal/service"
	"hao-admin/pkg/app"
	"hao-admin/pkg/errcode"
)

type User struct{}

func NewUser() User {
	return User{}
}

// 注册用户
func (u User) Register(c *gin.Context) {
	var params service.RegisterRequest
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err, user := svc.Register(&params)
	if err != nil {
		global.Logger.Errorf("svc.Register err: %v", err)
		response.ToErrorResponse(errcode.ErrorRegisterUserFail)
		return
	}
	response.ToResponse(gin.H{"user": user})
	return
}

// 删除用户
func (u User) DeleteUser(c *gin.Context) {
	var params service.DeleteUserRequest
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteUser(&params)
	if err != nil {
		global.Logger.Errorf("svc.DeleteUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteUserFail)
		return
	}
	response.ToResponse(gin.H{})
	return

}

// 修改密码
func (u User) ChangePassword(c *gin.Context) {

}

// 获取用户列表
func (u User) GetUserList(c *gin.Context) {

}

//  设置用户信息
func (u User) SetUserInfo(c *gin.Context) {

}
