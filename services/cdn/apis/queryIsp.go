package apis

import "github.com/baoyxing/ncloud_go/core"

type QueryIspRequest struct {
	core.NCloudRequest
}

func NewQueryIspRequest() *QueryIspRequest {
	return &QueryIspRequest{
		core.NCloudRequest{
			Path:   "isp",
			Method: core.MethodPost,
		},
	}
}

type QueryIspResponse struct {
	Ok   bool     `json:"ok"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}
