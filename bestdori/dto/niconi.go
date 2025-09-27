package dto

type LLSifDifficulty struct {
	LiveSettingId     int    `json:"live_setting_id"`
	Difficulty        int    `json:"difficulty"`
	StageLevel        int    `json:"stage_level"`
	NotesSettingAsset string `json:"notes_setting_asset"`
	SRankCombo        int    `json:"s_rank_combo"`
	Available         bool   `json:"available"`
	ACFlag            int    `json:"ac_flag"`
	SwingFlag         int    `json:"swing_flag"`
	FiveKeysFlag      int    `json:"five_keys_flag"`
}

type LLSifSongInfo struct {
	Name             string            `json:"name"`
	NameKana         string            `json:"name_kana"`
	Keyword          string            `json:"keyword"`
	LiveIconAsset    string            `json:"live_icon_asset"`
	SoundAsset       string            `json:"sound_asset"`
	AttributeIconId  int               `json:"attribute_icon_id"`
	LiveTime         float64           `json:"live_time"`
	MemberTag        string            `json:"member_tag"`
	UnitTypeId       int               `json:"unit_type_id,omitempty"`
	MemberFilterCond string            `json:"member_filter_cond"`
	MinId            int               `json:"min_id"`
	MaxId            int               `json:"max_id"`
	Difficulties     []LLSifDifficulty `json:"difficulties"`
}

type LLSifMisc map[string]LLSifSongInfo
