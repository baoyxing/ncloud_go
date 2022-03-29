package client

import (
	"fmt"
	"github.com/baoyxing/ncloud_go/core"
	"github.com/baoyxing/ncloud_go/services/cdn/apis"
	"testing"
)

const (
	accessKey = "#######"
	secretKey = "########"
)

func TestCdnClient_CreateRefreshTask(t *testing.T) {

	credentials := core.NewCredentials(accessKey, secretKey)
	cdnClient := NewCdnClient(credentials)
	cdnClient.DisableLogger()
	domain := "download-v1.xyuncloud.com"
	dataRequest := apis.NewCreateRefreshTaskRequestByRefresh([]string{"/site/15847292.mp4.f30.mp4"}, &domain)
	dataResponse, err := cdnClient.CreateRefreshTask(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("msg:", dataResponse.Msg)
		fmt.Println("taskID:", dataResponse.Data.TaskId)

	}
}

func TestCdnClient_QueryRefreshTaskById(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	cdnClient := NewCdnClient(credentials)
	cdnClient.DisableLogger()
	dataRequest := apis.
		NewQueryRefreshTaskByIdRequestWithAllParams("68792190-3e35-4d88-852c-661e28648d63", "20220328")

	dataResponse, err := cdnClient.QueryRefreshTaskById(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("msg:", dataResponse.Msg)
		for _, value := range dataResponse.Data {
			fmt.Println("url:", value.Url)
			fmt.Println("createAt", value.CreateTime)
		}
	}
}
