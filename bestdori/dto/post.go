package dto

type PostSongType string

const (
	PostSongTypeBandori PostSongType = "bandori"
	PostSongTypeLLSIF   PostSongType = "llsif"
	PostSongTypeCustom  PostSongType = "custom"
)

type PostSong struct {
	Type  PostSongType `json:"type"`
	Id    int          `json:"id,omitempty"`
	Audio string       `json:"audio,omitempty"`
	Cover string       `json:"cover,omitempty"`
}

type PostContentType string

const (
	PostContentTypeText    PostContentType = "text"
	PostContentTypeBr      PostContentType = "br"
	PostContentTypeEmoji   PostContentType = "emoji"
	PostContentTypeMention PostContentType = "mention"
	PostContentTypeHeading PostContentType = "heading"
	PostContentTypeImage   PostContentType = "image"
	PostContentTypeLink    PostContentType = "link"
	PostContentTypeList    PostContentType = "list"
)

type PostContentTarget string

const (
	PostContentTargetURL                 PostContentTarget = "url"
	PostContentTargetCharacterSingle     PostContentTarget = "character-single"
	PostContentTargetCardSingle          PostContentTarget = "card-single"
	PostContentTargetCostumeSingle       PostContentTarget = "costume-single"
	PostContentTargetEventSingle         PostContentTarget = "event-single"
	PostContentTargetGachaSingle         PostContentTarget = "gacha-single"
	PostContentTargetSongSingle          PostContentTarget = "song-single"
	PostContentTargetLoginCampaignSingle PostContentTarget = "logincampaign-single"
	PostContentTargetComicSingle         PostContentTarget = "comic-single"
	PostContentTargetMissionSingle       PostContentTarget = "mission-single"
	PostContentTargetCharacterInfo       PostContentTarget = "character-info"
	PostContentTargetCardInfo            PostContentTarget = "card-info"
	PostContentTargetCardIcon            PostContentTarget = "card-icon"
	PostContentTargetCostumeInfo         PostContentTarget = "costume-info"
	PostContentTargetEventInfo           PostContentTarget = "event-info"
	PostContentTargetGachaInfo           PostContentTarget = "gacha-info"
	PostContentTargetSongInfo            PostContentTarget = "song-info"
	PostContentTargetLoginCampaignInfo   PostContentTarget = "logincampaign-info"
	PostContentTargetComicInfo           PostContentTarget = "comic-info"
	PostContentTargetMissionInfo         PostContentTarget = "mission-info"
)

type PostContent struct {
	Type    PostContentType   `json:"type"`
	Data    string            `json:"data,omitempty"`
	Margin  string            `json:"margin,omitempty"`
	Display int               `json:"display,omitempty"`
	Object  []string          `json:"object,omitempty"`
	Target  PostContentTarget `json:"target,omitempty"`
}

type PostAuthorTitle struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Server Server `json:"server"`
}

type PostAuthor struct {
	Username string            `json:"username"`
	Nickname string            `json:"nickname"`
	Titles   []PostAuthorTitle `json:"titles"`
}

type PostTag struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type PostInfo struct {
	CategoryName string            `json:"categoryName"`
	CategoryId   string            `json:"categoryId"`
	Title        *string           `json:"title,omitempty"`
	Song         *PostSong         `json:"song,omitempty"`
	Artists      *string           `json:"artists,omitempty"`
	Diff         *ChartDifficulty  `json:"diff,omitempty"`
	Level        *int              `json:"level,omitempty"`
	Chart        *[]map[string]any `json:"chart,omitempty"`
	Content      []PostContent     `json:"content,omitempty"`
	Time         float64           `json:"time"`
	Author       PostAuthor        `json:"author"`
	Likes        int               `json:"likes"`
	Liked        bool              `json:"liked"`
	Tags         []PostTag         `json:"tags"`
}

type PostDetail struct {
	Result bool     `json:"result"`
	Post   PostInfo `json:"post"`
}

type PostBasicAuthor struct {
	Username string `json:"username"`
}

type PostBasic struct {
	Result bool            `json:"result"`
	Title  string          `json:"title"`
	Author PostBasicAuthor `json:"author"`
}

type PostListPost struct {
	Id           int              `json:"id"`
	CategoryName string           `json:"categoryName"`
	CategoryId   string           `json:"categoryId"`
	Title        *string          `json:"title,omitempty"`
	Song         *PostSong        `json:"song,omitempty"`
	Artists      *string          `json:"artists,omitempty"`
	Diff         *ChartDifficulty `json:"diff,omitempty"`
	Level        *int             `json:"level,omitempty"`
	Time         float64          `json:"time"`
	Content      []PostContent    `json:"content"`
	Author       PostAuthor       `json:"author"`
	Likes        int              `json:"likes"`
	Liked        bool             `json:"liked"`
	Tags         []PostTag        `json:"tags"`
}

type PostList struct {
	Result bool           `json:"result"`
	Posts  []PostListPost `json:"posts"`
	Count  int            `json:"count"`
}

type TagResultTag struct {
	Type  string `json:"type"`
	Data  string `json:"data"`
	Count int    `json:"count"`
}

type TagResult struct {
	Result bool           `json:"result"`
	Tags   []TagResultTag `json:"tags"`
}
