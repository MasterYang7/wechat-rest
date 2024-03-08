package util

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func ProxySendPubg(sendType string, url string, body, token string, header map[string]string) (result []byte, err error) {
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
	header["Accept"] = "application/vnd.api+json"
	header["Accept-Language"] = "ja,zh-CN;q=0.8,zh;q=0.6"
	header["Connection"] = "keep-alive"
	header["Authorization"] = "Bearer " + token
	for k, v := range header {
		reqest.Header.Set(k, v)
	}
	response, errc := client.Do(reqest)
	if errc != nil {
		return nil, errc
	} else {
		defer response.Body.Close()
		body, erra := ioutil.ReadAll(response.Body)

		if response.StatusCode == 429 {
			err = errors.New(string(body))
			return
		}
		return body, erra
	}
}

func round(value float64, places int) float64 {
	rounding := 1
	for i := 0; i < places; i++ {
		rounding *= 10
	}
	return float64(int(value*float64(rounding))) / float64(rounding)
}

func rankLevel(lv string) string {
	switch lv {
	case "Master":
		return "大师"
	case "Diamond":
		return "钻石"
	case "Platinum":
		return "铂金"
	case "Gold":
		return "黄金"
	case "Silver":
		return "白银"
	case "Bronze":
		return "青铜"
	default:
		return "废铁"
	}
}
func calculateLevel(score int) string {

	if score < 1000 {
		return "Ⅴ"
	} else if score >= 3500 {
		return ""
	} else {
		tmp := score - 1000
		tmp2 := tmp % 500
		if tmp2 < 100 {
			return "Ⅴ"
		}
		if tmp2 < 200 && tmp2 >= 100 {
			return "Ⅳ"
		}
		if tmp2 < 300 && tmp2 >= 200 {
			return "Ⅲ"
		}
		if tmp2 < 400 && tmp2 >= 300 {
			return "Ⅱ"
		}
		if tmp2 < 500 && tmp2 >= 400 {
			return "Ⅰ"
		}
	}

	return ""
}
func calculateLevel2(score string, scores int) string {
	if scores >= 3500 {
		return ""
	}
	switch score {
	case "1":
		return "Ⅰ"
	case "2":
		return "Ⅱ"
	case "3":
		return "Ⅲ"
	case "4":
		return "Ⅳ"
	case "5":
		return "Ⅴ"
	default:
		return "废铁"
	}

}

func normalizeKD(kd float64) float64 {
	return kd / 5.0
}

func normalizeRankingScore(score int) float64 {
	return float64(score) / 5000.0
}
func normalizeRankingDMA(score float64) float64 {
	return score / 1000.0
}

func calculatePlayerLevel(kd float64, score int, dma float64) string {

	kdNormalized := normalizeKD(kd)
	rankingScoreNormalized := normalizeRankingScore(score)
	dmaSS := normalizeRankingDMA(dma)
	totla := (kdNormalized)*0.5 + (dmaSS)*0.3 + (rankingScoreNormalized)*0.2
	if totla >= 0.8 {
		return "AMW"
	} else if totla >= 0.6 {
		return "MK14"
	} else if totla >= 0.4 {
		return "Groza"
	} else if totla >= 0.2 {
		return "Beryl M762"
	} else {
		return "平底锅"
	}
}
