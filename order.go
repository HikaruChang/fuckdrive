package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func (Store *Store) GetDaysOrder() (getDaysOrderResp *GetDaysOrderResp) {
	HTTPClient := &http.Client{
		Transport: &HTTPTransport,
	}
	Request, err := http.NewRequest("POST", URL.getDaysOrder, nil)
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
	dresp := new(Dresp)
	json.Unmarshal(ResponseBody, &dresp)
	newJson := RegReplace(RegDelete(dresp.D, `(\r|\n)`), `"`, `\"`)
	getDaysOrderResp = new(GetDaysOrderResp)
	json.Unmarshal([]byte(newJson), &getDaysOrderResp)
	return getDaysOrderResp
}

func (Store *Store) GetInfoByTimeId() (getInfoByTimeIdResp *GetInfoByTimeIdResp) {
	Params := GetInfoByTimeIdParams{
		ArgDay:     fmt.Sprint(time.Now().Year(), "-", int(time.Now().Month()), "-", time.Now().Day()+1),
		NPageIndex: 1,
		NUnitNO:    3,
	}
	ParamsJSON, err := json.Marshal(Params)
	if err != nil {
		logrus.Error(err)
	}
	HTTPClient := &http.Client{
		Transport: &HTTPTransport,
	}
	Request, err := http.NewRequest("POST", URL.getInfoByTimeId, bytes.NewBuffer(ParamsJSON))
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
	mobileLoginResp := new(MobileLoginResp)
	json.Unmarshal(ResponseBody, &mobileLoginResp)
	newJson := RegReplace(RegDelete(mobileLoginResp.D.Data, `(\r|\n)`), `"`, `\"`)
	getInfoByTimeIdResp = new(GetInfoByTimeIdResp)
	json.Unmarshal([]byte(newJson), &getInfoByTimeIdResp)
	Store.CoachID = getInfoByTimeIdResp.BusinessData.CoachInfo.Data[0].CoachID
	return getInfoByTimeIdResp
}

func (Store *Store) GetTimesInfoByCoachIDNew() (getTimesInfoByCoachIDNewResp *GetTimesInfoByCoachIDNewResp) {
	Params := GetTimesInfoByCoachIDNewParams{
		CoachID: Store.CoachID,
		Date:    fmt.Sprint(time.Now().Year(), "-", int(time.Now().Month()), "-", time.Now().Day()+1),
		Subid:   0,
	}
	ParamsJSON, err := json.Marshal(Params)
	if err != nil {
		logrus.Error(err)
	}
	HTTPClient := &http.Client{
		Transport: &HTTPTransport,
	}
	Request, err := http.NewRequest("POST", URL.getTimesInfoByCoachIDNew, bytes.NewBuffer(ParamsJSON))
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
	dresp := new(Dresp)
	json.Unmarshal(ResponseBody, &dresp)
	newJson := RegReplace(RegDelete(dresp.D, `(\r|\n)`), `"`, `\"`)
	getTimesInfoByCoachIDNewResp = new(GetTimesInfoByCoachIDNewResp)
	json.Unmarshal([]byte(newJson), &getTimesInfoByCoachIDNewResp)
	Store.Data = getTimesInfoByCoachIDNewResp.Data
	return getTimesInfoByCoachIDNewResp
}

func (Store *Store) OrderCoachNew() (orderCoachNewResp *OrderCoachNewResp) {
	Params := OrderCoachNewParams{
		PXResID: "",
		SubID:   strconv.Itoa(JSONConfig().Project),
		StateID: Store.StateID,
	}
	ParamsJSON, err := json.Marshal(Params)
	if err != nil {
		logrus.Error(err)
	}
	HTTPClient := &http.Client{
		Transport: &HTTPTransport,
	}
	Request, err := http.NewRequest("POST", URL.orderCoachNew, bytes.NewBuffer(ParamsJSON))
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
	orderCoachNewResp = new(OrderCoachNewResp)
	json.Unmarshal(ResponseBody, &orderCoachNewResp)
	return orderCoachNewResp
}
