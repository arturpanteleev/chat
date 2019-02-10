package config

import (
	"encoding/json"
	"io/ioutil"
)

type AppConfig struct {
	ApiPort string `json:"api_port"`

	MysqlHost   string `json:"mysql_host"`
	MysqlUser   string `json:"mysql_user"`
	MysqlPass   string `json:"mysql_pass"`
	MysqlDbName string `json:"mysql_dbName"`

	CentrifugoApiUrl string `json:"centrifugo_api_url"`
}

func Config() *AppConfig {
	var conf AppConfig
	plan, _ := ioutil.ReadFile("conf.json")
	err := json.Unmarshal(plan, &conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
