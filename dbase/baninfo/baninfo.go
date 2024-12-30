package baninfo

import (
	"time"

	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建黑名单

type CreateParam struct {
	Rd     uint   `json:"rd" binding:"required"`
	Roomid string `json:"roomid" `
	Name   string `json:"name"`
	Num    uint   `json:"num"`    // 警告次数
	Sender string `json:"sender"` // 对象
	Ban    int64  `json:"ban"`    // 是否黑名单 1黑名单 2 正常
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.BanInfo{
		Roomid: data.Roomid,
		Name:   data.Name,
		Num:    data.Num,
		Sender: data.Sender,
		Ban:    data.Ban,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新群聊

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.BanInfo{
			Roomid: data.Roomid,
			Sender: data.Sender,
		}).
		Updates(tables.BanInfo{
			Name: data.Name,
			Num:  data.Num,
			Ban:  data.Ban,
		})

	return result.Error

}

// 合并群聊

type ReplaceParam = CreateParam

func Replace(data *ReplaceParam) error {

	rq := &FetchParam{Rd: data.Rd}
	if rq.Rd == 0 {
		rq.Roomid = data.Roomid
	}

	item, err := Fetch(rq)
	if err == nil && item.Rd > 0 {
		data.Rd = item.Rd
		err = Update(data)
	} else {
		_, err = Create(data)
	}

	return err

}

// 获取群聊

type FetchParam struct {
	Rd     uint   `json:"rd"`
	Roomid string `json:"roomid"`
	Sender string `json:"sender"` // 对象
}

func Fetch(data *FetchParam) (*tables.BanInfo, error) {

	var item *tables.BanInfo

	result := dborm.Db.
		Where(&tables.BanInfo{
			Rd:     data.Rd,
			Roomid: data.Roomid,
		}).
		First(&item)

	if item == nil {
		item = &tables.BanInfo{Roomid: data.Roomid}
	}

	return item, result.Error

}
func FetchOne(data *FetchParam) (*tables.BanInfo, error) {

	var item *tables.BanInfo

	result := dborm.Db.
		Where(&tables.BanInfo{
			Sender: data.Sender,
			Roomid: data.Roomid,
		}).
		First(&item)

	if item == nil {
		item = &tables.BanInfo{Roomid: data.Roomid}
	}

	return item, result.Error

}

// 删除群聊

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	var item *tables.BanInfo

	result := dborm.Db.
		Where(&tables.BanInfo{
			Rd:     data.Rd,
			Roomid: data.Roomid,
		}).
		Delete(&item)

	return result.Error

}

// 获取群聊列表

type FetchAllParam struct {
	Level int32 `json:"level"`
}

func FetchAll(data *FetchAllParam) ([]*tables.BanInfo, error) {

	var items []*tables.BanInfo

	result := dborm.Db.
		Find(&items)

	return items, result.Error

}
func CloseOne() {
	roomMap := make(map[string]*tables.Chatroom)
	list, _ := FetchAll(nil)
	for _, v := range list {
		if roomMap[v.Roomid] == nil {
			room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: v.Roomid})
			if room.BanNum == 0 {
				room.BanNum = 10
			}
			roomMap[v.Roomid] = room
		}
		if int64(v.Num) > roomMap[v.Roomid].BanNum/2 && v.UpdatedAt+48*3600 > time.Now().Unix() {
			v.Num--
			dborm.Db.Where(&tables.BanInfo{Rd: v.Rd}).Updates(v)
		}

	}

}

// 获取群聊总数

type CountParam = FetchAllParam
