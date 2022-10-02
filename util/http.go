package util

import (
	"github.com/Mmx233/tool"
	"github.com/MultiMx/QPT/global"
	"net/http"
	"net/url"
	"time"
)

var Http *tool.Http

var HttpDownload *tool.Http

func init() {
	defaultTimeout := time.Second * 30
	downloadTimeout := time.Minute * 10

	var proxy func(r *http.Request) (*url.URL, error)
	if global.Config.Proxy == "" {
		proxy = http.ProxyFromEnvironment
	} else {
		u, _ := url.Parse(global.Config.Proxy)
		proxy = http.ProxyURL(u)
	}

	defaultTransport := tool.GenHttpTransport(&tool.HttpTransportOptions{
		Timeout: defaultTimeout,
	})
	defaultTransport.Proxy = proxy
	Http = tool.NewHttpTool(tool.GenHttpClient(&tool.HttpClientOptions{
		Transport: defaultTransport,
		Timeout:   defaultTimeout,
	}))

	downloadTransport := tool.GenHttpTransport(&tool.HttpTransportOptions{
		Timeout: downloadTimeout,
	})
	downloadTransport.Proxy = proxy
	HttpDownload = tool.NewHttpTool(tool.GenHttpClient(&tool.HttpClientOptions{
		Transport: downloadTransport,
		Timeout:   downloadTimeout,
	}))
}
