package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"hao-admin/pkg/util"
)

type User struct {
	*Model
	UUID        uuid.UUID `json:"uuid"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	NickName    string    `json:"nick_name"`
	HeaderImage string    `json:"header_image" gorm:"default:'http://qmplusimg.henrongyi.top/head.png';comment:'用户头像'"`
	Authority   Authority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:'用户角色'"`
	AuthorityId string    `json:"authorityId" gorm:"default:888;comment:'用户角色ID'"`
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

func (u *User) ChangePassword(db *gorm.DB, newPassword string) error {
	var user User
	err := db.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", util.EncodeMD5(newPassword)).Error
	return err
}

func (u *User) CountUser(db *gorm.DB) (int, error) {
	var count int
	if err := db.Model(&u).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (u *User) List(db *gorm.DB, pageOffset, pageSize int) ([]*User, error) {
	var users []*User
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if err = db.Where("is_del = ?", 0).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) Updates(db *gorm.DB)(*User,error) {
	err := db.Updates(u).Error
	return u,err
}
func (u *User) SetUserAuthority(db *gorm.DB,  values interface{}) error{
	return db.Model(&u).Where("uuid = ? and is_del = ?", u.UUID, 0).Updates(values).Error
}