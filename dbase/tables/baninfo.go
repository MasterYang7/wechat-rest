package tables

// 群聊配置

type BanInfo struct {
	Rd        uint   `json:"rd" gorm:"primaryKey"`      // 主键
	Roomid    string `json:"roomid" gorm:"uniqueIndex"` // 群聊 id
	Name      string `json:"name"`                      // 群聊名称
	Num       uint   `json:"num"`                       // 警告次数
	Sender    string `json:"sender"`                    // 对象
	Ban       int64  `json:"ban"`                       // 是否黑名单 1黑名单 2 正常
	CreatedAt int64  `json:"created_at"`                // 创建时间戳
	UpdatedAt int64  `json:"updated_at"`                // 最后更新时间戳
}
