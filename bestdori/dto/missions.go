package dto

type MissionsAll5Info struct {
	Type    string    `json:"type"`
	StartAt []*string `json:"startAt"`
	EndAt   []*string `json:"endAt"`
	Title   []*string `json:"title"`
}

type MissionsAll5 map[string]MissionsAll5Info

type MissionDetailReward struct {
	MissionId       *int   `json:"missionId,omitempty"`
	Seq             *int   `json:"seq,omitempty"`
	MissionRewardId *int   `json:"missionRewardId,omitempty"`
	ResourceType    string `json:"resourceType"`
	ResourceId      int    `json:"resourceId"`
	Quantity        int    `json:"quantity"`
}

type MissionDetail struct {
	Seq         int                 `json:"seq"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	MaxProgress int                 `json:"maxProgress"`
	Reward      MissionDetailReward `json:"reward"`
}

type MissionInfo struct {
	MissionsAll5Info `json:",inline"`
	Details          []*[]MissionDetail `json:"details"`
}
