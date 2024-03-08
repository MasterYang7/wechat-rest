package robot

import (
	"strings"

	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/util"
)

func pluginHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    0,
		Order:    11,
		ChatAble: true,
		RoomAble: true,
		Command:  "/pubg",
		Describe: "查询游戏信息",
		Callback: func(msg *wcferry.WxMsg) string {
			if msg.Content == "" {
				return "请在指令后输入游戏ID"
			}
			pubgsvc := util.PUBG{}
			id := strings.Split(msg.Content, "|")
			season := "28"
			if len(id) == 2 {
				season = id[1]
			}
			return pubgsvc.GetPlayerRank(id[0], season)
		},
	})

	return cmds

}
