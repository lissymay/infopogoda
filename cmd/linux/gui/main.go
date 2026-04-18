package main

import (
	"os"

	"github.com/lissymay/infopogoda.git/internal/pkg/app/gui"
	"github.com/lissymay/infopogoda.git/internal/pkg/flags"
	"github.com/lissymay/infopogoda.git/internal/pkg/gui/fyne"
	"github.com/lissymay/infopogoda.git/internal/pkg/providers"
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

	// Get weather provider
	provider := providers.GetProvider(c, log)

	// Create Fyne provider
	p := fyne.NewP()

	// Create and run GUI app
	g := gui.New(log, p, provider, c)

	err = g.Run()
	if err != nil {
		panic(err)
	}
}
