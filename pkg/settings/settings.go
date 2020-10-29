package settings

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Settings struct {
	vp *viper.Viper
}

func (s *Settings) WatchSettingChange() {
	go func() {
		s.vp.WriteConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()

}

func NewSettings(configs ...string) (*Settings, error) {
	vp := viper.New()

	vp.SetConfigName("config") //设置配置文件的名称为config
	vp.AddConfigPath("configs") // 设置配置文件的相对路径为 configs
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml") // 设置配置文件的类型为 yaml
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Settings{vp}
	s.WatchSettingChange()
	return s, nil
}
