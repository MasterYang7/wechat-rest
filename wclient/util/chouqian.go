package util

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var QianMap = make(map[string]struct {
	Qian string
	Jie  string
	Time time.Time
})

type CQ struct {
}

func (cq CQ) Chouqian(userId string) string {
	url := "https://api.t1qq.com/api/tool/cq?key=nYVgEm0QNGpzBnKU4RygwlnMw7"
	heard := map[string]string{"Content-Type": "application/x-www-form-urlencoded;charset:utf-8;"}
	resstu := struct {
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
		Title string `json:"title"`
		Qian  string `json:"qian"`
		Jie   string `json:"jie"`
		Time  string `json:"time"`
	}{}

	reslt, err := ProxySendRes("GET", url, "", heard)
	if err != nil {
		return "操作频繁，请稍后再试"
	}
	err = json.Unmarshal(reslt, &resstu)
	if err != nil {
		return err.Error()
	}
	qian := struct {
		Qian string
		Jie  string
		Time time.Time
	}{Qian: resstu.Qian, Jie: resstu.Jie, Time: time.Now()}
	QianMap[userId] = qian
	model := `
您抽到%s
----------------
🌌%s`
	go Clear()
	return fmt.Sprintf(model, resstu.Title, resstu.Qian)
}
func (cq CQ) Jieqian(userId string) string {
	model := `
----------------
🥁%s`
	if su, ok := QianMap[userId]; ok {
		delete(QianMap, userId)
		return fmt.Sprintf(model, su.Jie)
	} else {
		return "您未抽签喔（10分钟过期）"
	}
}

func Clear() {
	for k, v := range QianMap {
		if v.Time.Add(10 * time.Minute).Before(time.Now()) {
			delete(QianMap, k)
		}
	}
}

func ProxySendRes(sendType string, url string, body string, header map[string]string) (result []byte, err error) {
	if err != nil {
		return result, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	var reqest *http.Request
	if sendType == "GET" || sendType == "DELETE" {
		reqest, _ = http.NewRequest(sendType, url, nil)
	} else if sendType == "POST" || sendType == "PUT" {
		reqest, _ = http.NewRequest(sendType, url, strings.NewReader(body))
	}
	if header == nil {
		header = make(map[string]string)
	}
	header["Accept"] = "application/json, text/plain, */*"
	header["Accept-Language"] = "ja,zh-CN;q=0.8,zh;q=0.6"
	header["Connection"] = "keep-alive"
	for k, v := range header {
		reqest.Header.Set(k, v)
	}
	response, errc := client.Do(reqest)
	if errc != nil {
		return nil, errc
	} else {
		defer response.Body.Close()
		body, erra := ioutil.ReadAll(response.Body)
		if response.StatusCode >= 400 {
			err = errors.New(string(body))
			return
		}
		return body, erra
	}
}
