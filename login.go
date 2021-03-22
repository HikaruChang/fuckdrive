package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/sirupsen/logrus"
)

func (Store *Store) appType() (appTypeResp *AppTypeResp) {
	HTTPClient := &http.Client{
		Transport: &HTTPTransport,
	}
	Request, err := http.NewRequest("POST", URL.appType, nil)
	if err != nil {
		logrus.Error(err)
	}
	Request.Header.Set("content-type", "application/json")
	Response, err := HTTPClient.Do(Request)
	if err != nil {
		logrus.Error(err)
	}
	defer Response.Body.Close()
	ResponseBody, err := ioutil.ReadAll(Response.Body)
	if err != nil {
		logrus.Error(err)
	}
	appTypeResp = new(AppTypeResp)
	json.Unmarshal(ResponseBody, &appTypeResp)
	setCookie := Response.Header.Get("Set-Cookie")
	cookieRegexp := regexp.MustCompile(`^ASP.NET_SessionId=(.+?);`)
	JSONConfig := JSONConfig()
	JSONConfig.Cookie = cookieRegexp.FindString(setCookie)
	err = JSONEdit(JSONConfig)
	if err != nil {
		logrus.Error(err)
	}
	return appTypeResp
}

func (Store *Store) mobileLogin() (mobileLoginResp *MobileLoginResp) {
	if Store.appType().Result && Store.appType().IsSuccess {
		Params := MobileLoginParams{
			LoginType: 0,
			UserName:  JSONConfig().Username,
			Password:  JSONConfig().Password,
		}
		ParamsJSON, err := json.Marshal(Params)
		if err != nil {
			logrus.Error(err)
		}
		HTTPClient := &http.Client{
			Transport: &HTTPTransport,
		}
		Request, err := http.NewRequest("POST", URL.mobileLogin, bytes.NewBuffer(ParamsJSON))
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
		mobileLoginResp = new(MobileLoginResp)
		json.Unmarshal(ResponseBody, &mobileLoginResp)
		return mobileLoginResp
	}
	logrus.Error(string(Store.appType().Message))
	return
}

func (Store *Store) isLogin() (mobileLoginResp *MobileLoginResp) {
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
	mobileLoginResp = new(MobileLoginResp)
	json.Unmarshal(ResponseBody, &mobileLoginResp)
	return mobileLoginResp
}
