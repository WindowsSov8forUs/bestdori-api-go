package dto

type UserPosterCard struct {
	Id         int  `json:"id"`
	Offset     int  `json:"offset"`
	TrainedArt bool `json:"trainedArt"`
}

type UserServerId struct {
	Id     int    `json:"id"`
	Server Server `json:"server"`
}

type UserInfo struct {
	Result          bool              `json:"result"`
	FollowingCount  int               `json:"followingCount"`
	FollowedByCount int               `json:"followedByCount"`
	Followed        bool              `json:"followed"`
	Nickname        string            `json:"nickname"`
	Titles          []PostAuthorTitle `json:"titles"`
	PosterCard      UserPosterCard    `json:"posterCard"`
	SelfIntro       string            `json:"selfIntro"`
	ServerIds       []UserServerId    `json:"serverIds"`
	SocialMedia     string            `json:"socialMedia"`
	FavCharaters    []int             `json:"favCharacters"`
	FavCards        []int             `json:"favCards"`
	FavBands        []int             `json:"favBands"`
	FavSongs        []int             `json:"favSongs"`
	FavCostumes     []int             `json:"favCostumes"`
}

type MeInfo struct {
	Result       bool              `json:"result"`
	Username     string            `json:"username"`
	Nickname     string            `json:"nickname"`
	Titles       []PostAuthorTitle `json:"titles"`
	Email        string            `json:"email"`
	MessageCount int               `json:"messageCount"`
}
