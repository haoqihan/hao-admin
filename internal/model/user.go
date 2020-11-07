package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	*Model
	UUID        uuid.UUID `json:"uuid"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	NickName    string    `json:"nick_name"`
	HeaderImage string    `json:"header_image" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Authority   Authority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string    `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}

func (u *User) Register(db *gorm.DB) (err error, userInter *User) {
	var user User
	// 判断用户是否注册
	if !errors.Is(db.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已经注册"), userInter
	}

	return db.Create(&u).Error, u
}

func (u *User) Login(db *gorm.DB) (err error, user *User) {
	err = db.Where("username = ? and password = ?", u.Username, u.Password).Preload("Authority").First(user).Error
	return err, user
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", u.Model.ID, 0).Delete(&u).Error
}
