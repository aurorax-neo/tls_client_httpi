package TCHUtil

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// OutHttpRequest 打印请求.
func OutHttpRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
	} else {
		fmt.Println(string(dump))
	}
}

// OutHttpResponse 打印响应.
func OutHttpResponse(res *http.Response) {
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("Error dumping response:", err)
	} else {
		fmt.Println(string(dump))
	}
}
