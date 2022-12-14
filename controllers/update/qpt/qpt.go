package qpt

import (
	"fmt"
	"github.com/Mmx233/tool"
	"github.com/MultiMx/QPT/global"
	"github.com/MultiMx/QPT/modules/github"
	"github.com/MultiMx/QPT/util"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
)

const owner = "MultiMx"
const repo = "QPT"

func GetLatestVersion() (string, error) {
	return github.GetLatestRelease(owner, repo)
}

func Download(v string) (string, error) {
	spin := util.Spinner("正在下载")
	spin.Start()
	defer spin.Stop()

	res, e := util.HttpDownload.GetRequest(&tool.DoHttpReq{
		Url: fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/qpt_%s_%s", owner, repo, v, runtime.GOOS, runtime.GOARCH),
	})
	if e != nil {
		return "", e
	}
	return util.SaveTmpPack(res.Body)
}

func Install(file string) error {
	_, e := exec.Command("chmod", "+x", file).Output()
	if e != nil {
		return e
	}
	path, e := os.Executable()
	if e != nil {
		return e
	}
	_, e = exec.Command("mv", file, path).Output()
	return e
}

func MakeUpdate() {
	v, e := GetLatestVersion()
	if e != nil {
		log.Errorln("获取最新版信息失败：", e)
		return
	}

	if v == global.VERSION {
		log.Infoln("已经是最新版")
		return
	}

	log.Infoln("检测到新版本：", v)

	file, e := Download(v)
	if e != nil {
		log.Errorln("下载发行版本失败：", e)
		return
	}
	defer os.Remove(file)
	log.Infoln("发行版本下载完毕")

	if e = Install(file); e != nil {
		log.Errorln("安装失败：", e)
		return
	}

	log.Infoln("已升级至 QPT ", v)
}
