package global

import (
	"github.com/Mmx233/config"
	"github.com/MultiMx/QPT/global/models"
	log "github.com/sirupsen/logrus"
)

var Config models.Config

func init() {
	Config.Init(&config.Options{
		Path:      "/etc/qpt.yaml",
		Config:    &Config,
		Default:   &models.Config{},
		Overwrite: true,
	})
	if e := Config.Load(); e != nil && e != config.IsNewConfig {
		log.Fatalln("载入配置失败：", e)
	}
}
