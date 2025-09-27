package dto

type EventArchiveInfo struct {
	Cutoff []map[string]int `json:"cutoff"`
	Board  [][]int          `json:"board"`
}

type EventArchiveAll5 map[string]EventArchiveInfo

type EventsAll1Info struct {
	EventName []*string `json:"eventName"`
}

type EventsAll1 map[string]EventsAll1Info

type EventsAll3Info struct {
	EventsAll1Info        `json:",inline"`
	EventType             string    `json:"eventType"`
	AssetBundleName       string    `json:"assetBundleName"`
	BannerAssetBundleName string    `json:"bannerAssetBundleName"`
	StartAt               []*string `json:"startAt"`
	EndAt                 []*string `json:"endAt"`
}

type EventsAll3 map[string]EventsAll3Info

type EventsAll4Info struct {
	EventsAll3Info `json:",inline"`
	RewardCards    []int `json:"rewardCards"`
}

type EventsAll4 map[string]EventsAll4Info

type EventAttribute struct {
	EventId   *int          `json:"eventId,omitempty"`
	Attribute CardAttribute `json:"attribute"`
	Percent   int           `json:"percent"`
}

type EventCharacter struct {
	EventId     *int `json:"eventId,omitempty"`
	CharacterId int  `json:"characterId"`
	Percent     int  `json:"percent"`
	Seq         *int `json:"seq,omitempty"`
}

type EventAttributeAndCharacterBonus struct {
	EventId          *int `json:"eventId,omitempty"`
	PointPercent     int  `json:"pointPercent"`
	ParameterPercent int  `json:"parameterPercent"`
}

type EventCharacterParameterBonus struct {
	EventId     *int `json:"eventId,omitempty"`
	Performance int  `json:"performance"`
	Technique   int  `json:"technique"`
	Visual      int  `json:"visual"`
}

type EventMember struct {
	EventId     int `json:"eventId"`
	SituationId int `json:"situationId"`
	Percent     int `json:"percent"`
	Seq         int `json:"seq"`
}

type EventLimitBreak struct {
	Rarity  int     `json:"rarity"`
	Rank    int     `json:"rank"`
	Percent float64 `json:"percent"`
}

type EventsAll5Info struct {
	EventsAll4Info                  `json:",inline"`
	Attributes                      []EventAttribute                 `json:"attributes"`
	Characters                      []EventCharacter                 `json:"characters"`
	EventAttributeAndCharacterBonus *EventAttributeAndCharacterBonus `json:"eventAttributeAndCharacterBonus,omitempty"`
	EventCharacterParameterBonus    *EventCharacterParameterBonus    `json:"eventCharacterParameterBonus,omitempty"`
	Members                         []EventMember                    `json:"members"`
	LimitBreaks                     []EventLimitBreak                `json:"limitBreaks"`
}

type EventsAll5 map[string]EventsAll5Info

type EventsAll6Info EventsAll5Info

type EventsAll6 map[string]EventsAll6Info

type EventPointReward struct {
	Point          string `json:"point"`
	RewardType     string `json:"rewardType"`
	RewardId       *int   `json:"rewardId,omitempty"`
	RewardQuantity int    `json:"rewardQuantity"`
}

type EventRankingReward struct {
	FromRank       int    `json:"fromRank"`
	ToRank         int    `json:"toRank"`
	RewardType     string `json:"rewardType"`
	RewardId       int    `json:"rewardId"`
	RewardQuantity int    `json:"rewardQuantity"`
}

type EventStoryReward struct {
	RewardType     string `json:"rewardType"`
	RewardId       *int   `json:"rewardId,omitempty"`
	RewardQuantity int    `json:"rewardQuantity"`
}

type EventStory struct {
	ScenarioId        string             `json:"scenarioId"`
	CoverImage        string             `json:"coverImage"`
	BackgroundImage   string             `json:"backgroundImage"`
	ReleasePt         string             `json:"releasePt"`
	Rewards           []EventStoryReward `json:"rewards"`
	Caption           []*string          `json:"caption"`
	Title             []*string          `json:"title"`
	Synopsis          []*string          `json:"synopsis"`
	ReleaseConditions []*string          `json:"releaseConditions"`
}

type EventMusicRankingReward struct {
	FromRank     int    `json:"fromRank"`
	ToRank       int    `json:"toRank"`
	ResourceType string `json:"resourceType"`
	ResourceId   int    `json:"resourceId"`
	Quantity     int    `json:"quantity"`
}

type EventMusic struct {
	MusicId             int                       `json:"musicId"`
	MusicRankingRewards []EventMusicRankingReward `json:"musicRankingRewards"`
}

type EventInfo struct {
	EventsAll6Info      `json:",inline"`
	EnableFlag          []bool                  `json:"enableFlag"`
	PublicStartAt       []*string               `json:"publicStartAt"`
	PublicEndAt         []*string               `json:"publicEndAt"`
	DistributionStartAt []*string               `json:"distributionStartAt"`
	DistributionEndAt   []*string               `json:"distributionEndAt"`
	BGMAssetBundleName  string                  `json:"bgmAssetBundleName"`
	BGMFileName         string                  `json:"bgmFileName"`
	AggregateEndAt      []*string               `json:"aggregateEndAt"`
	ExchangeEndAt       []*string               `json:"exchangeEndAt"`
	PointRewards        []*[]EventPointReward   `json:"pointRewards"`
	RankingRewards      []*[]EventRankingReward `json:"rankingRewards"`
	Stories             []EventStory            `json:"stories"`
	Musics              []*[]EventMusic         `json:"musics"`
}

type EventTopPoint struct {
	Time  float64 `json:"time"`
	UID   int     `json:"uid"`
	Value int     `json:"value"`
}

type EventTopUser struct {
	UID          int    `json:"uid"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	Rank         int    `json:"rank"`
	SID          int    `json:"sid"`
	STrained     bool   `json:"strained"`
	Degrees      []int  `json:"degrees"`
}

type EventTopData struct {
	Points []EventTopPoint `json:"points"`
	Users  []EventTopUser  `json:"users"`
}

type EventTrackerRate struct {
	Type   string   `json:"type"`
	Server Server   `json:"server"`
	Tier   int      `json:"tier"`
	Rate   *float64 `json:"rate,omitempty"`
}

type EventTrackerCutoff struct {
	Time float64
	Ep   int
}

type EventTrackerData struct {
	Result  bool                 `json:"result"`
	Cutoffs []EventTrackerCutoff `json:"cutoffs"`
}

type FestivalRotationMusic struct {
	MusicId int    `json:"musicId"`
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}

type FestivalStage struct {
	Type    string `json:"type"`
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}
