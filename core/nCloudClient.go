package core

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type NCloudClient struct {
	Credential  Credential
	Config      Config
	ServiceName string
	Revision    string
	Logger      Logger
}

func (c NCloudClient) Send(method, path, data string) ([]byte, error) {
	signer := NewSigner(c.Credential)
	requrl := "/" + c.Revision + "/" + c.ServiceName + "/" + path
	sign, timestamp := signer.Sign(requrl)
	url := c.Config.Scheme + "://" + c.Config.Endpoint + requrl
	return c.doSend(method, url, data, sign, timestamp, c.Config.Timeout)

}

func (c NCloudClient) doSend(method, url, data, signed, t string, timeout time.Duration) ([]byte, error) {
	fmt.Println("data:", data)
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		c.Logger.Log(LogFatal, err.Error())
		return nil, err
	}
	c.setHeader(req, c.Credential.AccessKey, signed, t)
	resp, err := client.Do(req)
	if err != nil {
		c.Logger.Log(LogError, err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c NCloudClient) setHeader(req *http.Request, ak string, signed string, t string) {
	req.Header.Set("Content-Type", "application/json")
	stdEncoding := "ak=" + ak + "&sign=" + signed + "&t=" + t
	req.Header.Set("Authorization", base64.StdEncoding.EncodeToString([]byte(stdEncoding)))
	for k, v := range req.Header {
		c.Logger.Log(LogInfo, k, v)
	}
}
