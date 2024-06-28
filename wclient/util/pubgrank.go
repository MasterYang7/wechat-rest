package util

import "time"

type RankedPlayerStats struct {
	Data struct {
		Type       string `json:"type"` // 类型
		Attributes struct {
			RankedGameModeStats struct {
				Squad struct {
					CurrentTier       Tier    `json:"currentTier"`       // 当前段位
					CurrentRankPoint  int     `json:"currentRankPoint"`  // 当前积分
					BestTier          Tier    `json:"bestTier"`          // 最高段位
					BestRankPoint     int     `json:"bestRankPoint"`     // 最高积分
					RoundsPlayed      int     `json:"roundsPlayed"`      // 总局数
					AvgRank           float64 `json:"avgRank"`           // 平均排名
					AvgSurvivalTime   int     `json:"avgSurvivalTime"`   // 平均生存时间
					Top10Ratio        float64 `json:"top10Ratio"`        // 前10比例
					WinRatio          float64 `json:"winRatio"`          // 胜率
					Assists           int     `json:"assists"`           // 助攻数
					Wins              int     `json:"wins"`              // 胜利数
					KDA               float64 `json:"kda"`               // KDA
					KDR               int     `json:"kdr"`               // KDR
					Kills             int     `json:"kills"`             // 击杀数
					Deaths            int     `json:"deaths"`            // 死亡数
					RoundMostKills    int     `json:"roundMostKills"`    // 单局最高击杀数
					LongestKill       int     `json:"longestKill"`       // 最远击杀距离
					HeadshotKills     int     `json:"headshotKills"`     // 爆头击杀数
					HeadshotKillRatio int     `json:"headshotKillRatio"` // 爆头击杀比例
					DamageDealt       float64 `json:"damageDealt"`       // 造成伤害
					DBNOs             int     `json:"dBNOs"`             // 击倒数
					ReviveRatio       int     `json:"reviveRatio"`       // 复活比例
					Revives           int     `json:"revives"`           // 复活数
					Heals             int     `json:"heals"`             // 治疗数
					Boosts            int     `json:"boosts"`            // 加速道具数
					WeaponsAcquired   int     `json:"weaponsAcquired"`   // 获取武器数
					TeamKills         int     `json:"teamKills"`         // 团队击杀数
					PlayTime          int     `json:"playTime"`          // 游戏时间
					KillStreak        int     `json:"killStreak"`        // 连杀数
				} `json:"squad"` // 四人小队模式
				SquadFpp struct {
					CurrentTier       Tier    `json:"currentTier"`       // 当前段位
					CurrentRankPoint  int     `json:"currentRankPoint"`  // 当前积分
					BestTier          Tier    `json:"bestTier"`          // 最高段位
					BestRankPoint     int     `json:"bestRankPoint"`     // 最高积分
					RoundsPlayed      int     `json:"roundsPlayed"`      // 总局数
					AvgRank           float64 `json:"avgRank"`           // 平均排名
					AvgSurvivalTime   int     `json:"avgSurvivalTime"`   // 平均生存时间
					Top10Ratio        float64 `json:"top10Ratio"`        // 前10比例
					WinRatio          float64 `json:"winRatio"`          // 胜率
					Assists           int     `json:"assists"`           // 助攻数
					Wins              int     `json:"wins"`              // 胜利数
					KDA               float64 `json:"kda"`               // KDA
					KDR               int     `json:"kdr"`               // KDR
					Kills             int     `json:"kills"`             // 击杀数
					Deaths            int     `json:"deaths"`            // 死亡数
					RoundMostKills    int     `json:"roundMostKills"`    // 单局最高击杀数
					LongestKill       int     `json:"longestKill"`       // 最远击杀距离
					HeadshotKills     int     `json:"headshotKills"`     // 爆头击杀数
					HeadshotKillRatio int     `json:"headshotKillRatio"` // 爆头击杀比例
					DamageDealt       float64 `json:"damageDealt"`       // 造成伤害
					DBNOs             int     `json:"dBNOs"`             // 击倒数
					ReviveRatio       int     `json:"reviveRatio"`       // 复活比例
					Revives           int     `json:"revives"`           // 复活数
					Heals             int     `json:"heals"`             // 治疗数
					Boosts            int     `json:"boosts"`            // 加速道具数
					WeaponsAcquired   int     `json:"weaponsAcquired"`   // 获取武器数
					TeamKills         int     `json:"teamKills"`         // 团队击杀数
					PlayTime          int     `json:"playTime"`          // 游戏时间
					KillStreak        int     `json:"killStreak"`        // 连杀数
				} `json:"squad-fpp"` // 四人小队FPP模式
			} `json:"rankedGameModeStats"` // 排位赛游戏模式统计
		} `json:"attributes"` // 属性
	} `json:"data"` // 数据
}

type Tier struct {
	Tier    string `json:"tier"`    // 段位
	SubTier string `json:"subTier"` // 子段位
}

// FmRankData 表示 fm_rank_data 表中的一条数据

type FmRankData struct {
	Id               int       `json:"id"`                 // ID
	Name             string    `json:"name"`               // 唯一ID
	AccountId        string    `json:"account_id"`         // 唯一ID
	DamageDealt      float64   `json:"damage_dealt"`       // 爆头率
	Season           string    `json:"season"`             // 数据
	SubTier          string    `json:"sub_tier"`           // 段位等级
	Tier             string    `json:"tier"`               // 当前段位
	CurrentRankPoint int       `json:"current_rank_point"` // 当前分数
	Kills            int       `json:"kills"`              // 击杀数
	Model            string    `json:"model"`              // 模式
	RoundsPlayed     int       `json:"rounds_played"`      // 比赛场次
	Kda              float64   `json:"kda"`                // KDA
	Kd               float64   `json:"kd"`                 // KD
	UpdateTime       time.Time `json:"update_time"`        // 更新时间
	CreateTime       time.Time `json:"create_time"`        // 创建时间
}
