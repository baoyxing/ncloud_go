package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/baoyxing/ncloud_go/core"
	"github.com/baoyxing/ncloud_go/services/cdn/apis"
)

type CdnClient struct {
	core.NCloudClient
}

func NewCdnClient(credential *core.Credential) *CdnClient {
	if credential == nil {
		return nil
	}
	config := core.NewConfig()
	config.SetScheme(core.SchemeHttp)
	config.SetEndpoint("127.0.0.1:8888")
	return &CdnClient{
		core.NCloudClient{
			Credential:  *credential,
			Config:      *config,
			ServiceName: "cdn",
			Revision:    core.Version,
			Logger:      core.NewDefaultLogger(core.LogInfo),
		}}
}

func (c *CdnClient) DisableLogger() {
	c.Logger = core.NewDummyLogger()
}

/* 创建刷新预热任务， */
func (c *CdnClient) CreateRefreshTask(
	request *apis.CreateRefreshTaskRequest) (*apis.CreateRefreshTaskResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := json.Marshal(request.RefreshTaskRequest)
	if err != nil {
		return nil, err
	}
	resp, err := c.Send(request.Method, request.Path, string(jsonDataReq))

	if err != nil {
		return nil, err
	}
	jsonResp := &apis.CreateRefreshTaskResponse{}
	err = json.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

/* 根据taskId查询刷新预热任务 */
func (c *CdnClient) QueryRefreshTaskById(
	request *apis.QueryRefreshTaskByIdRequest) (*apis.QueryRefreshTaskByIdResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := json.Marshal(request.RefreshTaskById)
	if err != nil {
		return nil, err
	}
	resp, err := c.Send(request.Method, request.Path, string(jsonDataReq))

	if err != nil {
		return nil, err
	}
	jsonResp := &apis.QueryRefreshTaskByIdResponse{}
	err = json.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

/* 查询日志 */
func (c *CdnClient) QueryDomainLog(request *apis.QueryDomainLogRequest) (*apis.QueryDomainLogkResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := json.Marshal(request.QueryDomainLog)
	if err != nil {
		return nil, err
	}
	resp, err := c.Send(request.Method, request.Path, string(jsonDataReq))
	fmt.Println("resp:", string(resp))
	if err != nil {
		return nil, err
	}
	jsonResp := &apis.QueryDomainLogkResponse{}
	err = json.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

/* 查询域名 */
func (c *CdnClient) QueryDomain(request *apis.QueryDomainRequest) (*apis.QueryDomainResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request.Method, request.Path, "")
	if err != nil {
		return nil, err
	}
	jsonResp := &apis.QueryDomainResponse{}
	err = json.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

/* 查询运营商 */
func (c *CdnClient) QueryIsp(request *apis.QueryIspRequest) (*apis.QueryIspResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request.Method, request.Path, "")
	if err != nil {
		return nil, err
	}
	jsonResp := &apis.QueryIspResponse{}
	err = json.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

/* 查询地区*/
func (c *CdnClient) QueryRegion(request *apis.QueryRegionRequest) (*apis.QueryRegionResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request.Method, request.Path, "")
	if err != nil {
		return nil, err
	}
	jsonResp := &apis.QueryRegionResponse{}
	err = json.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

/* 查询域名带宽统计*/
func (c *CdnClient) QueryStatisticsDataByDomain(request *apis.QueryStatisticsDataByDomainRequest) (*apis.QueryStatisticsDataResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := json.Marshal(request.QueryStatisticsDataByDomain)
	fmt.Println("QueryStatisticsDataByDomain:", request.QueryStatisticsDataByDomain)
	if err != nil {
		return nil, err
	}
	resp, err := c.Send(request.Method, request.Path, string(jsonDataReq))
	if err != nil {
		return nil, err
	}
	jsonResp := &apis.QueryStatisticsDataResponse{}
	err = json.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err

}

/* 查询运营商带宽统计*/
func (c *CdnClient) QueryStatisticsDataByIsp(request *apis.QueryStatisticsDataByIspRequest) (*apis.QueryStatisticsDataResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := json.Marshal(request.StatisticsDataByIsp)
	if err != nil {
		return nil, err
	}
	resp, err := c.Send(request.Method, request.Path, string(jsonDataReq))
	if err != nil {
		return nil, err
	}
	jsonResp := &apis.QueryStatisticsDataResponse{}
	err = json.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}

/* 查询地区带宽统计*/
func (c *CdnClient) QueryStatisticsDataByRegion(request *apis.QueryStatisticsDataByRegionRequest) (*apis.QueryStatisticsDataResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	jsonDataReq, err := json.Marshal(request.StatisticsDataByRegion)
	if err != nil {
		return nil, err
	}
	resp, err := c.Send(request.Method, request.Path, string(jsonDataReq))
	if err != nil {
		return nil, err
	}
	jsonResp := &apis.QueryStatisticsDataResponse{}
	err = json.Unmarshal(resp, jsonResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}
	return jsonResp, err
}
