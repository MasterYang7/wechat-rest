package robot

import (
	"fmt"
	"time"

	"github.com/opentdp/wrest-chat/dbase/baninfo"
	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/wcferry"
)

func roomHandler() []*Handler {

	cmds := []*Handler{}

	rooms, err := chatroom.FetchAll(&chatroom.FetchAllParam{})

	if err == nil && len(rooms) > 0 {
		for k, v := range rooms {
			if len(v.JoinArgot) < 2 {
				continue
			}
			v := v // copy
			cmdkey := v.JoinArgot
			cmds = append(cmds, &Handler{
				Level:    -1,
				Order:    510 + int32(k),
				Roomid:   "-",
				Command:  cmdkey,
				Describe: v.Name,
				Callback: func(msg *wcferry.WxMsg) string {

					baninfo, _ := baninfo.FetchOne(&baninfo.FetchParam{Roomid: v.Roomid, Sender: msg.Sender})

					room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: v.Roomid})
					if baninfo.Ban == 1 {
						return "黑名单用户，无法使用此功能"
					}
					if room.BanNum == 0 {
						room.BanNum = 10
					}
					fmt.Println("违规测试", baninfo, baninfo.UpdatedAt+24*3600, time.Now().Unix())
					if baninfo.Num > uint(room.BanNum) && baninfo.UpdatedAt+24*3600 > time.Now().Unix() {
						return fmt.Sprintf("违规用户，%d分钟内无法进该群", (baninfo.UpdatedAt+24*3600-time.Now().Unix())/60)
					}

					resp := wc.CmdClient.InviteChatroomMembers(v.Roomid, msg.Sender)
					if resp == 1 {
						return "已发送群邀请，稍后请点击进入"
					} else {
						return "发送群邀请失败"
					}
				},
			})
		}
	}

	return cmds

}
