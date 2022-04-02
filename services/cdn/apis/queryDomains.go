package apis

import "github.com/baoyxing/ncloud_go/core"

type QueryDomainRequest struct {
	core.NCloudRequest
}

func NewQueryDomainsRequest() *QueryDomainRequest {
	return &QueryDomainRequest{
		core.NCloudRequest{
			Path:   "domain",
			Method: core.MethodPost,
		},
	}
}

type QueryDomainResponse struct {
	Ok   bool     `json:"ok"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}
