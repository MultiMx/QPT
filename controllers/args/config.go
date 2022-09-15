package args

import (
	"github.com/MultiMx/QPT/global"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
)

func configure() {
	switch os.Args[2] {
	case "proxy":
		_, e := url.Parse(os.Args[3])
		if e != nil {
			log.Errorln("url 不合法：", e)
			return
		}
		global.Config.Proxy = os.Args[3]
	default:
		log.Fatalln("配置键不存在")
	}

	if e := global.Config.Save(); e != nil {
		log.Errorln("保存配置失败：", e)
	}
}
