package v1

import "github.com/gin-gonic/gin"

type User struct{}

func NewUser() User {
	return User{}
}

// 注册用户
func (u User) Register(c *gin.Context) {

}

// 删除用户
func (u User) DeleteUser(c *gin.Context) {

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
