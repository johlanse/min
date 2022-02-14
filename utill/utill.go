package utill

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

type Cookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func Parse(text string) (*goquery.Document, error) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(text))
	if err != nil {
		return document, err
	}
	return document, err
}

func Do(client *http.Client, url, method string, body url.Values, headers map[string]string, cookies []*http.Cookie) *http.Response {
	request, err := http.NewRequest(method, url, strings.NewReader(body.Encode()))
	if err != nil {
		panic(err)
	}
	if headers != nil {
		for name, value := range headers {
			request.Header.Add(name, value)
		}
	}
	if cookies != nil {
		for _, cookie := range cookies {
			request.AddCookie(cookie)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		log.Errorln("请求" + url + "出现错误" + err.Error())
		return response
	}

	return response
}

func LoadCookie(cookies []*http.Cookie) {
	var cos []Cookie
	for _, cookie := range cookies {
		var co Cookie
		co.Name = cookie.Name
		co.Value = cookie.Value
		cos = append(cos, co)
	}

	file, err := os.OpenFile("cookies.json", os.O_RDWR|os.O_CREATE, 666)
	if err != nil {
		log.Errorln("写入cookie出现错误" + err.Error())
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(&cos)
	if err != nil {
		log.Errorln("加载到cookies文件出现错误")
	}
}

func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
