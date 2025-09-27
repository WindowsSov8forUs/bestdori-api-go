package endpoints

// API endpoints
const (
	PostBasic   = "/api/post/basic"
	PostDetails = "/api/post/details"
	PostList    = "/api/post/list"
	PostTag     = "/api/post/tag"
	PostPost    = "/api/post"
	PostFind    = "/api/post/find"
	PostLike    = "/api/post/like"

	ChartsInfo = "/api/charts/%d/%s.json"

	CharactersInfo = "/api/characters/%d.json"
	CharactersAll  = "/api/characters/all.%d.json"
	CharactersMain = "/api/characters/main.%d.json"

	CardsInfo = "/api/cards/%d.json"
	CardsAll  = "/api/cards/all.%d.json"

	CostumesInfo = "/api/costumes/%d.json"
	CostumesAll  = "/api/costumes/all.%d.json"

	EventsInfo = "/api/events/%d.json"
	EventsAll  = "/api/events/all.%d.json"

	FestivalStages         = "/api/festival/stages/%d.json"
	FestivalRotationMusics = "/api/festival/rotationMusics/%d.json"

	SongsInfo = "/api/songs/%d.json"
	SongsAll  = "/api/songs/all.%d.json"

	BandsAll  = "/api/bands/all.%d.json"
	BandsMain = "/api/bands/main.%d.json"

	MiscLLSif = "/api/misc/llsif.%d.json"

	TrackerEventTop     = "/api/eventtop/data"
	TrackerEventTracker = "/api/tracker/data"
	TrackerRates        = "/api/tracker/rates.json"

	StampsAll   = "/api/stamps/all.%d.json"
	ArchivesAll = "/api/archives/all.%d.json"
	ComicsAll   = "/api/comics/all.%d.json"
)
