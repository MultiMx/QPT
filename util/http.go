package util

import (
	"github.com/Mmx233/tool"
	"github.com/MultiMx/QPT/global"
	"net/http"
	"net/url"
	"time"
)

var Http *tool.Http

func init() {
	defaultTimeout := time.Second * 30

	transport := tool.GenHttpTransport(&tool.HttpTransportOptions{
		Timeout: defaultTimeout,
	})
	if global.Config.Proxy == "" {
		transport.Proxy = http.ProxyFromEnvironment
	} else {
		u, _ := url.Parse(global.Config.Proxy)
		transport.Proxy = http.ProxyURL(u)
	}
	Http = tool.NewHttpTool(tool.GenHttpClient(&tool.HttpClientOptions{
		Transport: transport,
		Timeout:   defaultTimeout,
	}))
}
