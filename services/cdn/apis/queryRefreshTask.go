package apis

import "github.com/baoyxing/ncloud_go/core"

type QueryRefreshTaskByIdRequest struct {
	core.NCloudRequest
	RefreshTaskById
}

type RefreshTaskById struct {
	TaskId string `json:"taskId"`
	Date   string `json:"date"`
}

func NewQueryRefreshTaskByIdRequest() *QueryRefreshTaskByIdRequest {
	return &QueryRefreshTaskByIdRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "refreshTask",
			Method: core.MethodPost,
		},
	}
}
func NewQueryRefreshTaskByIdRequestWithAllParams(
	taskId string,
	date string) *QueryRefreshTaskByIdRequest {
	return &QueryRefreshTaskByIdRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "refreshTask",
			Method: core.MethodPost,
		},
		RefreshTaskById: RefreshTaskById{
			taskId,
			date,
		},
	}
}

func (q QueryRefreshTaskByIdRequest) SetTaskId(taskId string) {
	q.TaskId = taskId
}

func (q QueryRefreshTaskByIdRequest) SetDate(date string) {
	q.Date = date
}

type QueryRefreshTaskByIdResponse struct {
	Ok   bool                    `json:"ok"`
	Msg  string                  `json:"msg"`
	Data []RefreshTaskByIdResult `json:"data"`
}

type RefreshTaskByIdResult struct {
	Url        string
	Status     core.RefreshStatus
	CreateTime int64
}
