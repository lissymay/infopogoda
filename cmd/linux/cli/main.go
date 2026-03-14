package main

import (
	"os"

	"github.com/lissymay/infopogoda.git/internal/pkg/app/cli"
	"github.com/lissymay/infopogoda.git/pkg/logger"
)

func main() {
	log := logger.New()
	app := cli.New(log)

	err := app.Run()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Info("Приложение завершено успешно")
	os.Exit(0)
}
