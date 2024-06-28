package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/opentdp/go-helper/logman"
)

var (
	lock sync.RWMutex
)

type PUBG struct{}

const model = `ID:%s  赛季:第%s赛季
TPP.KD:%.2f  段位:%s  场伤:%.2f 场次:%d
鉴定为：%s
FPP.KD:%.2f  段位:%s  场伤:%.2f 场次:%d
鉴定为:%s
账号状态:%s
`

func (p *PUBG) GetPlayerRank(name, seasonId string) string {
	userInfo, err := p.UserInfo(name)
	if err != nil {
		return err.Error()
	}
	if userInfo.BanType == "永久封禁" {
		return fmt.Sprintf("ID：%s  永久封禁数据封存！", name)
	}
	accountID := userInfo.AccountID
	seasonStr := ""
	if seasonId == "" {
		seasonStr = fmt.Sprintf("division.bro.official.pc-2018-%d", 28)
	} else {
		seasonStr = fmt.Sprintf("division.bro.official.pc-2018-%s", seasonId)
	}
	// https://api.pubg.com/shards/steam/players/account.7d61a1cda14c45da8beacc63804833d3/seasons/division.bro.official.pc-2018-27/ranked
	rankData := RankedPlayerStats{}

	key := APP_KEY.GetNext()
	url := fmt.Sprintf("https://api.pubg.com/shards/steam/players/%s/seasons/%s/ranked", accountID, seasonStr)
	reslt, err := ProxySendPubg("GET", url, "", key, nil)
	if err != nil {
		logman.Warn("UserInfo ", err.Error())
		return ""
	}
	tppData := FmRankData{}
	fppData := FmRankData{}
	json.Unmarshal(reslt, &rankData)

	if rankData.Data.Attributes.RankedGameModeStats.Squad.RoundsPlayed != 0 {
		roundsPlayed := rankData.Data.Attributes.RankedGameModeStats.Squad.RoundsPlayed
		damageAvg := round(rankData.Data.Attributes.RankedGameModeStats.Squad.DamageDealt/float64(roundsPlayed), 2)
		kd := round(float64(rankData.Data.Attributes.RankedGameModeStats.Squad.Kills)/float64(roundsPlayed), 2)
		kda := round(rankData.Data.Attributes.RankedGameModeStats.Squad.KDA, 2)
		tppData.AccountId = accountID
		tppData.Kd = kd
		tppData.Kda = kda
		tppData.CurrentRankPoint = rankData.Data.Attributes.RankedGameModeStats.Squad.CurrentRankPoint
		tppData.Kills = rankData.Data.Attributes.RankedGameModeStats.Squad.Kills
		tppData.Model = "tpp"
		tppData.RoundsPlayed = rankData.Data.Attributes.RankedGameModeStats.Squad.RoundsPlayed
		tppData.Tier = rankData.Data.Attributes.RankedGameModeStats.Squad.CurrentTier.Tier
		tppData.SubTier = rankData.Data.Attributes.RankedGameModeStats.Squad.CurrentTier.SubTier
		tppData.DamageDealt = damageAvg
		tppData.UpdateTime = time.Now()
		tppData.Season = seasonId

	}

	if rankData.Data.Attributes.RankedGameModeStats.SquadFpp.RoundsPlayed != 0 {

		roundsPlayed := rankData.Data.Attributes.RankedGameModeStats.SquadFpp.RoundsPlayed
		damageAvg := round(rankData.Data.Attributes.RankedGameModeStats.SquadFpp.DamageDealt/float64(roundsPlayed), 2)
		kd := round(float64(rankData.Data.Attributes.RankedGameModeStats.SquadFpp.Kills)/float64(roundsPlayed), 2)
		kda := round(rankData.Data.Attributes.RankedGameModeStats.SquadFpp.KDA, 2)
		fppData.Kd = kd
		fppData.Kda = kda
		fppData.AccountId = accountID

		fppData.CurrentRankPoint = rankData.Data.Attributes.RankedGameModeStats.SquadFpp.CurrentRankPoint
		fppData.Kills = rankData.Data.Attributes.RankedGameModeStats.SquadFpp.Kills
		fppData.Model = "fpp"
		fppData.RoundsPlayed = rankData.Data.Attributes.RankedGameModeStats.SquadFpp.RoundsPlayed
		fppData.Tier = rankData.Data.Attributes.RankedGameModeStats.SquadFpp.CurrentTier.Tier
		fppData.SubTier = rankData.Data.Attributes.RankedGameModeStats.SquadFpp.CurrentTier.SubTier
		fppData.UpdateTime = time.Now()
		fppData.DamageDealt = damageAvg
		fppData.Season = seasonId

	}

	result := fmt.Sprintf(model, name, seasonId, tppData.Kd,
		fmt.Sprintf("%s%s", rankLevel(tppData.Tier), calculateLevel2(tppData.SubTier, tppData.CurrentRankPoint)),
		tppData.DamageDealt, tppData.RoundsPlayed,
		calculatePlayerLevel(tppData.Kd, tppData.CurrentRankPoint, fppData.DamageDealt), fppData.Kd,
		fmt.Sprintf("%s%s", rankLevel(fppData.Tier), calculateLevel2(fppData.SubTier, fppData.CurrentRankPoint)),
		fppData.DamageDealt, fppData.RoundsPlayed,
		calculatePlayerLevel(fppData.Kd, fppData.CurrentRankPoint, fppData.DamageDealt),
		userInfo.BanType,
	)
	if name == "BHGS_Naruto" {
		result += `
		他是FM明星选手喔~1元一张签名照`
	}
	return result
}

func (p *PUBG) UserInfo(userName string) (PubgPlayOut, error) {
	res := PubgPlayOut{}
	url := "https://api.pubg.com/shards/steam/players"
	userName = strings.TrimSpace(userName)
	url = fmt.Sprintf(url+"?filter[playerNames]=%s", strings.ReplaceAll(userName, " ", ""))
	key := APP_KEY.GetNext()
	reslt, err := ProxySendPubg("GET", url, "", key, nil)
	if err != nil {
		logman.Warn("UserInfo ", err.Error())
		return res, fmt.Errorf("操作频繁，请稍后再试")
	}
	if strings.Contains(string(reslt), "Not Found") {
		return res, fmt.Errorf("游戏ID不存在,请检查后重新输入")
	}

	pubgPlayerInfo := PubgPlayerInfoList{}
	json.Unmarshal(reslt, &pubgPlayerInfo)
	if len(pubgPlayerInfo.Data) > 0 {
		res.Name = pubgPlayerInfo.Data[0].Attributes.Name
		res.AccountID = pubgPlayerInfo.Data[0].ID
		res.BanType = pubgPlayerInfo.Data[0].Attributes.BanType
	}

	switch res.BanType {
	case "TemporaryBan":
		res.BanType = "异常检测"
	case "PermanentBan":
		res.BanType = "永久封禁"
	case "Innocent":
		res.BanType = "正常"
	}
	return res, nil
}
