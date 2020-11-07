package dao

import (
	"hao-admin/internal/model"
	"hao-admin/pkg/util"
)

func (d *Dao) Login(username string, password string) (error, *model.User) {
	user := model.User{Username: username}
	user.Password = util.EncodeMD5(password)
	return user.Login(d.engine)
}
