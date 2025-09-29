package dto

type SkillsAll2Info struct {
	SimpleDescription []*string `json:"simpleDescription"`
}

type SkillsAll2 map[string]SkillsAll2Info

type SkillsAll5Info struct {
	SkillsAll2Info `json:",inline"`
	Description    []*string `json:"description"`
	Duration       []float64 `json:"duration"`
}

type SkillsAll5 map[string]SkillsAll5Info

type SkillsAll10ActivationEffectTypesScore struct {
	ActivateEffectValue     []int  `json:"activateEffectValue"`
	ActivateEffectValueType string `json:"activateEffectValueType"`
	ActivateCondition       string `json:"activateCondition"`
}

type SkillsAll10ActivationEffectTypes struct {
	Score SkillsAll10ActivationEffectTypesScore `json:"score"`
}

type SkillsAll10ActivationEffect struct {
	ActivationEffectTypes SkillsAll10ActivationEffectTypes `json:"activationEffectTypes"`
}

type SkillsAll10Info struct {
	SkillsAll5Info   `json:",inline"`
	ActivationEffect SkillsAll10ActivationEffect `json:"activationEffect"`
}

type SkillsAll10 map[string]SkillsAll10Info
