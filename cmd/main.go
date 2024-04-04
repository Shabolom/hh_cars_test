package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"hh_test_autho/config"
	"hh_test_autho/internal/routes"
	"hh_test_autho/internal/tools"
)

func main() {
	config.CheckFlagEnv()
	err := tools.InitLogger()
	if err != nil {
		fmt.Println(err)
	}

	// config.InitPgSQL инициализируем подключение к базе данных
	err = config.InitPgSQL()
	if err != nil {
		log.WithField("component", "initialization").Fatal(err)
	}

	// конфигурация (инициализация) end point
	r := routes.SetupRouter()

	// запуск сервера
	if err = r.Run(config.Env.Host + ":" + config.Env.Port); err != nil {
		log.WithField("component", "run").Fatal(err)
	}
}
