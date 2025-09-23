package dto

type CharactersAll2Info struct {
	CharacterType string    `json:"characterType"`       // 角色类型
	CharacterName []*string `json:"characterName"`       // 角色姓名 定长列表
	Nickname      []*string `json:"nickname"`            // 角色昵称 定长列表
	BandId        *int      `json:"bandId,omitempty"`    // 角色所在乐队 ID 。非主要角色不存在此项
	ColorCode     *string   `json:"colorCode,omitempty"` // 角色代表色十六进制码。非主要角色不存在此项
}

type CharactersAll2 map[string]CharactersAll2Info

type characterSeasonCostumeListMapEntrySeasonEntry struct {
	CharacterId           int    `json:"characterId"` // 角色 ID
	BasicSeasonId         int    `json:"basicSeasonId"`
	CostumeType           string `json:"costumeType"` // 服装类型
	SeasonCostumeType     string `json:"seasonCostumeType"`
	SdAssetBundleName     string `json:"sdAssetBundleName"`     // LIVESD 资源资源库名
	Live2DAssetBundleName string `json:"live2dAssetBundleName"` // Live2D 资源资源库名
	SeasonType            string `json:"seasonType"`            // 季 ID
}

type characterSeasonCostumeListMapEntrySeason struct {
	Entries []characterSeasonCostumeListMapEntrySeasonEntry `json:"entries"`
}

type characterSeasonCostumeListMap struct {
	Entries map[string]characterSeasonCostumeListMapEntrySeason `json:"entries"` // 季 ID 与服装列表映射字典
}

type CharactersAll5Info struct {
	CharactersAll2Info   `json:",inline"`
	FirstName            []*string                      `json:"firstName"`                      // 角色名 定长列表
	LastName             []*string                      `json:"lastName"`                       // 角色姓氏 定长列表
	SeasonCostumeListMap *characterSeasonCostumeListMap `json:"seasonCostumeListMap,omitempty"` // 角色某季服装列表映射表
}

type CharactersAll5 map[string]CharactersAll5Info

type CharactersMain1Info struct {
	CharacterType string `json:"characterType"` // 角色类型
	BandId        int    `json:"bandId"`        // 角色所在乐队 ID
}

type CharactersMain1 map[string]CharactersMain1Info

type CharactersMain2Info struct {
	CharactersMain1Info `json:",inline"`
	CharacterName       []*string `json:"characterName"` // 角色姓名 定长列表
	Nickname            []*string `json:"nickname"`      // 角色昵称 定长列表
	ColorCode           string    `json:"colorCode"`     // 角色代表色十六进制码
}

type CharactersMain2 map[string]CharactersMain2Info

type CharactersMain3Info struct {
	CharactersMain2Info `json:",inline"`
	FirstName           []*string `json:"firstName"` // 角色名 定长列表
	LastName            []*string `json:"lastName"`  // 角色姓氏 定长列表

}

type CharactersMain3 map[string]CharactersMain3Info

type characterProfile struct {
	CharacterVoice   []*string `json:"characterVoice"`   // 角色声优名 定长列表
	FavoriteFood     []*string `json:"favoriteFood"`     // 角色喜好食物 定长列表
	HatedFood        []*string `json:"hatedFood"`        // 角色厌恶食物 定长列表
	Hobby            []*string `json:"hobby"`            // 角色习惯 定长列表
	SelfIntroduction []*string `json:"selfIntroduction"` // 角色自我介绍 定长列表
	School           []*string `json:"school"`           // 角色学校名称 定长列表
	SchoolCls        []*string `json:"schoolCls"`        // 角色所在班级 定长列表
	SchoolYear       []*string `json:"schoolYear"`       // 角色所在学年 定长列表
	Part             string    `json:"part"`             // 角色担当
	Birthday         string    `json:"birthday"`         // 角色生日时间戳
	Constellation    string    `json:"constellation"`    // 角色星座名
	Height           string    `json:"height"`           // 角色身高 (cm)
}

type CharacterInfo struct {
	CharactersAll5Info `json:",inline"`
	SdAssetBundleName  string            `json:"sdAssetBundleName"`          // 角色 SD 资源资源库名
	DefaultCostumeId   *int              `json:"defaultCostumeId,omitempty"` // 默认服装 ID 。非主要角色不存在此项
	Ruby               []*string         `json:"ruby"`                       // 	角色读音注释 定长列表
	Profile            *characterProfile `json:"profile,omitempty"`          // 角色个人资料
}
