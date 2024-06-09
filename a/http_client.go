package a

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

var httpRestyClient *resty.Client

func HttpClient() *resty.Client {
	if httpRestyClient == nil {
		httpRestyClient = resty.New().EnableTrace()
		httpRestyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: false})
		httpRestyClient.SetTimeout(5 * time.Minute)
	}
	return httpRestyClient
}
func HttpRequest() *resty.Request {
	return HttpClient().R()
}

func HttpGet(url string) []byte {
	resp, err := HttpClient().R().EnableTrace().Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp.Body()
}
