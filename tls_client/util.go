package tls_client

import (
	"fmt"

	fhttp "github.com/bogdanfinn/fhttp"
	"github.com/bogdanfinn/fhttp/httputil"
)

// OutFHttpRequest 打印请求.
func OutFHttpRequest(req *fhttp.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
	} else {
		fmt.Println(string(dump))
	}
}

// OutFHttpResponse 打印响应.
func OutFHttpResponse(res *fhttp.Response) {
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("Error dumping response:", err)
	} else {
		fmt.Println(string(dump))
	}
}
