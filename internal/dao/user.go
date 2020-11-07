package dao

import (
	uuid "github.com/satori/go.uuid"
	"hao-admin/internal/model"
	"hao-admin/pkg/util"
)

func (d *Dao) Register(username, nickname, password, header_image, authority_id string) (error, *model.User) {
	user := &model.User{
		Username:    username,
		NickName:    nickname,
		HeaderImage: header_image,
		AuthorityId: authority_id,
	}
	user.Password = util.EncodeMD5(password)
	user.UUID = uuid.NewV4()
	return user.Register(d.engine)
}

func (d *Dao) DeleteUser(id uint32) error {
	user := model.User{Model: &model.Model{ID: id}}
	return user.Delete(d.engine)
}