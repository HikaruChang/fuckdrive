package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (Store *Store) GetUserInfo() (getUserInfoResp *GetUserInfoResp) {
	HTTPClient := &http.Client{
		Transport: &HTTPTransport,
	}
	Request, err := http.NewRequest("POST", URL.isLogin, nil)
	if err != nil {
		logrus.Error(err)
	}
	Request.Header.Set("content-type", "application/json")
	Request.Header.Set("Cookie", JSONConfig().Cookie)

	Response, err := HTTPClient.Do(Request)
	if err != nil {
		logrus.Error(err)
	}
	defer Response.Body.Close()
	ResponseBody, err := ioutil.ReadAll(Response.Body)
	if err != nil {
		logrus.Error(err)
	}
	getUserInfoResp = new(GetUserInfoResp)
	json.Unmarshal(ResponseBody, &getUserInfoResp)
	return getUserInfoResp
}
