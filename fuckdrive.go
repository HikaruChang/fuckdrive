package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	logFile, logErr := os.OpenFile(*LogFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		logFile, err := os.Create(*LogFileName)
		if err != nil {
			fmt.Println("Fail to find", *logFile, "PorterMan start Failed")
		}
		defer logFile.Close()
	}
	defer logFile.Close()
	logrus.SetOutput(logFile)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	Drive := new(Store)
	Drive.OrderNo = 0
	if Drive.isLogin().D.Data == "0" {
		Drive.appType()
		Drive.mobileLogin()
	}
	Drive.GetDaysOrder()
	Drive.GetInfoByTimeId()
	Drive.GetTimesInfoByCoachIDNew()
	Drive.OrderQueue()
}
