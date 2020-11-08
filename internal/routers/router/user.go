package router

import (
	"github.com/gin-gonic/gin"
	v1 "hao-admin/internal/routers/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup) {
	user := v1.NewUser()
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", user.Register)                 // 注册用户
		UserRouter.DELETE("deleteUser", user.DeleteUser)           // 删除用户
		UserRouter.POST("changePassword", user.ChangePassword)     // 修改密码
		UserRouter.POST("getUserList", user.GetUserList)           // 分页获取用户列表
		UserRouter.PUT("setUserInfo", user.SetUserInfo)            // 设置用户信息
		UserRouter.POST("setUserAuthority", user.SetUserAuthority) // 设置用户权限

	}
}
