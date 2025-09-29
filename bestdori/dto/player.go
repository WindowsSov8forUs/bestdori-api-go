package dto

type PlayerDataProfileMainDeckUserSituationsEntryUserAppendParameter struct {
	UserId                        string `json:"userId"`
	SituationId                   int    `json:"situationId"`
	Performance                   int    `json:"performance"`
	Technique                     int    `json:"technique"`
	Visual                        int    `json:"visual"`
	CharacterPotentialPerformance int    `json:"characterPotentialPerformance"`
	CharacterPotentialTechnique   int    `json:"characterPotentialTechnique"`
	CharacterPotentialVisual      int    `json:"characterPotentialVisual"`
	CharacterBonusPerformance     int    `json:"characterBonusPerformance"`
	CharacterBonusTechnique       int    `json:"characterBonusTechnique"`
	CharacterBonusVisual          int    `json:"characterBonusVisual"`
}

type PlayerDataProfileMainDeckUserSituationsEntry struct {
	UserId              string                                                          `json:"userId"`
	SituationId         int                                                             `json:"situationId"`
	Level               int                                                             `json:"level"`
	Exp                 int                                                             `json:"exp"`
	CreatedAt           string                                                          `json:"createdAt"`
	AddExp              int                                                             `json:"addExp"`
	TrainingStatus      string                                                          `json:"trainingStatus"`
	DuplicateCount      int                                                             `json:"duplicateCount"`
	Illust              string                                                          `json:"illust"`
	SkillExp            int                                                             `json:"skillExp"`
	SkillLevel          int                                                             `json:"skillLevel"`
	UserAppendParameter PlayerDataProfileMainDeckUserSituationsEntryUserAppendParameter `json:"userAppendParameter"`
	LimitBreakRank      int                                                             `json:"limitBreakRank"`
}

type PlayerDataProfileMainDeckUserSituations struct {
	Entries []PlayerDataProfileMainDeckUserSituationsEntry `json:"entries"`
}

type PlayerDataProfileEnabledUserAreaItemsEntry struct {
	UserId           string `json:"userId"`
	AreaItemId       int    `json:"areaItemId"`
	AreaItemCategory string `json:"areaItemCategory"`
	Level            int    `json:"level"`
}

type PlayerDataProfileEnabledUserAreaItems struct {
	Entries []PlayerDataProfileEnabledUserAreaItemsEntry `json:"entries"`
}

type PlayerDataProfileBandRankMap struct {
	Entries map[string]int `json:"entries"`
}

type PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicListEntry struct {
	MusicId    int                 `json:"musicId"`
	Difficulty ChartDifficultyName `json:"difficulty"`
	Rating     int                 `json:"rating"`
}

type PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicList struct {
	Entries []PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicListEntry `json:"entries"`
}

type PlayerDataProfileUserHighScoreRating struct {
	UserPoppinPartyHighScoreMusicList     PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicList `json:"userPoppinPartyHighScoreMusicList"`
	UserAfterglowHighScoreMusicList       PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicList `json:"userAfterglowHighScoreMusicList"`
	UserPastelPalettesHighScoreMusicList  PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicList `json:"userPastelPalettesHighScoreMusicList"`
	UserHelloHappyWorldHighScoreMusicList PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicList `json:"userHelloHappyWorldHighScoreMusicList"`
	UserRoseliaHighScoreMusicList         PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicList `json:"userRoseliaHighScoreMusicList"`
	UserMorfonicaHighScoreMusicList       PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicList `json:"userMorfonicaHighScoreMusicList"`
	UserRaiseASuilenHighScoreMusicList    PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicList `json:"userRaiseASuilenHighScoreMusicList"`
	UserMyGOHighScoreMusicList            PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicList `json:"userMyGOHighScoreMusicList"`
	UserOtherHighScoreMusicList           PlayerDataProfileUserHighScoreRatingUserBandHighScoreMusicList `json:"userOtherHighScoreMusicList"`
}

type PlayerDataProfileMainUserDeck struct {
	DeckId   int    `json:"deckId"`
	DeckName string `json:"deckName"`
	DeckType string `json:"deckType"`
	Leader   int    `json:"leader"`
	Member1  int    `json:"member1"`
	Member2  int    `json:"member2"`
	Member3  int    `json:"member3"`
	Member4  int    `json:"member4"`
}

type PlayerDataProfileUserProfileSituation struct {
	UserId                     string `json:"userId"`
	SituationId                int    `json:"situationId"`
	Illust                     string `json:"illust"`
	ViewProfileSituationStatus string `json:"viewProfileSituationStatus"`
}

type PlayerDataProfileUserProfileDegreeMapEntry struct {
	UserId            string `json:"userId"`
	ProfileDegreeType string `json:"profileDegreeType"`
	DegreeId          int    `json:"degreeId"`
}

type PlayerDataProfileUserProfileDegreeMap struct {
	Entries map[string]PlayerDataProfileUserProfileDegreeMapEntry `json:"entries"`
}

type PlayerDataProfileUserTwitter struct {
	TwitterId       string `json:"twitterId"`
	TwitterName     string `json:"twitterName"`
	ScreenName      string `json:"screenName"`
	Url             string `json:"url"`
	ProfileImageUrl string `json:"profileImageUrl"`
}

