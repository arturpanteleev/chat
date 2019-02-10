package main

import (
	"chat/api"
	"chat/config"
	"chat/delivery"
	"chat/storage"
	"github.com/labstack/echo"
)

func main() {

	conf := config.Config()

	chatApi := api.New(
		conf.ApiPort,
		echo.New(),
		storage.NewMysqlStorage(conf.MysqlHost, conf.MysqlUser, conf.MysqlPass, conf.MysqlDbName),
		delivery.NewCentrifugoDelivery(conf.CentrifugoApiUrl),
	)

	chatApi.Start()
}
