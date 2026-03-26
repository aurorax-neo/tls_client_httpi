package tls_client

import (
	"io"
	"net/http"

	"github.com/aurorax-neo/tls_client_httpi"
	fhttp "github.com/bogdanfinn/fhttp"
	tlsClient "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

type TlsClient struct {
	Client    tlsClient.HttpClient
	ReqBefore handler
}

type handler func(req *fhttp.Request) error

func NewClientOptions(timeOutSec int, profile profiles.ClientProfile) []tlsClient.HttpClientOption {
	return []tlsClient.HttpClientOption{
		tlsClient.WithCookieJar(tlsClient.NewCookieJar()),
		tlsClient.WithTimeoutSeconds(timeOutSec),
		tlsClient.WithClientProfile(profile),
	}
}

func NewClient(options []tlsClient.HttpClientOption) *TlsClient {
	client, err := tlsClient.NewHttpClient(tlsClient.NewNoopLogger(), options...)
	if err != nil {
		panic(err)
	}
	return &TlsClient{Client: client}
}

func DefaultClient() *TlsClient {
	options := NewClientOptions(30, profiles.Chrome_124)
	return NewClient(options)
}

func convertResponse(resp *fhttp.Response) *http.Response {
	response := &http.Response{
		Status:           resp.Status,
		StatusCode:       resp.StatusCode,
		Proto:            resp.Proto,
		ProtoMajor:       resp.ProtoMajor,
		ProtoMinor:       resp.ProtoMinor,
		Header:           http.Header(resp.Header),
		Body:             resp.Body,
		ContentLength:    resp.ContentLength,
		TransferEncoding: resp.TransferEncoding,
		Close:            resp.Close,
		Uncompressed:     resp.Uncompressed,
		Trailer:          http.Header(resp.Trailer),
	}
	return response
}

func (TC *TlsClient) handleHeaders(req *fhttp.Request, headers tls_client_httpi.Headers) {
	if headers == nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func (TC *TlsClient) handleCookies(req *fhttp.Request, cookies tls_client_httpi.Cookies) {
	if cookies == nil {
		return
	}
	for _, c := range cookies {
		req.AddCookie(&fhttp.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			Expires:    c.Expires,
			RawExpires: c.RawExpires,
			MaxAge:     c.MaxAge,
			Secure:     c.Secure,
			HttpOnly:   c.HttpOnly,
			SameSite:   fhttp.SameSite(c.SameSite),
			Raw:        c.Raw,
			Unparsed:   c.Unparsed,
		})
	}
}

func (TC *TlsClient) Request(method tls_client_httpi.Method, rawURL string, headers tls_client_httpi.Headers, cookies tls_client_httpi.Cookies, body io.Reader) (*http.Response, error) {
	req, err := fhttp.NewRequest(string(method), rawURL, body)
	if err != nil {
		return nil, err
	}
	TC.handleHeaders(req, headers)
	TC.handleCookies(req, cookies)
	if TC.ReqBefore != nil {
		if err := TC.ReqBefore(req); err != nil {
			return nil, err
		}
	}
	do, err := TC.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return convertResponse(do), nil
}

func (TC *TlsClient) SetProxy(rawUrl string) error {
	return TC.Client.SetProxy(rawUrl)
}

func (TC *TlsClient) GetProxy() string {
	return TC.Client.GetProxy()
}

func (TC *TlsClient) SetFollowRedirect(followRedirect bool) {
	TC.Client.SetFollowRedirect(followRedirect)
}

func (TC *TlsClient) GetFollowRedirect() bool {
	return TC.Client.GetFollowRedirect()
}
