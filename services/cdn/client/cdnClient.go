package client

import (
	"encoding/json"
	"errors"
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
	config.SetEndpoint("ad.y2.xyuncloud.com")
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
