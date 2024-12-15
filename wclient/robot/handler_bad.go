package robot

import (
	"fmt"
	"regexp"

	"github.com/importcjj/sensitive"
	cmap "github.com/orcaman/concurrent-map/v2"

	"github.com/opentdp/wrest-chat/dbase/baninfo"
	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/dbase/keyword"
	"github.com/opentdp/wrest-chat/dbase/profile"
	"github.com/opentdp/wrest-chat/dbase/tables"
	"github.com/opentdp/wrest-chat/wcferry"
)

// var badMember = map[string]int{}
var badFilter *sensitive.Filter

var roomMemberAlias = map[string]string{}

func badHandler() []*Handler {

	updateBadWordFilter()

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    310,
		Roomid:   "*",
		Command:  "/bad",
		Describe: "添加违禁词",
		Callback: func(msg *wcferry.WxMsg) string {
			_, err := keyword.Create(&keyword.CreateParam{
				Group: "badword", Roomid: prid(msg), Phrase: msg.Content, Level: 1,
			})
			if err == nil {
				badFilter.AddWord(msg.Content)
				return "违禁词添加成功"
			}
			return "违禁词已存在"
		},
		PreCheck: badPreCheck,
	})

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    311,
		Roomid:   "*",
		Command:  "/bad:rm",
		Describe: "删除违禁词",
		Callback: func(msg *wcferry.WxMsg) string {
			err := keyword.Delete(&keyword.DeleteParam{
				Group: "badword", Roomid: prid(msg), Phrase: msg.Content,
			})
			if err == nil {
				badFilter.DelWord(msg.Content)
				return "违禁词删除成功"
			}
			return "违禁词删除失败"
		},
	})

	return cmds

}

func badPreCheck(msg *wcferry.WxMsg) string {

	// 私聊豁免
	if !msg.IsGroup {
		return ""
	}

	// 管理豁免
	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
	if up.Level > 6 {
		return ""
	}

	// 清洗并查找
	expr := regexp.MustCompile("[[:space:]]|[\x00-\x1F]|[\u2000-\u22ff]")
	text := roomMemberName(msg.Sender, msg.Roomid) + msg.Content
	keys := badFilter.FindAll(expr.ReplaceAllString(text, ""))
	if len(keys) == 0 {
		return ""
	}

	// 判断违禁级别
	level := 0
	for _, k := range keys {
		v, _ := keyword.Fetch(&keyword.FetchParam{Group: "badword", Phrase: k})
		if v.Level > 0 && (v.Roomid == msg.Roomid || v.Roomid == "*" || v.Roomid == "+") {
			level += int(v.Level)
		}
	}

	// 等级违规积分
	if level > 0 {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		info := updateBan(msg, uint(level))
		if info.Num > uint(room.BanNum) {
			// defer delete(badMember, msg.Sender)
			defer wc.CmdClient.DelChatRoomMembers(msg.Roomid, msg.Sender)
			str := "违规累计 %d；送你离开，天涯之外你是否还在"
			return fmt.Sprintf(str, info.Num)
		}
		str := "违规风险 +%d，当前累计：%d，大于 %d 将被请出群聊"
		return fmt.Sprintf(str, level, info.Num, room.BanNum)
	}

	return ""

}

func updateBan(msg *wcferry.WxMsg, level uint) tables.BanInfo {
	info, ok := BaninfoMap.Get(msg.Roomid)
	if !ok {
		info = cmap.New[*tables.BanInfo]()
	}
	sender, oks := info.Get(msg.Sender)
	if !oks {
		sender = &tables.BanInfo{
			Roomid: msg.Roomid,
			Name:   msg.Roomid,
			Num:    level,
			Sender: msg.Sender,
			Ban:    2,
		}
		baninfo.Create(&baninfo.CreateParam{Roomid: msg.Roomid, Sender: msg.Sender, Num: sender.Num, Ban: sender.Ban})
	} else {
		sender.Num += level
		baninfo.Update(&baninfo.CreateParam{Roomid: msg.Roomid, Sender: msg.Sender, Num: sender.Num, Ban: sender.Ban})
	}
	info.Set(msg.Sender, sender)
	BaninfoMap.Set(msg.Roomid, info)
	return *sender

}

func roomMemberName(wxid, roomid string) string {

	k := fmt.Sprintf("%s@%s", wxid, roomid)

	if roomMemberAlias[k] == "" {
		roomMemberAlias[k] = wc.CmdClient.GetAliasInChatRoom(wxid, roomid)
	}

	return roomMemberAlias[k]

}

func updateBadWordFilter() {

	filter := sensitive.New()

	items, _ := keyword.FetchAll(&keyword.FetchAllParam{
		Group: "badword",
	})
	for _, v := range items {
		filter.AddWord(v.Phrase)
	}

	badFilter = filter

}
