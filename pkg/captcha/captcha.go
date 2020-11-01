package captcha

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"hao-admin/global"
	"hao-admin/pkg/app"
	"hao-admin/pkg/errcode"
)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	response := app.NewResponse(c)
	driver := base64Captcha.NewDriverDigit(global.CaptchaSetting.ImgHeight,
		global.CaptchaSetting.ImgWidth,
		global.CaptchaSetting.KeyLong,
		0.7,
		80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, errs := cp.Generate()
	if errs != nil {
		global.Logger.Errorf("cp.Generate errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Error()))
		return
	} else {
		response.ToResponse(gin.H{
			"CaptchaId": id,
			"PicPath":   b64s,
		})
	}
	return
}
