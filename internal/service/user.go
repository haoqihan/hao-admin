package service

import (
	uuid "github.com/satori/go.uuid"
	"hao-admin/internal/model"
	"hao-admin/pkg/app"
)

type RegisterRequest struct {
	Username    string `json:"username" form:"username" binding:"required"`
	NickName    string `json:"nick_name" form:"nick_name" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	AuthorityId string `json:"authority_id" form:"authority_id" binding:"required"`
	HeaderImage string `json:"header_image" form:"header_image"`
}

type LoginRequest struct {
	CaptchaId string `json:"captcha_id" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type DeleteUserRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid" binding:"required"`
	AuthorityId string    `json:"authority_id" binding:"required"`
}

func (svc *Service) Register(params *RegisterRequest) (error, *model.User) {
	return svc.dao.Register(params.Username, params.NickName, params.Password, params.HeaderImage, params.AuthorityId)
}

func (svc *Service) Login(params *LoginRequest) (error, *model.User) {
	return svc.dao.Login(params.Username, params.Password)
}

func (svc *Service) DeleteUser(params *DeleteUserRequest) error {
	return svc.dao.DeleteUser(params.ID)
}
func (svc *Service) ChangePassword(params *ChangePasswordRequest) error {
	return svc.dao.ChangePassword(params.Username, params.Password, params.NewPassword)
}

func (svc *Service) CountUser() (int, error) {
	return svc.dao.CountUser()
}

func (svc *Service) GetUserList(pager *app.Pager) ([]*model.User, error) {
	return svc.dao.GetUserList(pager.Page, pager.PageSize)

}
func (svc *Service) SetUserInfo(user *model.User) (*model.User, error) {
	return svc.dao.SetUserInfo(user)
}

func (svc *Service) SetUserAuthority(params *SetUserAuth) error {
	return svc.dao.SetUserAuthority(params.UUID, params.AuthorityId)
}
