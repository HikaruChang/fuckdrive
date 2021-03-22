package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var LogFileName = flag.String("log", fmt.Sprint(RunPath(), "fuckdrive.log"), "Log file name")

var (
	URL = URLConst{
		base:                     "http://yyhzxcjx.91jiexun.com",
		mobileLogin:              "http://yyhzxcjx.91jiexun.com/Server/AccountServer.asmx/MobileLogin",
		appType:                  "http://yyhzxcjx.91jiexun.com/Account/AppType",
		isLogin:                  "http://yyhzxcjx.91jiexun.com/Server/AlipayServerNew.asmx/IsLogin",
		getUserInfo:              "http://yyhzxcjx.91jiexun.com/Server/BaseInfoServer.asmx/GetUserInfo",
		getDaysOrder:             "http://yyhzxcjx.91jiexun.com/Server/OrderCoachServer.asmx/getDaysOrder",
		getInfoByTimeId:          "http://yyhzxcjx.91jiexun.com/Server/OrderCoachServer.asmx/GetInfoByTimeId",
		getTimesInfoByCoachIDNew: "http://yyhzxcjx.91jiexun.com/Server/OrderCoachServer.asmx/GetTimesInfoByCoachIDNew",
		orderCoachNew:            "http://yyhzxcjx.91jiexun.com/Server/OrderCoachServer.asmx/orderCoachNew",
	}

	Client = ClientConst{
		accept:      "application/json",
		contentType: "application/json; charset=UTF-8",
		host:        "yyhzxcjx.91jiexun.com",
		origin:      "http://yyhzxcjx.91jiexun.com",
		referer:     "http://yyhzxcjx.91jiexun.com/NMobile/page/time.html?version=20160927",
		userAgent:   "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
	}

	HTTPTransport = http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, time.Second*60) //设置建立连接超时
			if err != nil {
				logrus.Error(err)
			}
			conn.SetDeadline(time.Now().Add(time.Second * 300)) //设置发送接受数据超时
			return conn, nil
		},
		ResponseHeaderTimeout: time.Second * 60,
		DisableKeepAlives:     false,
	}
)
