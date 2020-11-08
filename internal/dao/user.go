package dao

import (
	uuid "github.com/satori/go.uuid"
	"hao-admin/internal/model"
	"hao-admin/pkg/app"
	"hao-admin/pkg/util"
)

func (d *Dao) Register(username, nickname, password, headerImage, authorityId string) (error, *model.User) {
	user := &model.User{
		Username:    username,
		NickName:    nickname,
		HeaderImage: headerImage,
		AuthorityId: authorityId,
	}
	user.Password = util.EncodeMD5(password)
	user.UUID = uuid.NewV4()
	return user.Register(d.engine)
}

func (d *Dao) DeleteUser(id uint32) error {
	user := model.User{Model: &model.Model{ID: id}}
	return user.Delete(d.engine)
}

func (d *Dao) ChangePassword(username, password, newPassword string) error {
	user := &model.User{
		Username: username,
	}
	user.Password = util.EncodeMD5(password)
	return user.ChangePassword(d.engine, newPassword)
}

func (d *Dao) CountUser() (int, error) {
	user := &model.User{}
	return user.CountUser(d.engine)
}

func (d *Dao) GetUserList( page, pageSize int) ([]*model.User, error) {
	user := model.User{}
	pageOffset := app.GetPageOffset(page, pageSize)
	return user.List(d.engine, pageOffset, pageSize)
}
func (d *Dao) SetUserInfo(user *model.User)(*model.User,error) {
	return user.Updates(d.engine)
}

func (d *Dao) SetUserAuthority(uuid uuid.UUID,authorityId string) error {
	user := model.User{UUID:uuid}
	values := map[string]interface{}{
		"authority_id":authorityId,
	}
	return user.SetUserAuthority(d.engine,values)
}