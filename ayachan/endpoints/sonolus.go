package endpoints

import "strconv"

const (
	sonolusLevels    = "/levels"
	sonolusLevelsGet = "/bdv2.json"
)

func LevelsPost() string {
	return sonolusLevels
}

func LevelsInfo(uid int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(sonolusLevelsGet)
	builder.WriteString("/")
	builder.WriteString(strconv.Itoa(uid))
	return builder.String()
}

func LevelsGet(uid int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(sonolusLevels)
	builder.WriteString("/")
	builder.WriteString(strconv.Itoa(uid))
	builder.WriteString(sonolusLevelsGet)
	return builder.String()
}
