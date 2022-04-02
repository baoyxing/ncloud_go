package apis

import "github.com/baoyxing/ncloud_go/core"

type QueryStatisticsDataByIspRequest struct {
	core.NCloudRequest
	StatisticsDataByIsp
}

type StatisticsDataByIsp struct {
	/* 查询起始时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
	StartTime *string `json:"startTime"`
	/* 查询截止时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss ，示例:2018-10-21 10:00:00 (Optional) */
	EndTime *string   `json:"endTime"`
	Domains *[]string `json:"domains"`
	Isp     *[]string `json:"isp"`
}

func NewQueryStatisticsDataByIspRequest() *QueryStatisticsDataByIspRequest {
	return &QueryStatisticsDataByIspRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "statisticsData/isp",
			Method: core.MethodPost,
		},
	}
}

/*
*查询用户pin下有权限的所有域名统计带宽
* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 */
func NewQueryStatisticsDataByIspAllDomainRequest(
	startTime *string,
	endTime *string,
	isp *[]string) *QueryStatisticsDataByIspRequest {
	return &QueryStatisticsDataByIspRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "statisticsData/isp",
			Method: core.MethodPost,
		},
		StatisticsDataByIsp: StatisticsDataByIsp{
			StartTime: startTime,
			EndTime:   endTime,
			Isp:       isp,
		},
	}
}

/*
*查询用户pin下有权限的所有域名下所有运营商统计带宽
* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 */
func NewQueryStatisticsDataByIspAllDomainAndIspRequest(
	startTime *string,
	endTime *string) *QueryStatisticsDataByIspRequest {
	return &QueryStatisticsDataByIspRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "statisticsData/isp",
			Method: core.MethodPost,
		},
		StatisticsDataByIsp: StatisticsDataByIsp{
			StartTime: startTime,
			EndTime:   endTime,
		},
	}
}

/*
*查询用户pin下有权限的域名下所有运营商统计带宽
* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 */
func NewQueryStatisticsDataByAlIspRequest(
	startTime *string,
	endTime *string,
	domains *[]string,
) *QueryStatisticsDataByIspRequest {
	return &QueryStatisticsDataByIspRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "statisticsData/isp",
			Method: core.MethodPost,
		},
		StatisticsDataByIsp: StatisticsDataByIsp{
			StartTime: startTime,
			EndTime:   endTime,
			Domains:   domains,
		},
	}
}

/*
* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param Domains: 需要查询的域名, 必须为用户pin下有权限的域名 (Optional)
 */
func NewQueryStatisticsDataByIspRequestWithAllParams(
	startTime *string,
	endTime *string,
	domains *[]string,
	isp *[]string) *QueryStatisticsDataByIspRequest {
	return &QueryStatisticsDataByIspRequest{
		core.NCloudRequest{
			Path:   "statisticsData/domain",
			Method: core.MethodPost,
		},
		StatisticsDataByIsp{
			StartTime: startTime,
			EndTime:   endTime,
			Domains:   domains,
			Isp:       isp,
		},
	}
}

/* 查询起始时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
func (r *QueryStatisticsDataByIspRequest) SetStartTime(startTime string) {
	r.StartTime = &startTime
}

/* 查询截止时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
func (r *QueryStatisticsDataByIspRequest) SetEndTime(endTime string) {
	r.StartTime = &endTime
}

/*param domains: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryStatisticsDataByIspRequest) SetDomains(domains []string) {
	r.Domains = &domains
}

/*param isp:运营商(Optional) */
func (r *QueryStatisticsDataByIspRequest) SetIsps(isp []string) {
	r.Isp = &isp
}
