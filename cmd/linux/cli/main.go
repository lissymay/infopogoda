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
	// Parse command line arguments
	arguments := flags.Parse()

	// Open config file
	r, err := os.Open(arguments.Path)
	if err != nil {
		panic(err)
	}

	// Parse config
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
	case "pogoda":
		wi = pogodaby.New(l)
	default:
		wi = weather.New(l)
	}
	return wi
}
