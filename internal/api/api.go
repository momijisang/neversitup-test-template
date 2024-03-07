package api

import (
	"fmt"
	"neversitup-test-template/internal/api/router"
	"neversitup-test-template/internal/pkg/MongoDB"
	"neversitup-test-template/internal/pkg/config"
	"neversitup-test-template/internal/pkg/persistence"
)

func setUp(configPath string) {
	config.Setup(configPath)
	//DB.Setup()
}

func Run() {
	setUp("data/config.yml")

	conf := config.GetConfig()
	web := router.Setup()
	MongoDB.Db = MongoDB.Db.Connect(config.Config.Mongo)

	schedule := persistence.Schedule()
	go schedule.SetSchedule()

	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")

	_ = web.Run(":" + conf.Server.Port)

}
