package apis

import "github.com/baoyxing/ncloud_go/core"

type QueryRegionRequest struct {
	core.NCloudRequest
}

func NewQueryRegionRequest() *QueryRegionRequest {
	return &QueryRegionRequest{
		core.NCloudRequest{
			Path:   "region",
			Method: core.MethodPost,
		},
	}
}

type QueryRegionResponse struct {
	Ok   bool     `json:"ok"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}
