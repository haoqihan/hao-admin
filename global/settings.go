package global

import (
	"hao-admin/pkg/logger"
	"hao-admin/pkg/settings"
)

var (
	ServerSetting   *settings.ServerSettingS
	AppSetting      *settings.AppSettingS
	DataBaseSetting *settings.DataBaseSettingS
	JWTSetting      *settings.JWTSettingS
	EmailSetting    *settings.EmailSettingS
	Logger          *logger.Logger
	CaptchaSetting  *settings.CaptchaSettingS
)
