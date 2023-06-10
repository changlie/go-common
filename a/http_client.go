package a

import (
	"github.com/go-resty/resty/v2"
	"log"
)

var httpRestyClient *resty.Client

func httpClient() *resty.Client {
	if httpRestyClient == nil {
		httpRestyClient = resty.New()
	}
	return httpRestyClient
}

func HttpGet(url string) []byte {
	resp, err := httpClient().R().EnableTrace().Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp.Body()
}
