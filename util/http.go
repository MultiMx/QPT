package util

import (
	"github.com/Mmx233/tool"
	"net/http"
	"time"
)

var Http *tool.Http

func init() {
	defaultTimeout := time.Second * 30

	transport := tool.GenHttpTransport(&tool.HttpTransportOptions{
		Timeout: defaultTimeout,
	})
	transport.Proxy = http.ProxyFromEnvironment
	Http = tool.NewHttpTool(tool.GenHttpClient(&tool.HttpClientOptions{
		Transport: transport,
		Timeout:   defaultTimeout,
	}))
}
