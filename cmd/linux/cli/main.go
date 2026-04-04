package main

import (
	"os"

	pogodaby "github.com/lissymay/infopogoda.git/internal/adapters/pogoda_by"
	"github.com/lissymay/infopogoda.git/internal/adapters/weather"
	"github.com/lissymay/infopogoda.git/internal/pkg/app/cli"
	"github.com/lissymay/infopogoda.git/internal/pkg/flags"
	"github.com/lissymay/infopogoda.git/pkg/config"
	"github.com/lissymay/infopogoda.git/pkg/logger"
)

func main() {
<<<<<<< HEAD
	// Parse command line arguments
	arguments := flags.Parse()

	// Open config file
=======
	// Парсим аргументы командной строки
	arguments := flags.Parse()

	// Открываем конфигурационный файл
>>>>>>> aecab4d0c3cc7544596581e380fd0a9f06d16b88
	r, err := os.Open(arguments.Path)
	if err != nil {
		panic(err)
	}

<<<<<<< HEAD
	// Parse config
=======
	// Парсим конфигурацию
>>>>>>> aecab4d0c3cc7544596581e380fd0a9f06d16b88
	c, err := config.Parse(r)
	if err != nil {
		panic(err)
	}

	log := logger.New()

	wi := getProvider(c, log)

	app := cli.New(log, wi, c)

	err = app.Run()
	if err != nil {
		log.Error("Application error", err)
		os.Exit(1)
	}

	log.Info("Application completed successfully")
	os.Exit(0)
}

func getProvider(c config.Config, l cli.Logger) cli.WeatherInfo {
	var wi cli.WeatherInfo
	switch c.P.Type {
	case "open-meteo":
		wi = weather.New(l)
<<<<<<< HEAD
	case "pogoda":
		wi = pogodaby.New(l)
=======
>>>>>>> aecab4d0c3cc7544596581e380fd0a9f06d16b88
	default:
		wi = weather.New(l)
	}
	return wi
}
