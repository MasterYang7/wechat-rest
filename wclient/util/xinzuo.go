package util

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/opentdp/wrest-chat/dbase/setting"
	cmap "github.com/orcaman/concurrent-map/v2"
)

const api = "https://apis.tianapi.com/star/index?key=5443580973e075b985b326333b4d4f3c&astro=%s"

var (
	BgDict cmap.ConcurrentMap[string, int]
)

func init() {
	BgDict = cmap.New[int]()
}

var AstroMap = map[string]string{"白羊": "aries", "金牛": "taurus", "双子": "gemini",
	"巨蟹": "cancer", "狮子": "leo", "处女": "virgo", "天秤": "libra", "天蝎": "scorpio", "射手": "sagittarius",
	"摩羯": "capricorn", "水瓶": "aquarius", "双鱼": "pisces"}

type Plugin struct{}
type Astro struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		List []struct {
			Type    string `json:"type"`
			Content string `json:"content"`
		} `json:"list"`
	} `json:"result"`
}

const astroModel = `查收您的今日运势

✨综合指数：%s
❤爱情指数：%s
👔工作指数：%s
💴财运指数：%s
⛑健康指数：%s
🤞幸运数字：%s

🔔%s🔔%s
`

func (Plugin) CheckAstro(userid, astro string) string {
	ln := len(astro)
	if ln == 6 {
		astro = AstroMap[astro]
	}
	if astro == "" {
		return ""
	}
	if num, ok := BgDict.Get(userid); !ok || num < setting.Astro {
		as := Astro{}
		reslt, err := ProxySendRes("GET", fmt.Sprintf(api, astro), "", nil)
		if err != nil {
			return "操作频繁，请稍后再试"
		}
		err = json.Unmarshal(reslt, &as)
		if err != nil {
			return err.Error()
		}
		if as.Code != 200 {
			return "今日天机已消耗完了,请明日再来"
		}
		num++
		BgDict.Set(userid, num)
		data := as.Result.List
		return fmt.Sprintf(astroModel,
			a(data[0].Content),
			a(data[1].Content),
			a(data[2].Content),
			a(data[3].Content),
			a(data[4].Content),
			a(data[6].Content),
			data[8].Content, getDuJiTang())
	} else {
		return fmt.Sprintf("每人每日只能查询%d次喔", setting.Astro)
	}
}
func a(s string) string {
	return strings.ReplaceAll(s, "%", "")
}

func (Plugin) Clear() {
	BgDict.Clear()
}

func getDuJiTang() string {
	reslt, err := ProxySendRes("GET", "https://api.shadiao.pro/du", "", nil)
	if err != nil {
		return ""
	}
	data := struct {
		Data struct {
			Type string `json:"type"`
			Text string `json:"text"`
		}
	}{}
	err = json.Unmarshal(reslt, &data)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("\n送你一碗鸡汤:%s", data.Data.Text)

}
