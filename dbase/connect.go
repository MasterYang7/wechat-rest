package dbase

import (
	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/go-helper/filer"

	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

func Connect() {

	dbname := "wrest.db3"
	if !filer.Exists(dbname) {
		dbname = args.Web.Storage + "/" + dbname
	}

	// 连接数据库
	db := dborm.Connect(&dborm.Config{
		Type: "sqlite", DbName: dbname,
	})

	// 开启外键约束
	db.Exec("PRAGMA foreign_keys=ON;")

	// 实施自动迁移
	db.AutoMigrate(
		&tables.Chatroom{},
		&tables.BanInfo{},
		&tables.Cronjob{},
		&tables.Contact{},
		&tables.Keyword{},
		&tables.LLModel{},
		&tables.Message{},
		&tables.Profile{},
		&tables.Setting{},
		&tables.Webhook{},
	)

	// 加载全局配置
	setting.DataMigrate()
	setting.Laod()

}
