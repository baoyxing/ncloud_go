package apis

import (
	"fmt"
	"github.com/baoyxing/ncloud_go/core"
	"github.com/baoyxing/ncloud_go/models"
)

type QueryDomainLogRequest struct {
	core.NCloudRequest
	QueryDomainLog
}

type QueryDomainLog struct {
	/* 用户域名  */
	Domain string `json:"domain"`
	/* 查询起始时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
	StartTime string `json:"startTime"`

	/* 查询截止时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
	EndTime string `json:"endTime"`
	/* 时间间隔，取值(hour，day，fiveMin)，不传默认小时。 目前只支持hour (Optional) */
	Interval string `json:"interval"`

	/* 页面大小，默认值10 (Optional) */
	PageSize int `json:"pageSize"`

	/* 分页页数，默认值1 (Optional) */
	PageNumber int `json:"pageNumber"`
}

/*
 * param domain: 用户域名 (Required)
 * param startTime: 查询起始时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional)
 * param endTime: 查询截止时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional)
 * param interval: 时间间隔，取值(hour，day，fiveMin)，不传默认小时。 (Optional)
 * param logType: 日志类型，取值(log，zip,gz)，不传默认gz。 (Optional)
 * param pageSize: 页面大小，默认值10 (Optional)
 * param pageNumber: 分页页数，默认值1 (Optional)
 */
func NewQueryDomainLogRequestWithAllParams(
	domain string,
	startTime string,
	endTime string,
	interval string,
	pageSize int,
	pageNumber int,
) *QueryDomainLogRequest {
	fmt.Println("startTime:", startTime)
	fmt.Println("endTime:", endTime)

	return &QueryDomainLogRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "downLog",
			Method: core.MethodPost,
		},
		QueryDomainLog: QueryDomainLog{
			domain,
			startTime,
			endTime,
			interval,
			pageSize,
			pageNumber,
		},
	}
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryDomainLogRequestWithoutParam() *QueryDomainLogRequest {

	return &QueryDomainLogRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "/downLog",
			Method: core.MethodPost,
		},
	}
}

/* param domain: 用户域名(Required) */
func (r *QueryDomainLogRequest) SetDomain(domain string) {
	r.Domain = domain
}

/* param startTime: 查询起始时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional)*/
func (r *QueryDomainLogRequest) SetStartTime(startTime string) {
	r.StartTime = startTime
}

/* param endTime: 查询截止时间,北京时间，格式为:yyyy-MM-dd HH:mm:ss，示例:2018-10-21 10:00:00 (Optional) */
func (r *QueryDomainLogRequest) SetEndTime(endTime string) {
	r.EndTime = endTime
}

/* param interval: 时间间隔，取值(hour，day，fiveMin)，不传默认小时。(Optional) */
func (r *QueryDomainLogRequest) SetInterval(interval string) {
	r.Interval = interval
}

/* param pageSize: 页面大小，默认值10(Optional) */
func (r *QueryDomainLogRequest) SetPageSize(pageSize int) {
	r.PageSize = pageSize
}

/* param pageNumber: 分页页数，默认值1(Optional) */
func (r *QueryDomainLogRequest) SetPageNumber(pageNumber int) {
	r.PageNumber = pageNumber
}

type QueryDomainLogkResponse struct {
	Ok   bool                 `json:"ok"`
	Msg  string               `json:"msg"`
	Data QueryDomainLogResult `json:"data"`
}
type QueryDomainLogResult struct {
	Total      int                `json:"total"`
	PageSize   int                `json:"pageSize"`
	PageNumber int                `json:"pageNumber"`
	Urls       []models.DomainLog `json:"urls"`
}
