package request

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Request() {

}

// get 请求
func Get(url string, params map[string]string) (body []byte, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	body, _ = ioutil.ReadAll(resp.Body)
	return []byte(body), err
}

// post 请求
func Post(url string, params map[string]string, headers map[string]string) (body []byte, err error) {
	if len(headers) < 1 {
		headers = map[string]string{
			"Content-type": "Application/json",
		}
		newHeaders := map[string]string{}
		var (
			newKey string
		)
		for key, value := range headers {
			newKey = strings.ToLower(key)
			newHeaders[newKey] = value
		}
		headers = newHeaders
	}
	var defaultContentType string = "Application/json"
	if tmpV, ok := headers["content-type"]; ok {
		defaultContentType = tmpV
	}
	var buf = bytes.NewBuffer(make([]byte, 1024))
	resp, err := http.Post(url, defaultContentType, buf)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	body, _ = ioutil.ReadAll(resp.Body)
	return []byte(body), err
}
