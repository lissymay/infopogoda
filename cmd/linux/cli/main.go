package main

import (
	"os"

	"github.com/lissymay/infopogoda.git/internal/adapters/weather"
	"github.com/lissymay/infopogoda.git/internal/pkg/app/cli"
	"github.com/lissymay/infopogoda.git/pkg/logger"
)

func main() {
	log := logger.New()

	wi := weather.New(log)

	app := cli.New(log, wi)

	err := app.Run()
	if err != nil {
		log.Error("Ошибка выполнения приложения: " + err.Error())
		os.Exit(1)
	}

	log.Info("Приложение завершено успешно")
	os.Exit(0)
}
