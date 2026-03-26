package TCHI_test

import (
	"fmt"
	"testing"

	"github.com/aurorax-neo/tls_client_httpi/TCHUtil"
	"github.com/aurorax-neo/tls_client_httpi/tls_client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func TestGetReq(t *testing.T) {
	c := tls_client.NewClient(tls_client.NewClientOptions(30, profiles.Chrome_124))
	response, err := c.Request("GET", "https://tls.browserleaks.com/json", nil, nil, nil)
	if err != nil {
		return
	}
	TCHUtil.OutHttpResponse(response)
}

func TestGetProxy(t *testing.T) {
	c := tls_client.DefaultClient()
	c.SetProxy("http://127.0.0.1:7890")
	response, err := c.Request("GET", "https://ipv4.ip.sb", nil, nil, nil)
	if err != nil {
		return
	}

	fmt.Println("c")
	TCHUtil.OutHttpResponse(response)

}
