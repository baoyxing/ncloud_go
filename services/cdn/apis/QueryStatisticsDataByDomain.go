package apis

import "github.com/baoyxing/ncloud_go/core"

type QueryStatisticsDataByDomainRequest struct {
	core.NCloudRequest
	QueryStatisticsDataByDomain
}
type QueryStatisticsDataByDomain struct {
	/* 查询起始时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
	StartTime *string `json:"startTime"`

	/* 查询截止时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss ，示例:2018-10-21 10:00:00 (Optional) */
	EndTime *string `json:"endTime"`

	Domains *[]string `json:"domains"`
}

func NewQueryStatisticsDataByDomainRequest() *QueryStatisticsDataByDomainRequest {
	return &QueryStatisticsDataByDomainRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "statisticsData/domain",
			Method: core.MethodPost,
		},
	}
}

/*
* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param Domains: 需要查询的域名, 必须为用户pin下有权限的域名 (Optional)
 */
func NewQueryStatisticsDataByDomainRequestWithAllParams(
	startTime *string,
	endTime *string,
	domains *[]string) *QueryStatisticsDataByDomainRequest {
	return &QueryStatisticsDataByDomainRequest{
		core.NCloudRequest{
			Path:   "statisticsData/domain",
			Method: core.MethodPost,
		},
		QueryStatisticsDataByDomain{
			StartTime: startTime,
			EndTime:   endTime,
			Domains:   domains,
		},
	}
}

/*
*查询用户pin下有权限的所有域名统计信息
* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 */
func NewQueryStatisticsDataByAllDomainRequest(
	startTime *string,
	endTime *string) *QueryStatisticsDataByDomainRequest {
	return &QueryStatisticsDataByDomainRequest{
		core.NCloudRequest{
			Path:   "statisticsData/domain",
			Method: core.MethodPost,
		},
		QueryStatisticsDataByDomain{
			StartTime: startTime,
			EndTime:   endTime,
		},
	}
}

/* 查询起始时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
func (r *QueryStatisticsDataByDomainRequest) SetStartTime(startTime string) {
	r.StartTime = &startTime
}

/* 查询截止时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
func (r *QueryStatisticsDataByDomainRequest) SetEndTime(endTime string) {
	r.StartTime = &endTime
}

/*param domains: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryStatisticsDataByDomainRequest) SetDomains(domains []string) {
	r.Domains = &domains
}

type QueryStatisticsDataResponse struct {
	Ok   bool                        `json:"ok"`
	Msg  string                      `json:"msg"`
	Data []QueryStatisticsDataResult `json:"data"`
}

type QueryStatisticsDataResult struct {
	StartTime  string               `json:"startTime"`
	EndTime    string               `json:"endTime"`
	Statistics []StatisticsDataItem `json:"statistics"`
}

type StatisticsDataItem struct {
	/* UTC时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
	StartTime string `json:"startTime"`
	/* UTC时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
	EndTime   string `json:"endTime"`
	Domain    string `json:"domain"`
	Isp       string `json:"isp"`
	Region    string `json:"region"`
	Bandwidth int64  `json:"bandwidth"`
}
