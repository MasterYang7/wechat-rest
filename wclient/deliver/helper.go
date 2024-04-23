package deliver

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/wrest-chat/wclient"
	cmap "github.com/orcaman/concurrent-map/v2"
)

var RoomMemberMap = cmap.New[map[string]string]()

func Send(deliver, content string) error {

	content = strings.TrimSpace(content)
	delivers := strings.Split(deliver, "\n")

	for _, dr := range delivers {
		logman.Warn("deliver "+dr, "content", content)
		// 解析参数
		args := strings.Split(strings.TrimSpace(dr), ",")
		if len(args) < 2 {
			return errors.New("deliver is error")
		}
		// 分渠道投递
		switch args[0] {
		case "wechat":
			time.Sleep(1 * time.Second)
			wechatMessage(args[1:], content)
		}
	}

	return nil

}

func CheckOut(deliver, content string) error {
	content = strings.TrimSpace(content)
	model := strings.TrimPrefix(content, "CHECK_OUT")
	delivers := strings.Split(deliver, "\n")
	client := wclient.Register()
	for _, dr := range delivers {
		// 解析参数
		args := strings.Split(strings.TrimSpace(dr), ",")
		if len(args) < 2 {
			return errors.New("deliver is error")
		}
		// 分渠道投递
		switch args[0] {
		case "wechat":
			time.Sleep(1 * time.Second)
			roomid := args[1]

			data := client.CmdClient.GetChatRoomMembers(roomid)
			memberMap := make(map[string]string)
			for _, val := range data {
				memberMap[val.Wxid] = val.Name
			}
			if len(memberMap) == 0 {
				return nil
			}
			if old, ok := RoomMemberMap.Get(roomid); ok {
				tmp := old
				if len(memberMap) < len(old) {

					for key, _ := range old {
						if _, has := memberMap[key]; has {
							delete(tmp, key)
						}
					}
					for wxid, name := range tmp {
						text := fmt.Sprintf(model, name, name, wxid, time.Now().Format("2006-01-02 15:04:05"))
						logman.Warn("text "+roomid, "content", text)
						wechatMessage([]string{roomid}, text)

						time.Sleep(1 * time.Second)
					}
					RoomMemberMap.Set(roomid, memberMap)
				}

			} else {
				RoomMemberMap.Set(roomid, memberMap)
			}

		}
	}

	return nil
}
