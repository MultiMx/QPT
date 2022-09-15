package golang

import (
	"errors"
	"fmt"
	"github.com/Mmx233/tool"
	"github.com/MultiMx/QPT/util"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetLocalVersion() (string, error) {
	d, e := exec.Command("go", "version").Output()
	if e != nil {
		return "", e
	}
	return strings.Split(string(d), " ")[2], nil
}

func GetLatestVersion() (string, error) {
	res, e := util.Http.GetRequest(&tool.DoHttpReq{
		Url: "https://go.dev/dl/",
	})
	if e != nil {
		return "", e
	}
	defer res.Body.Close()

	doc, e := goquery.NewDocumentFromReader(res.Body)
	if e != nil {
		return "", e
	}
	version, ok := doc.Find("#stable + .toggleVisible").First().Attr("id")
	if !ok {
		return "", errors.New("数据解析失败")
	}

	return version, nil
}

func Download(v string) (string, error) {
	spin := util.Spinner()
	spin.Suffix = "正在下载"
	spin.Start()
	defer spin.Stop()

	res, e := util.Http.GetRequest(&tool.DoHttpReq{
		Url: fmt.Sprintf("https://go.dev/dl/%s.%s-%s.tar.gz", v, runtime.GOOS, runtime.GOARCH),
	})
	if e != nil {
		return "", e
	}
	return util.SaveTmpPack(res.Body)
}

func InstallFiles(file string) error {
	spin := util.Spinner()
	spin.Suffix = "正在解压"
	spin.Start()
	defer spin.Stop()

	_, e := exec.Command("tar", "-C", "/usr/local", "-xzf", file).Output()
	return e
}

func LinkFiles() error {
	return os.Symlink("/usr/local/go/bin/*", "/bin/")
}

func Install(v string) error {
	file, e := Download(v)
	if e != nil {
		log.Errorln("下载最新版本失败")
		return e
	}
	defer os.Remove(file)
	log.Infoln("安装包下载完毕")

	e = InstallFiles(file)
	if e != nil {
		return e
	}
	return LinkFiles()
}

func MakeInstall() {
	vl, e := GetLocalVersion()
	if e == nil {
		log.Errorf("本地已存在 golang %s\n", vl)
		return
	}
	v, e := GetLatestVersion()
	if e != nil {
		log.Errorln("获取最新版本失败：", e)
		return
	}

	log.Infoln("开始安装 ", v)

	if e = Install(v); e != nil {
		log.Errorln("安装异常退出：", e)
	}
}
