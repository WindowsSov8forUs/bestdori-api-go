package endpoints

import "strconv"

const (
	apiVersion = "/version"

	apiChartMetricsBandori  = "/chart/metrics/bandori/"
	apiChartMetricsBestdori = "/chart/metrics/bestdori/"
	apiChartMetricsCustom   = "/chart/metrics/custom/"
)

func VersionGet() string {
	return apiVersion
}

func ChartMetricsBandori(chartId int, diffStr string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiChartMetricsBandori)
	builder.WriteString(strconv.Itoa(chartId))
	builder.WriteByte('/')
	builder.WriteString(diffStr)
	return builder.String()
}

func ChartMetricsBestdori(chartId int) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiChartMetricsBestdori)
	builder.WriteString(strconv.Itoa(chartId))
	return builder.String()
}

func ChartMetricsCustom(diffStr string) string {
	builder := getBuilder()
	defer putBuilder(builder)

	builder.WriteString(apiChartMetricsCustom)
	builder.WriteString(diffStr)
	return builder.String()
}
