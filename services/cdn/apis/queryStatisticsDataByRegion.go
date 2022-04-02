package apis

import "github.com/baoyxing/ncloud_go/core"

type QueryStatisticsDataByRegionRequest struct {
	core.NCloudRequest
	StatisticsDataByRegion
}

type StatisticsDataByRegion struct {
	/* 查询起始时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
	StartTime *string `json:"startTime"`
	/* 查询截止时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss ，示例:2018-10-21 10:00:00 (Optional) */
	EndTime *string   `json:"endTime"`
	Domains *[]string `json:"domains"`
	Regions *[]string `json:"regions"`
	Isp     *[]string `json:"isp"`
}

func NewQueryStatisticsDataByRegionRequest() *QueryStatisticsDataByRegionRequest {
	return &QueryStatisticsDataByRegionRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "statisticsData/region",
			Method: core.MethodPost,
		},
	}
}

/*
*查询用户pin下有权限的所有域名统计带宽
* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 */
func NewQueryStatisticsDataByRegionAllDomainRequest(
	startTime *string,
	endTime *string,
	isp *[]string,
	regions *[]string) *QueryStatisticsDataByRegionRequest {
	return &QueryStatisticsDataByRegionRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "statisticsData/region",
			Method: core.MethodPost,
		},
		StatisticsDataByRegion: StatisticsDataByRegion{
			StartTime: startTime,
			EndTime:   endTime,
			Isp:       isp,
			Regions:   regions,
		},
	}
}

/*
*查询用户pin下有权限的所有域名下所有运营商下所有地区统计带宽
* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 */
func NewQueryStatisticsDataByRegionAllDomainAndIspRequest(
	startTime *string,
	endTime *string) *QueryStatisticsDataByRegionRequest {
	return &QueryStatisticsDataByRegionRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "statisticsData/region",
			Method: core.MethodPost,
		},
		StatisticsDataByRegion: StatisticsDataByRegion{
			StartTime: startTime,
			EndTime:   endTime,
		},
	}
}

/*
*查询用户pin下有权限的域名下所有运营商下所有运营商统计带宽
* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 */
func NewQueryStatisticsDataByAllRegionRequest(
	startTime *string,
	endTime *string,
	domains *[]string,
	isps *[]string,
) *QueryStatisticsDataByRegionRequest {
	return &QueryStatisticsDataByRegionRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "statisticsData/region",
			Method: core.MethodPost,
		},
		StatisticsDataByRegion: StatisticsDataByRegion{
			StartTime: startTime,
			EndTime:   endTime,
			Domains:   domains,
			Isp:       isps,
		},
	}
}

/*
* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
* param Domains: 需要查询的域名, 必须为用户pin下有权限的域名 (Optional)
 */
func NewQueryStatisticsDataByRegionRequestWithAllParams(
	startTime *string,
	endTime *string,
	domains *[]string,
	regions *[]string,
	isp *[]string) *QueryStatisticsDataByRegionRequest {
	return &QueryStatisticsDataByRegionRequest{
		core.NCloudRequest{
			Path:   "statisticsData/domain",
			Method: core.MethodPost,
		},
		StatisticsDataByRegion{
			StartTime: startTime,
			EndTime:   endTime,
			Domains:   domains,
			Isp:       isp,
			Regions:   regions,
		},
	}
}

/* 查询起始时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
func (r *QueryStatisticsDataByRegionRequest) SetStartTime(startTime string) {
	r.StartTime = &startTime
}

/* 查询截止时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
func (r *QueryStatisticsDataByRegionRequest) SetEndTime(endTime string) {
	r.StartTime = &endTime
}

/*param domains: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryStatisticsDataByRegionRequest) SetDomains(domains []string) {
	r.Domains = &domains
}

/*param isp:运营商(Optional) */
func (r *QueryStatisticsDataByRegionRequest) SetIsps(isp []string) {
	r.Isp = &isp
}

/*param regions:地区(Optional) */
func (r *QueryStatisticsDataByRegionRequest) SetRegion(regions []string) {
	r.Regions = &regions
}
