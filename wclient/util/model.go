package util

type PubgPlayerInfo struct {
	Data  PlayerData `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct{} `json:"meta"`
}
type PubgPlayerInfoList struct {
	Data  []PlayerData `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct{} `json:"meta"`
}

type PlayerData struct {
	Type          string                  `json:"type"`
	ID            string                  `json:"id"`
	Attributes    PlayerDataAttributes    `json:"attributes"`
	Relationships PlayerDataRelationships `json:"relationships"`
	Links         struct {
		Self   string `json:"self"`
		Schema string `json:"schema"`
	} `json:"links"`
}

type PlayerDataAttributes struct {
	TitleID      string      `json:"titleId"`
	ShardID      string      `json:"shardId"`
	PatchVersion string      `json:"patchVersion"`
	BanType      string      `json:"banType"`
	Name         string      `json:"name"`
	Stats        interface{} `json:"stats"`
}

type PlayerDataRelationships struct {
	Assets struct {
		Data []interface{} `json:"data"`
	} `json:"assets"`
	Matches PlayerDataRelationshipsMatch `json:"matches"`
}
type PlayerDataRelationshipsMatch struct {
	Data []PlayerDataRelationshipsMatchData `json:"data"`
}
type PlayerDataRelationshipsMatchData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type PubgPlayOut struct {
	BanType   string `json:"banType"`
	Name      string `json:"name"`
	AccountID string `json:"accountId"`
}

type Stats struct {
	DBNOs           int     `json:"DBNOs"`
	Assists         int     `json:"assists"`
	Boosts          int     `json:"boosts"`
	DamageDealt     float64 `json:"damageDealt"`
	DeathType       string  `json:"deathType"`
	HeadshotKills   int     `json:"headshotKills"`
	Heals           int     `json:"heals"`
	KillPlace       int     `json:"killPlace"`
	KillStreaks     int     `json:"killStreaks"`
	Kills           int     `json:"kills"`
	LongestKill     float64 `json:"longestKill"`
	Name            string  `json:"name"`
	PlayerID        string  `json:"playerId"`
	Revives         int     `json:"revives"`
	RideDistance    int     `json:"rideDistance"`
	RoadKills       int     `json:"roadKills"`
	SwimDistance    int     `json:"swimDistance"`
	TeamKills       int     `json:"teamKills"`
	TimeSurvived    int     `json:"timeSurvived"`
	VehicleDestroys int     `json:"vehicleDestroys"`
	WalkDistance    float64 `json:"walkDistance"`
	WeaponsAcquired int     `json:"weaponsAcquired"`
	WinPlace        int     `json:"winPlace"`
}

type Attributes struct {
	Stats         interface{} `json:"stats"`
	GameMode      string      `json:"gameMode"`
	TitleID       string      `json:"titleId"`
	Tags          interface{} `json:"tags"`
	MapName       string      `json:"mapName"`
	MatchType     string      `json:"matchType"`
	CreatedAt     string      `json:"createdAt"`
	Duration      int         `json:"duration"`
	ShardID       string      `json:"shardId"`
	IsCustomMatch bool        `json:"isCustomMatch"`
	SeasonState   string      `json:"seasonState"`
}

type Asset struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Roster struct {
	Type          string      `json:"type"`
	ID            string      `json:"id"`
	Won           string      `json:"won"`
	ShardID       string      `json:"shardId"`
	Stats         interface{} `json:"stats"`
	Relationships struct {
		Team         interface{} `json:"team"`
		Participants struct {
			Data []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"participants"`
	} `json:"relationships"`
}

type Participant struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		ShardID string `json:"shardId"`
		Stats   Stats  `json:"stats"`
		Actor   string `json:"actor"`
	} `json:"attributes"`
}

type Data struct {
	Type          string     `json:"type"`
	ID            string     `json:"id"`
	Attributes    Attributes `json:"attributes"`
	Relationships struct {
		Assets struct {
			Data []Asset `json:"data"`
		} `json:"assets"`
		Rosters struct {
			Data []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"rosters"`
	} `json:"relationships"`
	Links struct {
		Self   string `json:"self"`
		Schema string `json:"schema"`
	} `json:"links"`
}

type Included struct {
	Type          string      `json:"type"`
	ID            string      `json:"id"`
	Attributes    interface{} `json:"attributes"`
	Relationships struct {
		Team         interface{} `json:"team"`
		Participants struct {
			Data []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"participants"`
	} `json:"relationships"`
}

type Pubg struct {
	Data     Data       `json:"data"`
	Included []Included `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta interface{} `json:"meta"`
}

type PubgPlayerExp struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Tier               int    `json:"tier"`
			Level              int    `json:"level"`
			LastMatchID        string `json:"lastMatchId"`
			TotalMatchesPlayed int    `json:"totalMatchesPlayed"`
			XP                 int    `json:"xp"`
			Ban                string `json:"ban"`
		} `json:"attributes"`
	} `json:"data"`
}

type Season struct {
	ID     string
	Season string
}

type TPPData struct {
	DamageAvg  float64
	HaveData   int
	KD         float64
	KDA        float64
	RankPoint  float64
	Round      float64
	SeasonName string
}

type FPPData struct {
	DamageAvg  float64
	HaveData   int
	KD         float64
	KDA        float64
	RankPoint  float64
	Round      float64
	SeasonName string
}

type TppRankData struct {
	DamageAvg  float64
	HaveData   int
	KD         float64
	KDA        float64
	RankPoint  int
	Round      int
	SeasonName string
}

type FppRankData struct {
	DamageAvg  float64
	HaveData   int
	KD         float64
	KDA        float64
	RankPoint  int
	Round      int
	SeasonName string
}
