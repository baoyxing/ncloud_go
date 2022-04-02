package apis

import "github.com/baoyxing/ncloud_go/core"

type CreateRefreshTaskRequest struct {
	core.NCloudRequest
	RefreshTaskRequest
}

type RefreshTaskRequest struct {
	/* 刷新预热类型,(url:url刷新,dir:目录刷新,prefetch:预热)*/
	TaskType string `json:"taskType"`
	/* 域名 目前只支持单域名*/
	Domain string `json:"domain"`
	/*  文件名称 */
	Files []string `json:"files"`
}

func NewCreateRefreshTaskRequest() *CreateRefreshTaskRequest {
	return &CreateRefreshTaskRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "refresh",
			Method: core.MethodPost,
		},
	}
}

func NewCreateRefreshTaskRequestWithAllParams(
	taskType string,
	files []string,
	domain string) *CreateRefreshTaskRequest {
	return &CreateRefreshTaskRequest{
		NCloudRequest: core.NCloudRequest{
			Path:   "refresh",
			Method: core.MethodPost,
		},
		RefreshTaskRequest: RefreshTaskRequest{
			TaskType: taskType,
			Files:    files,
			Domain:   domain,
		},
	}
}

func NewCreateRefreshTaskRequestByRefresh(
	files []string,
	domain string) *CreateRefreshTaskRequest {
	taskType := "url"
	return NewCreateRefreshTaskRequestWithAllParams(taskType, files, domain)
}

func NewCreateRefreshTaskRequestByPrefetch(
	files []string,
	domain string) *CreateRefreshTaskRequest {
	taskType := "prefetch"
	return NewCreateRefreshTaskRequestWithAllParams(taskType, files, domain)
}

func NewCreateRefreshTaskRequestByDir(
	files []string,
	domain string) *CreateRefreshTaskRequest {
	taskType := "dir"
	return NewCreateRefreshTaskRequestWithAllParams(taskType, files, domain)
}

/* param taskType: 刷新预热类型,(url:url刷新,dir:目录刷新,prefetch:预热)*/
func (c *CreateRefreshTaskRequest) SetTaskType(taskType string) {
	c.TaskType = taskType
}

/* param files: (Optional) */
func (c *CreateRefreshTaskRequest) SetFiles(files []string) {
	c.Files = files
}

/* param domain: (Optional) */
func (c *CreateRefreshTaskRequest) SetDomain(domain string) {
	c.Domain = domain
}

type CreateRefreshTaskResponse struct {
	Ok   bool                    `json:"ok"`
	Msg  string                  `json:"msg"`
	Data CreateRefreshTaskResult `json:"data"`
}

type CreateRefreshTaskResult struct {
	TaskId string `json:"taskId"` //任务id
	Date   string `json:"date"`   //日期[yyyyMMdd]
}
