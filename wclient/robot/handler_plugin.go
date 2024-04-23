package robot

import (
	"strings"

	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient/util"
)

func pluginHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    -1,
		Order:    11,
		Roomid:   "+",
		Command:  "/pubg",
		Describe: "查询游戏信息",
		Callback: func(msg *wcferry.WxMsg) string {
			if msg.Content == "" {
				return "请在指令后输入游戏ID"
			}
			pubgsvc := util.PUBG{}
			id := strings.Split(msg.Content, "|")
			season := "29"
			if len(id) == 2 {
				season = id[1]
			}
			return pubgsvc.GetPlayerRank(id[0], season)
		},
	}, &Handler{
		Level:    0,
		Order:    12,
		Roomid:   "+",
		Command:  "抽签",
		Describe: "抽取",
		Callback: func(msg *wcferry.WxMsg) string {

			cqsvc := util.CQ{}

			return cqsvc.Chouqian(msg.Sender)
		},
	}, &Handler{
		Level:    0,
		Order:    12,
		Roomid:   "+",
		Command:  "解签",
		Describe: "解签",
		Callback: func(msg *wcferry.WxMsg) string {
			cqsvc := util.CQ{}
			return cqsvc.Jieqian(msg.Sender)
		},
	})

	return cmds

}
