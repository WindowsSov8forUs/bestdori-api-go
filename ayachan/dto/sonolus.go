package dto

type LevelItemResource struct {
	Hash string `json:"hash"`
	URL  string `json:"url"`
}

type LevelData struct {
	Name     string `json:"name"`
	Source   string `json:"source"`
	Version  int    `json:"version"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Author   string `json:"author"`
	Tags     []any  `json:"tags"`
}

type LevelItemEngineData struct {
	LevelData `json:",inline"`
	Thumbnail LevelItemResource `json:"thumbnail"`
	Data      LevelItemResource `json:"data"`
}

type LevelItemEngineSkin struct {
	LevelItemEngineData `json:",inline"`
	Texture             LevelItemResource `json:"texture"`
}

type LevelItemEngineBackground struct {
	LevelItemEngineData `json:",inline"`
	Image               LevelItemResource `json:"image"`
	Configuration       LevelItemResource `json:"configuration"`
}

type LevelItemEngineEffect struct {
	LevelItemEngineData `json:",inline"`
	Audio               LevelItemResource `json:"audio"`
}

type LevelItemEngineParticle struct {
	LevelItemEngineData `json:",inline"`
	Texture             LevelItemResource `json:"texture"`
}

type LevelItemEngine struct {
	LevelData     `json:",inline"`
	Tags          []string                  `json:"tags"`
	Skin          LevelItemEngineSkin       `json:"skin"`
	Background    LevelItemEngineBackground `json:"background"`
	Effect        LevelItemEngineEffect     `json:"effect"`
	Particle      LevelItemEngineParticle   `json:"particle"`
	Thumbnail     LevelItemResource         `json:"thumbnail"`
	PlayData      LevelItemResource         `json:"playData"`
	WatchData     LevelItemResource         `json:"watchData"`
	PreviewData   LevelItemResource         `json:"previewData"`
	TutorialData  LevelItemResource         `json:"tutorialData"`
	Configuration LevelItemResource         `json:"configuration"`
}

type LevelItemUse struct {
	UseDefault bool `json:"useDefault"`
}

type LevelItem struct {
	LevelData     `json:",inline"`
	Engine        LevelItemEngine   `json:"engine"`
	UseSkin       LevelItemUse      `json:"useSkin"`
	UseBackground LevelItemUse      `json:"useBackground"`
	UseEffect     LevelItemUse      `json:"useEffect"`
	UseParticle   LevelItemUse      `json:"useParticle"`
	Cover         LevelItemResource `json:"cover"`
	BGM           LevelItemResource `json:"bgm"`
	Data          LevelItemResource `json:"data"`
}

type Level struct {
	Item         LevelItem `json:"item"`
	Description  string    `json:"description"`
	HasCommunity bool      `json:"hasCommunity"`
	LeaderBoards []any     `json:"leaderBoards"`
	Actions      []any     `json:"actions"`
	Sections     []any     `json:"sections"`
}
