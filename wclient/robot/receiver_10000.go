package robot

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/opentdp/wrest-chat/dbase/baninfo"
	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
)

// 处理系统消息
func receiver10000(msg *wcferry.WxMsg) {

	if msg.IsGroup {
		receiver10000Public(msg)
	} else {
		receiver10000Private(msg)
	}

}

func receiver10000Public(msg *wcferry.WxMsg) {

	// 自动回应群聊拍一拍（私聊不支持）
	if strings.Contains(msg.Content, "拍了拍我") {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if room.PatReturn == "true" {
			wc.CmdClient.SendPatMsg(msg.Roomid, msg.Sender)
		}
		return
	}

	// 邀请"xxx"加入了群聊
	r1 := regexp.MustCompile(`邀请"(.+)"加入了群聊`)
	if matches := r1.FindStringSubmatch(msg.Content); len(matches) > 1 {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		time.Sleep(3 * time.Second) // 延迟1秒
		list := wc.CmdClient.GetChatRoomMembers(msg.Roomid)
		for _, v := range list {
			fmt.Println(v.Name, v)
			if v.Name == matches[1] {
				msg.Sender = v.Wxid
				break
			}

		}
		fmt.Println("发送人", msg.Sender, "内容", msg.Content, msg.Roomid)
		isban, remsg := checkUserIsBan(msg.Roomid, msg.Sender, uint(room.BanNum))
		if isban {

			defer wc.CmdClient.DelChatRoomMembers(msg.Roomid, msg.Sender)
			reply(msg, remsg)
			return
		}

		if len(room.WelcomeMsg) > 1 {
			time.Sleep(1 * time.Second) // 延迟1秒
			if room.Roomid == "47697253318@chatroom" {
				avatar := wc.CmdClient.GetAvatar([]string{msg.Sender})
				avatarurl := "https://cache.jg110.cn/fm/1702083164"
				if len(avatar) > 0 {
					avatarurl = avatar[0].SmallHeadImgUrl
				}
				wc.CmdClient.SendRichText("👆点我查看FM联赛赛事手册", "", fmt.Sprintf("👏欢迎 %s 加入FM选手群🎇", matches[1]), "请第一时间修改群昵称为游戏ID，并阅读赛事手册和群规", "https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FurGeywGtGmFYW2lJ0i_52g", avatarurl, msg.Roomid)
			} else {

				reply(msg, room.WelcomeMsg)
			}
		}
		return
	}

	// "xxx"通过扫描"xxx"分享的二维码加入群聊
	r2 := regexp.MustCompile(`"(.+)"通过扫描"(.+)"分享的二维码加入群聊`)
	if matches := r2.FindStringSubmatch(msg.Content); len(matches) > 1 {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		time.Sleep(3 * time.Second) // 延迟1秒
		list := wc.CmdClient.GetChatRoomMembers(msg.Roomid)
		for _, v := range list {
			if v.Name == matches[1] {
				msg.Sender = v.Wxid
				break
			}
		}
		isban, remsg := checkUserIsBan(msg.Roomid, msg.Sender, uint(room.BanNum))
		if isban {

			defer wc.CmdClient.DelChatRoomMembers(msg.Roomid, msg.Sender)
			reply(msg, remsg)
			return
		}
		if len(room.WelcomeMsg) > 1 {
			time.Sleep(1 * time.Second) // 延迟1秒
			reply(msg, room.WelcomeMsg)
		}
		return
	}

}

func receiver10000Private(msg *wcferry.WxMsg) {

	// 接受好友后响应
	if strings.Contains(msg.Content, "现在可以开始聊天了") {
		if len(setting.FriendHello) > 1 {
			reply(msg, setting.FriendHello)
		}
		return
	}

}
func checkUserIsBan(Roomid, Sender string, bannum uint) (bool, string) {
	if bannum == 0 {
		bannum = 10
	}
	baninfo, err := baninfo.FetchOne(&baninfo.FetchParam{Roomid: Roomid, Sender: Sender})
	if err != nil {
		return false, "系统数据库异常，请联系管理员：MasterYang77"
	}
	fmt.Println("违规测试", baninfo, baninfo.UpdatedAt+24*3600, time.Now().Unix())
	if baninfo.Num > bannum && baninfo.UpdatedAt+24*3600 > time.Now().Unix() {
		return true, fmt.Sprintf("违规用户，%d分钟内无法进该群", (baninfo.UpdatedAt+24*3600-time.Now().Unix())/60)
	}
	return false, ""
}
