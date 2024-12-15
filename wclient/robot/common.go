package robot

import (
	"strings"

	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/wrest-chat/dbase/baninfo"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/dbase/tables"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient"
	cmap "github.com/orcaman/concurrent-map/v2"
)

var wc *wcferry.Client
var BaninfoMap cmap.ConcurrentMap[string, cmap.ConcurrentMap[string, *tables.BanInfo]]

func Start() {

	if !setting.BotEnable {
		logman.Warn("robot disabled")
		return
	}

	if wc != nil {
		logman.Warn("robot already started")
		return
	}

	wc = wclient.Register()
	_, err := wc.EnrollReceiver(true, receiver)
	if err != nil {
		logman.Fatal("robot start failed", "error", err)
	}
	initBan()

	ResetHandlers()

}

func Reset() {

	setting.Laod()
	ResetHandlers()

}

///////////////////////// COMMON METHODS /////////////////////////

// 会话场景
func prid(msg *wcferry.WxMsg) string {

	if msg.IsGroup {
		return msg.Roomid
	}
	return "-"

}

// 回复消息
func reply(msg *wcferry.WxMsg, text string) int32 {

	if msg.IsSelf {
		return -2
	}

	if text = strings.TrimSpace(text); text == "" {
		return -1
	}

	if msg.IsGroup {
		if msg.Sender != "" && wcferry.ContactType(msg.Sender) == "好友" {
			user := wc.CmdClient.GetInfoByWxid(msg.Sender)
			if user != nil && user.Name != "" {
				text = "@" + user.Name + "\n" + text
			}
		}
		return wc.CmdClient.SendTxt(text, msg.Roomid, msg.Sender)
	}

	return wc.CmdClient.SendTxt(text, msg.Sender, "")

}

// 获取违禁名单
func initBan() {
	BaninfoMap = cmap.New[cmap.ConcurrentMap[string, *tables.BanInfo]]()
	banlist, err := baninfo.FetchAll(nil)
	if err != nil {
		logman.Fatal("baninfo start failed", "error", err)
	}

	for _, v := range banlist {
		info, ok := BaninfoMap.Get(v.Roomid)
		if !ok {
			info = cmap.New[*tables.BanInfo]()
		}
		info.Set(v.Sender, v)
		BaninfoMap.Set(v.Roomid, info)
	}
}
