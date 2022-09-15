package golang

import (
	"errors"
	"fmt"
	"github.com/Mmx233/tool"
	"github.com/MultiMx/QPT/util"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"io"
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

func Download(v string) (io.ReadCloser, error) {
	res, e := util.Http.GetRequest(&tool.DoHttpReq{
		Url: fmt.Sprintf("https://go.dev/dl/%s.%s-%s.tar.gz", v, runtime.GOOS, runtime.GOARCH),
	})
	if e != nil {
		return nil, e
	}
	return res.Body, nil
}

func SaveTmpPack(i io.ReadCloser) (string, error) {
	defer i.Close()

	file, e := os.CreateTemp("", "go-install-*****.tar.gz")
	if e != nil {
		return "", e
	}
	defer file.Close()

	_, e = io.Copy(file, i)
	return file.Name(), e
}

func InstallFiles(file string) error {
	e := exec.Command("tar", "-C", "/usr/local", "-xzf", file).Wait()
	if e != nil {
		return e
	}
	return os.Symlink("/usr/local/go/bin/*", "/bin/")
}

func Install() error {
	v, e := GetLatestVersion()
	if e != nil {
		log.Errorln("获取最新版本失败")
		return e
	}
	res, e := Download(v)
	if e != nil {
		log.Errorln("下载最新版本失败")
		return e
	}
	file, e := SaveTmpPack(res)
	if e != nil {
		return e
	}
	defer os.Remove(file)

	return InstallFiles(file)
}
