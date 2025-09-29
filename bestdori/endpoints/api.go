package endpoints

import "strconv"

// Private constants for API endpoints
const (
	apiPostBasic   = "/api/post/basic"
	apiPostDetails = "/api/post/details"
	apiPostList    = "/api/post/list"
	apiPostTag     = "/api/post/tag"
	apiPost        = "/api/post"
	apiPostFind    = "/api/post/find"
	apiPostLike    = "/api/post/like"

	apiCharts = "/api/charts/"

	apiCharacters     = "/api/characters/"
	apiCharactersAll  = "/api/characters/all."
	apiCharactersMain = "/api/characters/main."

	apiCards    = "/api/cards/"
	apiCardsAll = "/api/cards/all."

	apiCostumes    = "/api/costumes/"
	apiCostumesAll = "/api/costumes/all."

	apiEvents    = "/api/events/"
	apiEventsAll = "/api/events/all."

	apiFestivalStages         = "/api/festival/stages/"
	apiFestivalRotationMusics = "/api/festival/rotationMusics/"

	apiGacha    = "/api/gacha/"
	apiGachaAll = "/api/gacha/all."

	apiSongs    = "/api/songs/"
	apiSongsAll = "/api/songs/all."

	apiLoginCampaignsInfo = "/api/loginCampaigns/"
	apiLoginCampaignsAll  = "/api/loginCampaigns/all."

	apiMissionsInfo = "/api/missions/"
	apiMissionsAll  = "/api/missions/all."

	apiBandsAll  = "/api/bands/all."
	apiBandsMain = "/api/bands/main."

	apiMiscLLSif = "/api/misc/llsif."

	apiPlayerInfo = "/api/player/"

	apiTrackerEventTop     = "/api/eventtop/data"
	apiTrackerEventTracker = "/api/tracker/data"
	apiTrackerRates        = "/api/tracker/rates.json"

	apiStampsAll                 = "/api/stamps/all."
	apiArchivesAll               = "/api/archives/all."
	apiMiracleTicketExchangesAll = "/api/miracleTicketExchanges/all."
	apiComicsAll                 = "/api/comics/all."
)

// API endpoints functions
func PostBasic() string {
	return apiPostBasic
}

func PostDetails() string {
	return apiPostDetails
}

func PostList() string {
	return apiPostList
}

func PostTag() string {
	return apiPostTag
}

func PostPost() string {
	return apiPost
}

func PostFind() string {
	return apiPostFind
}

func PostLike() string {
	return apiPostLike
}

func ChartsInfo(id int, diff string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiCharts)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString("/")
	builder.WriteString(diff)
	builder.WriteString(json)
	return builder.String()
}

func CharactersInfo(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiCharacters)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(json)
	return builder.String()
}

func CharactersAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiCharactersAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func CharactersMain(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiCharactersMain)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func CardsInfo(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiCards)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(json)
	return builder.String()
}

func CardsAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiCardsAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func CostumesInfo(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiCostumes)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(json)
	return builder.String()
}

func CostumesAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiCostumesAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func EventsInfo(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiEvents)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(json)
	return builder.String()
}

func EventsAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiEventsAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func FestivalStages(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiFestivalStages)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(json)
	return builder.String()
}

func FestivalRotationMusics(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiFestivalRotationMusics)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(json)
	return builder.String()
}

func GachaInfo(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiGacha)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(json)
	return builder.String()
}

func GachaAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiGachaAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func SongsInfo(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiSongs)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(json)
	return builder.String()
}

func SongsAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiSongsAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func LoginCampaignsInfo(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiLoginCampaignsInfo)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(json)
	return builder.String()
}

func LoginCampaignsAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiLoginCampaignsAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func MissionsInfo(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiMissionsInfo)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(json)
	return builder.String()
}

func MissionsAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiMissionsAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func BandsAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiBandsAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func BandsMain(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiBandsMain)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func MiscLLSif(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiMiscLLSif)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func PlayerInfo(server string, id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiPlayerInfo)
	builder.WriteString(server)
	builder.WriteString("/")
	builder.WriteString(strconv.Itoa(id))
	return builder.String()
}

func TrackerEventTop() string {
	return apiTrackerEventTop
}

func TrackerEventTracker() string {
	return apiTrackerEventTracker
}

func TrackerRates() string {
	return apiTrackerRates
}

func StampsAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiStampsAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func ArchivesAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiArchivesAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func MiracleTicketExchangesAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiMiracleTicketExchangesAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}

func ComicsAll(index int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiComicsAll)
	builder.WriteString(strconv.Itoa(index))
	builder.WriteString(json)
	return builder.String()
}
