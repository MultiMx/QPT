package args

import (
	"github.com/MultiMx/QPT/controllers/install/golang"
	log "github.com/sirupsen/logrus"
	"os"
)

func installers() {
	if len(os.Args) < 3 {
		log.Fatalln("参数缺失")
	}
	switch os.Args[2] {
	case "go":
		v, e := golang.GetLocalVersion()
		if e == nil {
			log.Errorln("本地已存在 golang %s\n", v)
			return
		}
		if e = golang.Install(); e != nil {
			log.Errorln("安装异常退出：", e)
		}
	}
}