type PlayerDataProfileUserDeckTotalRatingMapEntry struct {
	Rank        string `json:"rank"`
	Score       int    `json:"score"`
	Level       int    `json:"level"`
	LowerRating int    `json:"lowerRating"`
	UpperRating int    `json:"upperRating"`
}

type PlayerDataProfileUserDeckTotalRatingMap struct {
	Entries map[string]PlayerDataProfileUserDeckTotalRatingMapEntry `json:"entries"`
}

type PlayerDataProfileStageChallengeAchievementConditionsMap struct {
	Entries map[string]int `json:"entries"`
}

type PlayerDataProfileUserMusicClearInfoMapEntry struct {
	ClearedMusicCount    int `json:"clearedMusicCount"`
	FullComboMusicCount  int `json:"fullComboMusicCount"`
	AllPerfectMusicCount int `json:"allPerfectMusicCount"`
}

type PlayerDataProfileUserMusicClearInfoMap struct {
	Entries map[ChartDifficultyName]PlayerDataProfileUserMusicClearInfoMapEntry `json:"entries"`
}

type PlayerDataProfileUserCharacterRankMapEntry struct {
	Rank                   int    `json:"rank"`
	Exp                    string `json:"exp"`
	AddExp                 string `json:"addExp"`
	NextExp                string `json:"nextExp"`
	TotalExp               string `json:"totalExp"`
	ReleasedPotentialLevel string `json:"releasedPotentialLevel"`
}

type PlayerDataProfileUserCharacterRankMap struct {
	Entries map[string]PlayerDataProfileUserCharacterRankMapEntry `json:"entries"`
}

type PlayerDataProfile struct {
	UserId                                        string                                                  `json:"userId"`
	UserName                                      string                                                  `json:"userName"`
	Rank                                          int                                                     `json:"rank"`
	Degree                                        int                                                     `json:"degree"`
	Introduction                                  string                                                  `json:"introduction"`
	PublishTotalDeckPowerFlg                      bool                                                    `json:"publishTotalDeckPowerFlg"`
	PublishBandRankFlg                            bool                                                    `json:"publishBandRankFlg"`
	PublishMusicClearedFlg                        bool                                                    `json:"publishMusicClearedFlg"`
	PublishMusicFullComboFlg                      bool                                                    `json:"publishMusicFullComboFlg"`
	PublishMusicAllPerfectFlg                     bool                                                    `json:"publishMusicAllPerfectFlg"`
	PublishHighScoreRatingFlg                     bool                                                    `json:"publishHighScoreRatingFlg"`
	PublishUserIdFlg                              bool                                                    `json:"publishUserIdFlg"`
	PublishUpdatedAtFlg                           bool                                                    `json:"publishUpdatedAtFlg"`
	PublishDeckRankFlg                            bool                                                    `json:"publishDeckRankFlg"`
	PublishStageChallengeAchievementConditionsFlg bool                                                    `json:"publishStageChallengeAchievementConditionsFlg"`
	PublishStageChallengeFriendRankingFlg         bool                                                    `json:"publishStageChallengeFriendRankingFlg"`
	PublishCharacterRankFlg                       bool                                                    `json:"publishCharacterRankFlg"`
	SearchableFlg                                 bool                                                    `json:"searchableFlg"`
	FriendApplicableFlg                           bool                                                    `json:"friendApplicableFlg"`
	MainDeckUserSituations                        PlayerDataProfileMainDeckUserSituations                 `json:"mainDeckUserSituations"`
	EnabledUserAreaItems                          PlayerDataProfileEnabledUserAreaItems                   `json:"enabledUserAreaItems"`
	BandRankMap                                   PlayerDataProfileBandRankMap                            `json:"bandRankMap"`
	UserHighScoreRating                           PlayerDataProfileUserHighScoreRating                    `json:"userHighScoreRating"`
	MainUserDeck                                  PlayerDataProfileMainUserDeck                           `json:"mainUserDeck"`
	UserProfileSituation                          PlayerDataProfileUserProfileSituation                   `json:"userProfileSituation"`
	UserProfileDegreeMap                          PlayerDataProfileUserProfileDegreeMap                   `json:"userProfileDegreeMap"`
	UserTwitter                                   PlayerDataProfileUserTwitter                            `json:"userTwitter,omitempty"`
	UserDeckTotalRatingMap                        PlayerDataProfileUserDeckTotalRatingMap                 `json:"userDeckTotalRatingMap"`
	StageChallengeAchievementConditionsMap        PlayerDataProfileStageChallengeAchievementConditionsMap `json:"stageChallengeAchievementConditionsMap"`
	UserMusicClearInfoMap                         PlayerDataProfileUserMusicClearInfoMap                  `json:"userMusicClearInfoMap"`
	UserCharacterRankMap                          PlayerDataProfileUserCharacterRankMap                   `json:"userCharacterRankMap"`
}

type PlayerData struct {
	Cache   bool               `json:"cache"`
	Time    float64            `json:"time"`
	Profile *PlayerDataProfile `json:"profile,omitempty"`
}

type PlayerInfo struct {
	Result bool       `json:"result"`
	Data   PlayerData `json:"data"`
}
