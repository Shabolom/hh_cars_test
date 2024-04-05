package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"hh_test_autho/config"
	"hh_test_autho/internal/routes"
	"hh_test_autho/internal/tools"

	_ "hh_test_autho/docs"
)

func main() {
	//	@title		Cars_API
	//	@version	1.0.0

	// 	@description 	Это hh_cars_project с использованием свагера
	// 	@termsOfService  тут были-бы условия использования, еслибы я их мог обозначить
	// 	@contact.url    https://t.me/Timuchin3
	// 	@contact.email  tima.gorenskiy@mail.ru

	//	@host		localhost:8800

	// host необходимо передать свой для корректной работы swagger

	// заполняем структуру в Env
	config.CheckFlagEnv()

	err := tools.InitLogger()
	if err != nil {
		fmt.Println(err)
	}

	// инициализируем инфо логи с измененным выводом
	var logInf = tools.InfoLogs()

	// config.InitPgSQL инициализируем подключение к базе данных
	err = config.InitPgSQL()
	if err != nil {
		log.WithField("component", "initialization").Fatal(err)
	}
	log.WithField("component", "initialization").Info("подключились к базе")

	// конфигурация (инициализация) end point
	r := routes.SetupRouter()

	// запуск сервера
	logInf.WithField("component", "initialization").Info("запускаем сервер")
	if err = r.Run(config.Env.Host + ":" + config.Env.Port); err != nil {
		log.WithField("component", "run").Fatal(err)
	}
}
