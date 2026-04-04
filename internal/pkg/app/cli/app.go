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
<<<<<<< HEAD
	l    Logger
	wi   WeatherInfo
	conf config.Config
}

func New(l Logger, wi WeatherInfo, c config.Config) *cliApp {
	return &cliApp{
		l:    l,
		wi:   wi,
		conf: c,
=======
	l   Logger
	wi  WeatherInfo
	cfg config.Config
}

func New(l Logger, wi WeatherInfo, cfg config.Config) *cliApp {
	return &cliApp{
		l:   l,
		wi:  wi,
		cfg: cfg,
>>>>>>> aecab4d0c3cc7544596581e380fd0a9f06d16b88
	}
}

func (c *cliApp) Run() error {
<<<<<<< HEAD
	tempInfo, err := c.wi.GetTemperature(c.conf.L.Lat, c.conf.L.Long)
	if err != nil {
		c.l.Error("can't get temperature info", err)
		return err
	}
=======
	// Используем координаты из конфигурации
	lat := c.cfg.L.Lat
	long := c.cfg.L.Long

	c.l.Info(fmt.Sprintf("Запрашиваем данные о погоде для координат: %f, %f", lat, long))
	tempInfo := c.wi.GetTemperature(lat, long)
>>>>>>> aecab4d0c3cc7544596581e380fd0a9f06d16b88

	fmt.Printf(
		"Temperature - %.2f degrees Celsius\n",
		tempInfo.Temp,
	)

	return nil
}
