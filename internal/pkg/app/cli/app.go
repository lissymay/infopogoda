package cli

import (
	"fmt"

	"github.com/lissymay/infopogoda.git/internal/domain/models"
	"github.com/lissymay/infopogoda.git/pkg/config"
)

type Logger interface {
	Info(string)
	Debug(string)
	Error(string, error)
}

type WeatherInfo interface {
	GetTemperature(float64, float64) (models.TempInfo, error)
}

type cliApp struct {
	l   Logger
	wi  WeatherInfo
	cfg config.Config
}

func New(l Logger, wi WeatherInfo, cfg config.Config) *cliApp {
	return &cliApp{
		l:   l,
		wi:  wi,
		cfg: cfg, 
	}
}

func (c *cliApp) Run() error {
	tempInfo, err := c.wi.GetTemperature(c.cfg.L.Lat, c.cfg.L.Long)
	if err != nil {
		c.l.Error("can't get temperature info", err)
		return err
	}

	fmt.Printf(
		"Temperature - %.2f degrees Celsius\n",
		tempInfo.Temp,
	)

	return nil
}
