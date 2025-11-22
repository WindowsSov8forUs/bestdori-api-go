package chartmetrics

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/ayachan/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/ayachan/endpoints"
	bdDto "github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// ChartMetricsBandori BanG Dream! 谱面分析
func ChartMetricsBandori(api *uniapi.UniAPI, chartId int, diffStr bdDto.ChartDifficultyName) (*dto.ChartMetrics, error) {
	endpoint := endpoints.ChartMetricsBandori(chartId, string(diffStr))
	return uniapi.Get[dto.ChartMetrics](api, endpoint, nil)
}

// ChartMetricsBestdori Bestdori 社区谱面分析
func ChartMetricsBestdori(api *uniapi.UniAPI, chartId int) (*dto.ChartMetrics, error) {
	endpoint := endpoints.ChartMetricsBestdori(chartId)
	return uniapi.Get[dto.ChartMetrics](api, endpoint, nil)
}

// ChartMetricsCustom 自定义谱面分析
func ChartMetricsCustom(api *uniapi.UniAPI, diffStr string, chart *bdDto.Chart) (*dto.ChartMetrics, error) {
	endpoint := endpoints.ChartMetricsCustom(diffStr)
	return uniapi.Post[dto.ChartMetrics](api, endpoint, chart, nil)
}
