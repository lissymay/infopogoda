package providers

import (
	pogodaby "github.com/lissymay/infopogoda.git/internal/adapters/pogoda_by"
	"github.com/lissymay/infopogoda.git/internal/adapters/weather"
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

func GetProvider(c config.Config, l Logger) WeatherInfo {
	switch c.P.Type {
	case "open-meteo":
		return weather.New(l)
	case "pogoda":
		return pogodaby.New(l)
	default:
		return weather.New(l)
	}
}
