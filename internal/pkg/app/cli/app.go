package cli

import (
	"fmt"

	"github.com/lissymay/infopogoda.git/internal/domain/models"
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
	l  Logger
	wi WeatherInfo
}

func New(l Logger, wi WeatherInfo) *cliApp {
	return &cliApp{
		l:  l,
		wi: wi,
	}
}

func (c *cliApp) Run() error {
	// Координаты Гродно
	lat := 53.6688
	long := 23.8223

	c.l.Info("Запрашиваем данные о погоде")
	tempInfo := c.wi.GetTemperature(lat, long)

	fmt.Printf(
		"Температура воздуха - %.2f градусов цельсия\n",
		tempInfo.Temp,
	)

	return nil
}
