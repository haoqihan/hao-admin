package service

type RegisterRequest struct {
	Username    string `json:"username" form:"username" binding:"required"`
	NickName    string `json:"nick_name" form:"nick_name" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	AuthorityId string `json:"authority_id" form:"authority_id" binding:"required"`
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
