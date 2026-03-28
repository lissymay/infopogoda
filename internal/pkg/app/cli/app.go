package cli

import (
	"fmt"

	"github.com/lissymay/infopogoda.git/internal/domain/models"
	"github.com/lissymay/infopogoda.git/pkg/config"
)

type Logger interface {
	Info(string)
	Debug(string)
	Error(string)
}

type WeatherInfo interface {
	GetTemperature(float64, float64) models.TempInfo
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
	// Используем координаты из конфигурации
	lat := c.cfg.L.Lat
	long := c.cfg.L.Long

	c.l.Info(fmt.Sprintf("Запрашиваем данные о погоде для координат: %f, %f", lat, long))
	tempInfo := c.wi.GetTemperature(lat, long)

	fmt.Printf(
		"Температура воздуха - %.2f градусов цельсия\n",
		tempInfo.Temp,
	)

	return nil
}
