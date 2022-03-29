package client

import (
	"fmt"
	"github.com/baoyxing/ncloud_go/core"
	"github.com/baoyxing/ncloud_go/services/cdn/apis"
	"testing"
)

const (
	accessKey = "OF31SmcMfwOPcuVf"
	secretKey = ""
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
