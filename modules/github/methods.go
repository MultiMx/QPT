package github

import (
	"fmt"
	"github.com/Mmx233/tool"
	"github.com/MultiMx/QPT/util"
)

func GetLatestRelease(owner, repo string) (string, error) {
	_, res, e := util.Http.Get(&tool.DoHttpReq{
		Url: fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo),
	})
	if e != nil {
		return "", e
	}
	return res["tag_name"].(string), e
}
