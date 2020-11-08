package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"hao-admin/global"
	"hao-admin/internal/service"
	"hao-admin/pkg/app"
	"hao-admin/pkg/errcode"
)

type Base struct{}

func NewBase() Base {
	return Base{}
}

var store = base64Captcha.DefaultMemStore
// 登录
func (b Base) Login(c *gin.Context) {
	var params service.LoginRequest
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errors)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}
	if !store.Verify(params.CaptchaId, params.Captcha, true) {
		global.Logger.Errorf("验证码错误")
		response.ToErrorResponse(errcode.ErrorCaptchaErr)
		return
	}
	svc := service.New(c.Request.Context())
	err, user := svc.Login(&params)
	if err != nil {
		global.Logger.Errorf("svc.Login err:%v", err)
		response.ToErrorResponse(errcode.ErrorLoginFail)
		return
	}
	token, err := app.GenerateToken(user.UUID,user.ID,user.NickName,user.Password,user.AuthorityId)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{"token": token})
	return
}

// 生成验证码
func (b Base) Captcha(c *gin.Context) {
	// 生成默认数字的driver
	response := app.NewResponse(c)
	driver := base64Captcha.NewDriverDigit(global.CaptchaSetting.ImgHeight, global.CaptchaSetting.ImgWidth, global.CaptchaSetting.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global.Logger.Errorf("cp.Generate err: %v", err)
		response.ToErrorResponse(errcode.ErrorCaptchaFail)
		return
	}
	response.ToResponse(gin.H{"CaptchaId": id, "picPath": b64s})
	return
}

// 登录之后签发jwt
