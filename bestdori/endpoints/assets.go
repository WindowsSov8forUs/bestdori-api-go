package endpoints

import (
	"fmt"
	"strconv"
)

// Private constants for assets endpoints
const (
	assets = "/assets/"

	assetsCharactersKvImage1 = "/ui/character_kv_image/"
	assetsCharactersKvImage2 = "_rip/image.png"

	assetsCharactersResourceSet1 = "/characters/resourceset/"
	assetsCharactersResourceSet2 = "_rip/"

	assetsCharactersLiveSD1 = "/characters/livesd/"
	assetsCharactersLiveSD2 = "_rip/sdchara.png"

	assetsEventBanner1 = "/event/"
	assetsEventBanner2 = "/images_rip/banner.png"

	assetsEventLogo1 = "/event/"
	assetsEventLogo2 = "/images_rip/logo.png"

	assetsEventTopScreen1 = "/event/"
	assetsEventTopScreen2 = "/topscreen_rip/"
	assetsEventTopScreen3 = "_eventtop.png"

	assetsEventLoginBonus1 = "/event/loginbonus/"
	assetsEventLoginBonus2 = "_rip/background.png"

	assetsSongsMusicJacket1 = "/musicjacket/musicjacket"
	assetsSongsMusicJacket2 = "_rip/assets-star-forassetbundle-startapp-musicjacket-musicjacket"
	assetsSongsMusicJacket3 = "-jacket.png"

	assetsSongsSound1 = "/sound/bgm"
	assetsSongsSound2 = "_rip/bgm"

	assetsThumbChara1 = "/thumb/chara/card"
	assetsThumbChara2 = "_rip/"

	assetsThumbDegree = "/thumb/degree_rip/"

	assetsThumbCostume1 = "/thumb/costume/group"
	assetsThumbCostume2 = "_rip/"

	assetsStampGet = "/stamp/01_rip/"

	assetsHomebannerGet = "/homebanner_rip/"

	assetsGachaScreen1 = "/gacha/screen/gacha"
	assetsGachaScreen2 = "_rip/"

	assetsComicComic1 = "/comic/comic_"
	assetsComicComic2 = "_rip/"

	assetsComicThumbnail1 = "/comic/comic_"
	assetsComicThumbnail2 = "_thumbnail/"
	assetsComicThumbnail3 = "_rip/"

	assetsBandLogo1 = "/band/logo/"
	assetsBandLogo2 = "_rip/"
)

// ASSETS endpoints functions
func CharactersKvImage(server string, id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsCharactersKvImage1)
	fmt.Fprintf(builder, "%03d", id)
	builder.WriteString(assetsCharactersKvImage2)
	return builder.String()
}

func CharactersResourceSet(server, resourceSetName, name, typ string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsCharactersResourceSet1)
	builder.WriteString(resourceSetName)
	builder.WriteString(assetsCharactersResourceSet2)
	builder.WriteString(name)
	builder.WriteString("_")
	builder.WriteString(typ)
	builder.WriteString(png)
	return builder.String()
}

func CharactersLiveSD(server, sdResourceName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsCharactersLiveSD1)
	builder.WriteString(sdResourceName)
	builder.WriteString(assetsCharactersLiveSD2)
	return builder.String()
}

func EventBanner(server, assetBundleName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsEventBanner1)
	builder.WriteString(assetBundleName)
	builder.WriteString(assetsEventBanner2)
	return builder.String()
}

func EventLogo(server, assetBundleName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsEventLogo1)
	builder.WriteString(assetBundleName)
	builder.WriteString(assetsEventLogo2)
	return builder.String()
}

func EventTopScreen(server, assetBundleName, typ string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsEventTopScreen1)
	builder.WriteString(assetBundleName)
	builder.WriteString(assetsEventTopScreen2)
	builder.WriteString(typ)
	builder.WriteString(assetsEventTopScreen3)
	return builder.String()
}

func EventLoginBonus(server, assetBundleName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsEventLoginBonus1)
	builder.WriteString(assetBundleName)
	builder.WriteString(assetsEventLoginBonus2)
	return builder.String()
}

func SongsMusicJacket(server string, index int, jacketImage string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsSongsMusicJacket1)
	fmt.Fprintf(builder, "%02d", index)
	builder.WriteString(assetsSongsMusicJacket2)
	fmt.Fprintf(builder, "%02d", index)
	builder.WriteString("-")
	builder.WriteString(jacketImage)
	builder.WriteString(assetsSongsMusicJacket3)
	return builder.String()
}

func SongsSound(server string, id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsSongsSound1)
	fmt.Fprintf(builder, "%03d", id)
	builder.WriteString(assetsSongsSound2)
	fmt.Fprintf(builder, "%03d", id)
	builder.WriteString(mp3)
	return builder.String()
}

func ThumbChara(server string, id int, resourceSetName, typ string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsThumbChara1)
	fmt.Fprintf(builder, "%05d", id)
	builder.WriteString(assetsThumbChara2)
	builder.WriteString(resourceSetName)
	builder.WriteString("_")
	builder.WriteString(typ)
	builder.WriteString(png)
	return builder.String()
}

func ThumbDegree(server, degreeName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsThumbDegree)
	builder.WriteString(degreeName)
	builder.WriteString(png)
	return builder.String()
}

func ThumbCostume(server string, id int, assetBundleName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsThumbCostume1)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(assetsThumbCostume2)
	builder.WriteString(assetBundleName)
	builder.WriteString(png)
	return builder.String()
}

func StampGet(server, imageName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsStampGet)
	builder.WriteString(imageName)
	builder.WriteString(png)
	return builder.String()
}

func HomebannerGet(server, BannerAssetBundleName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsHomebannerGet)
	builder.WriteString(BannerAssetBundleName)
	builder.WriteString(png)
	return builder.String()
}

func GachaScreen(server string, id int, assetName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsGachaScreen1)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(assetsGachaScreen2)
	builder.WriteString(assetName)
	builder.WriteString(png)
	return builder.String()
}

func ComicComic(server, typ, assetBundleName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsComicComic1)
	builder.WriteString(typ)
	builder.WriteString("/")
	builder.WriteString(assetBundleName)
	builder.WriteString(assetsComicComic2)
	builder.WriteString(assetBundleName)
	builder.WriteString(png)
	return builder.String()
}

func ComicThumbnail(server, typ, assetBundleName string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsComicThumbnail1)
	builder.WriteString(typ)
	builder.WriteString(assetsComicThumbnail2)
	builder.WriteString(assetBundleName)
	builder.WriteString(assetsComicThumbnail3)
	builder.WriteString(assetBundleName)
	builder.WriteString(png)
	return builder.String()
}

func BandLogo(server string, id int, typ string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(assets)
	builder.WriteString(server)
	builder.WriteString(assetsBandLogo1)
	fmt.Fprintf(builder, "%03d", id)
	builder.WriteString(assetsBandLogo2)
	builder.WriteString(typ)
	builder.WriteString(png)
	return builder.String()
}
