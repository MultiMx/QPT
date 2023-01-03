package golang

import (
	"github.com/MultiMx/QPT/controllers/install/golang"
	log "github.com/sirupsen/logrus"
)

func MakeUpdate() {
	vl, e := golang.GetLocalVersion()
	if e != nil {
		log.Errorln("获取本地 golang 版本失败：", e)
		return
	}
	v, e := golang.GetLatestVersion()
	if e != nil {
		log.Errorln("获取最新版本失败：", e)
		return
	}

	if v == vl {
		log.Infoln("已经是最新版")
		return
	}
	log.Infoln("检测到新版 ", v)

	if e = golang.Install(v); e != nil {
		log.Errorln("升级异常退出：", e)
		return
	}

	log.Infoln("已升级至 ", v)
}
