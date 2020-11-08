package v1

import (
	"github.com/gin-gonic/gin"
	"hao-admin/global"
	"hao-admin/internal/model"
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
	var params service.ChangePasswordRequest
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.ChangePassword(&params)
	if err != nil {
		global.Logger.Errorf("svc.ChangePassword err: %v", err)
		response.ToErrorResponse(errcode.ErrorChangePasswordFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// 获取用户列表
func (u User) GetUserList(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	page := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountUser()
	if err != nil {
		global.Logger.Errorf("svc.CountUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountUserFail)
		return
	}
	users, err := svc.GetUserList(&page)
	if err != nil {
		global.Logger.Errorf("svc.GetUserList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserListFail)
		return
	}
	response.ToResponseList(users, totalRows)
	return

}

//  设置用户信息
func (u User) SetUserInfo(c *gin.Context) {
	var params model.User
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	reqUser, err := svc.SetUserInfo(&params)
	if err != nil {
		global.Logger.Errorf("svc.SetUserInfo err:%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateUserFail)
		return
	}
	response.ToResponse(gin.H{"userInfo": reqUser})
	return
}

// 设置用户权限
func (u User) SetUserAuthority(c *gin.Context) {
	var params service.SetUserAuth
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.SetUserAuthority(&params)
	if err != nil {
		global.Logger.Errorf("svc.SetUserAuthority err:%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateUserAuthFaile)
		return
	}
	response.ToResponse(gin.H{})
	return
}
