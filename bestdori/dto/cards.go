package dto

type CardRarity int

const (
	CardRarity1 CardRarity = iota + 1 // ★
	CardRarity2                       // ★★
	CardRarity3                       // ★★★
	CardRarity4                       // ★★★★
	CardRarity5                       // ★★★★★
)

type CardAttribute string

const (
	CardAttributePowerful CardAttribute = "powerful"
	CardAttributeHappy    CardAttribute = "happy"
	CardAttributePure     CardAttribute = "pure"
	CardAttributeCool     CardAttribute = "cool"
)

type CardsAll2Info struct {
	CharacterId int           `json:"characterId"` // 角色 ID
	Attribute   CardAttribute `json:"attribute"`   // 卡牌属性
}

type CardsAll2 map[string]CardsAll2Info

type CardsAll3Info struct {
	CardsAll2Info `json:",inline"`
	Prefix        []*string `json:"prefix"` // 卡牌名称 定长列表
}

type CardsAll3 map[string]CardsAll3Info

type cardStatInfo struct {
	Performance int `json:"performance"` // 卡牌演出值
	Technique   int `json:"technique"`   // 卡牌技巧值
	Visual      int `json:"visual"`      // 卡牌形象值
}

type cardStatTraining struct {
	cardStatInfo `json:",inline"`
	LevelLimit   int `json:"levelLimit"` // 特训等级上限
}

type cardStat struct {
	Levels   map[string]cardStatInfo // 卡牌各等级下数据值。在 all.5.json 中只会包括最低等级与最高等级，在 CardInfo 中则会包括所有等级。
	Episodes *[]cardStatInfo         `json:"episodes,omitempty"` // 卡牌故事阅读后增加数据值。数量为卡牌所有的卡牌故事数量，若没有卡牌故事则不存在该字段。
	Training *cardStatTraining       `json:"training,omitempty"` // 卡牌特训增加数据值。若卡牌无特训则不存在该字段。
}

type CardsAll5Info struct {
	CardsAll3Info   `json:",inline"`
	Rarity          CardRarity `json:"rarity"`          // 卡牌稀有度
	LevelLimit      int        `json:"levelLimit"`      // 卡牌等级上限
	ResourceSetName string     `json:"resourceSetName"` // 卡牌资源所在集合名。提取卡面等资源时需要提供。
	ReleasedAt      []*string  `json:"releasedAt"`      // 卡牌上线时间戳 定长列表
	SkillId         int        `json:"skillId"`         // 技能 ID
	Type            string     `json:"type"`            // 卡牌类型
	Stat            cardStat   `json:"stat"`            // 卡牌数据信息
}

type CardsAll5 map[string]CardsAll5Info

type cardEpisodesEntryCostsEntry struct {
	ResourceId   int    `json:"resourceId"`
	ResourceType string `json:"resourceType"`
	Quantity     int    `json:"quantity"`
	LbBonus      int    `json:"lbBonus"`
}

type cardEpisodesEntryCosts struct {
	Entries []cardEpisodesEntryCostsEntry `json:"entries"`
}

type cardEpisodesEntryRewardsEntry struct {
	ResourceType string `json:"resourceType"`
	Quantity     int    `json:"quantity"`
	LbBonus      int    `json:"lbBonus"`
}

type cardEpisodesEntryRewards struct {
	Entries []cardEpisodesEntryRewardsEntry `json:"entries"`
}

type cardEpisodesEntry struct {
	EpisodeId         int                      `json:"episodeId"`   // 故事 ID
	EpisodeType       string                   `json:"episodeType"` // 故事类型
	SituationId       int                      `json:"situationId"`
	ScenarioId        string                   `json:"scenarioId"`        // 场景 ID
	AppendPerformance int                      `json:"appendPerformance"` // 增加演出值
	AppendTechnique   int                      `json:"appendTechnique"`   // 增加技巧值
	AppendVisual      int                      `json:"appendVisual"`      // 增加形象值
	ReleaseLevel      int                      `json:"releaseLevel"`      // 开放等级
	Costs             cardEpisodesEntryCosts   `json:"costs"`             // 故事解锁所需道具
	Rewards           cardEpisodesEntryRewards `json:"rewards"`           // 故事奖励
	Title             []*string                `json:"title"`             // 故事标题 定长列表
	CharacterId       int                      `json:"characterId"`       // 故事角色 ID
}

type cardEpisodes struct {
	Entries []cardEpisodesEntry `json:"entries"`
}

type cardSourceGacha struct {
	Probability float64 `json:"probability"` // 招募抽取概率
}

type cardSource struct {
	Gacha map[string]cardSourceGacha `json:"gacha,omitempty"` // 招募 ID 与抽取概率字典
}

type CardInfo struct {
	CardsAll5Info  `json:",inline"`
	SdResourceName string       `json:"sdResourceName"` // 卡牌 LIVE 服装资源名
	Episodes       cardEpisodes `json:"episodes"`       // 卡牌故事
	CostumeId      int          `json:"costumeId"`      // 卡牌服装 ID
	GachaText      []*string    `json:"gachaText"`      // 卡牌上线招募名 定长列表
	SkillName      []*string    `json:"skillName"`      // 卡牌技能名 定长列表
	Source         []cardSource `json:"source"`         // 卡牌招募信息 定长列表
}

type CardTrain string

const (
	CardTrainNormal        CardTrain = "normal"         // 特训前
	CardTrainAfterTraining CardTrain = "after_training" // 特训后
)

type StarType string

const (
	StarTypeStar        StarType = "star"         // 特训前星标
	StarTypeStarTrained StarType = "star_trained" // 特训后星标
)
