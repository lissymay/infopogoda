package main

import (
	"os"

	"github.com/lissymay/infopogoda.git/internal/adapters/weather"
	"github.com/lissymay/infopogoda.git/internal/pkg/app/cli"
	"github.com/lissymay/infopogoda.git/internal/pkg/flags"
	"github.com/lissymay/infopogoda.git/pkg/config"
	"github.com/lissymay/infopogoda.git/pkg/logger"
)

func main() {
	// Парсим аргументы командной строки
	arguments := flags.Parse()

	// Открываем конфигурационный файл
	r, err := os.Open(arguments.Path)
	if err != nil {
		panic(err)
	}

	// Парсим конфигурацию
	c, err := config.Parse(r)
	if err != nil {
		panic(err)
	}

	log := logger.New()

	wi := getProvider(c, log)

	app := cli.New(log, wi, c)

	err = app.Run()
	if err != nil {
		log.Error("Ошибка выполнения приложения: " + err.Error())
		os.Exit(1)
	}

	log.Info("Приложение завершено успешно")
	os.Exit(0)
}

func getProvider(c config.Config, l cli.Logger) cli.WeatherInfo {
	var wi cli.WeatherInfo
	switch c.P.Type {
	case "open-meteo":
		wi = weather.New(l)
	default:
		wi = weather.New(l)
	}
	return wi
}
