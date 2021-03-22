package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// JSONConfig 读取json配置文件
func JSONConfig() (config *Config) {
	configFile, err := ioutil.ReadFile(fmt.Sprint(RunPath(), "config.json"))
	if err != nil {
		logrus.Error(err)
	}
	config = &Config{}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		logrus.Error(err)
	}
	return config
}

// JSONEdit 设置json
func JSONEdit(config *Config) (err error) {
	data, err := json.MarshalIndent(config, "", "	")
	if err != nil {
		logrus.Error(err)
	}
	err = ioutil.WriteFile(fmt.Sprint(RunPath(), "config.json"), data, 0644)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return
}
