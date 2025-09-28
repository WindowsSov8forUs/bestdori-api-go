package endpoints

import "strconv"

// RES endpoints
const (
	resIcon  = "/res/icon/"
	resImage = "/res/image/"

	charaIcon = "chara_icon_"
)

func ResIconSvg(name string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(resIcon)
	builder.WriteString(name)
	builder.WriteString(svg)
	return builder.String()
}

func ResIconPng(name string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(resIcon)
	builder.WriteString(name)
	builder.WriteString(png)
	return builder.String()
}

func ResImagePng(name string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(resImage)
	builder.WriteString(name)
	builder.WriteString(png)
	return builder.String()
}

func CharaIcon(id int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(charaIcon)
	builder.WriteString(strconv.Itoa(id))
	return builder.String()
}
