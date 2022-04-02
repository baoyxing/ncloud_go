package client

import (
	"fmt"
	"github.com/baoyxing/ncloud_go/core"
	"github.com/baoyxing/ncloud_go/services/cdn/apis"
	"testing"
)

const (
	accessKey = "OF31SmcMfwOPcuVf"
	secretKey = "2b79dwyBxW8DRqarSipo9SqvaudM8Stn"
)

func TestCdnClient_CreateRefreshTask(t *testing.T) {

	credentials := core.NewCredentials(accessKey, secretKey)
	cdnClient := NewCdnClient(credentials)
	cdnClient.DisableLogger()
	domain := "download-v1.xyuncloud.com"
	dataRequest := apis.NewCreateRefreshTaskRequestByRefresh([]string{"/site/15847292.mp4.f30.mp4"}, domain)
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

func TestCdnClient_QueryDomainLog(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	cdnClient := NewCdnClient(credentials)
	//cdnClient.DisableLogger()
	endTime := "2022-04-01 10:35:50"
	startTime := "2022-03-03 10:35:50"
	dataRequest := apis.
		NewQueryDomainLogRequestWithAllParams("download-v1.xyuncloud.com", startTime, endTime, "", 2, 1)

	dataResponse, err := cdnClient.QueryDomainLog(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		for _, value := range dataResponse.Data.Urls {
			fmt.Println("url:", value.Url)
		}
	}
}
func TestCdnClient_QueryDomain(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	cdnClient := NewCdnClient(credentials)
	//cdnClient.DisableLogger()
	dataRequest := apis.NewQueryDomainsRequest()
	dataResponse, err := cdnClient.QueryDomain(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("data:", dataResponse.Data)
	}
}

func TestCdnClient_QueryRegion(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	cdnClient := NewCdnClient(credentials)
	//cdnClient.DisableLogger()
	dataRequest := apis.NewQueryRegionRequest()
	dataResponse, err := cdnClient.QueryRegion(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("data:", dataResponse.Data)
	}
}

func TestCdnClient_QueryIsp(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	cdnClient := NewCdnClient(credentials)
	//cdnClient.DisableLogger()
	dataRequest := apis.NewQueryIspRequest()
	dataResponse, err := cdnClient.QueryIsp(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("data:", dataResponse.Data)
	}
}

func TestCdnClient_QueryStatisticsDataByDomain(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	cdnClient := NewCdnClient(credentials)
	//cdnClient.DisableLogger()
	endTime := "2022-04-01 10:35:00"
	startTime := "2022-04-01 10:25:00"
	domain := []string{"download-v1.xyuncloud.com"}
	dataRequest := apis.NewQueryStatisticsDataByDomainRequestWithAllParams(&startTime, &endTime, &domain)
	dataResponse, err := cdnClient.QueryStatisticsDataByDomain(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("msg:", dataResponse.Msg)
		for _, values := range dataResponse.Data {
			fmt.Println("startTime:", values.StartTime)
			fmt.Println("endTime:", values.EndTime)

			for _, value := range values.Statistics {
				fmt.Println("domain:", value.Domain)
				fmt.Println("bandwidth:", value.Bandwidth)
			}
		}
	}
}

func TestCdnClient_QueryStatisticsDataByIsp(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	cdnClient := NewCdnClient(credentials)
	//cdnClient.DisableLogger()
	endTime := "2022-04-01 10:35:00"
	startTime := "2022-04-01 10:25:00"
	dataRequest := apis.NewQueryStatisticsDataByIspAllDomainAndIspRequest(&startTime, &endTime)
	dataResponse, err := cdnClient.QueryStatisticsDataByIsp(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("msg:", dataResponse.Msg)
		for _, values := range dataResponse.Data {
			fmt.Println("startTime:", values.StartTime)
			fmt.Println("endTime:", values.EndTime)

			for _, value := range values.Statistics {
				fmt.Println("domain:", value.Domain)
				fmt.Println("bandwidth:", value.Bandwidth)
				fmt.Println("isp:", value.Isp)
			}
		}
	}
}

func TestCdnClient_QueryStatisticsDataByRegion(t *testing.T) {
	credentials := core.NewCredentials(accessKey, secretKey)
	cdnClient := NewCdnClient(credentials)
	//cdnClient.DisableLogger()
	endTime := "2022-04-01 10:35:00"
	startTime := "2022-04-01 10:25:00"
	dataRequest := apis.NewQueryStatisticsDataByRegionAllDomainAndIspRequest(&startTime, &endTime)
	dataResponse, err := cdnClient.QueryStatisticsDataByRegion(dataRequest)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("msg:", dataResponse.Msg)
		for _, values := range dataResponse.Data {
			fmt.Println("startTime:", values.StartTime)
			fmt.Println("endTime:", values.EndTime)

			for _, value := range values.Statistics {
				fmt.Println("domain:", value.Domain)
				fmt.Println("bandwidth:", value.Bandwidth)
				fmt.Println("isp:", value.Isp)
				fmt.Println("region:", value.Region)
			}
		}
	}
}
