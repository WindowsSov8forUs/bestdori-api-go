package dto

type LoginCampaignsAll1Info struct {
	Caption []*string `json:"caption"`
}

type LoginCampaignsAll1 map[string]LoginCampaignsAll1Info

type LoginCampaignsAll5Info struct {
	LoginCampaignsAll1Info `json:",inline"`
	LoginBonusType         string    `json:"loginBonusType"`
	AssetBundleName        []*string `json:"assetBundleName"`
	PublishedAt            []*string `json:"publishedAt"`
	ClosedAt               []*string `json:"closedAt"`
}

type LoginCampaignsAll5 map[string]LoginCampaignsAll5Info

type LoginCampaignDetail struct {
	LoginBonusId int    `json:"loginBonusId"`
	Days         int    `json:"days"`
	ResourceType string `json:"resourceType"`
	ResourceId   int    `json:"resourceId"`
	Quantity     int    `json:"quantity"`
	VoiceId      *int   `json:"voiceId,omitempty"`
	Seq          int    `json:"seq"`
	GrantType    string `json:"grantType"`
}

type LoginCampaignInfo struct {
	LoginCampaignsAll5Info `json:",inline"`
	AssetMap               map[string]any           `json:"assetMap"`
	Details                []*[]LoginCampaignDetail `json:"details"`
}
