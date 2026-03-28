package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/lissymay/infopogoda.git/internal/domain/models"
	"github.com/lissymay/infopogoda.git/internal/pkg/cache"
)

const apiURL = "https://api.open-meteo.com/v1/forecast"

type Logger interface {
	Info(string)
	Debug(string)
	Error(string)
}

type current struct {
	Temp float32 `json:"temperature_2m"`
}

type response struct {
	Curr current `json:"current"`
}

type weatherInfo struct {
	l        Logger
	cache    cache.Cache
	cacheTTL time.Duration
	current  current
}

func New(l Logger, c cache.Cache, ttl time.Duration) *weatherInfo {
	return &weatherInfo{
		l:        l,
		cache:    c,
		cacheTTL: ttl,
	}
}

func (wi *weatherInfo) getWeatherInfo(lat, long float64) error {
	var respData response

	params := fmt.Sprintf(
		"latitude=%f&longitude=%f&current=temperature_2m",
		lat,
		long,
	)

	url := fmt.Sprintf("%s?%s", apiURL, params)

	wi.l.Debug(fmt.Sprintf("URL: %s", url))

	resp, err := http.Get(url)
	if err != nil {
		wi.l.Error("не удалось получить данные о погоде")
		return errors.Join(errors.New("не удалось получить данные от openmeteo"), err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		wi.l.Error("не удалось прочитать тело ответа")
		return errors.Join(errors.New("не удалось прочитать данные из ответа"), err)
	}

	if err := json.Unmarshal(data, &respData); err != nil {
		wi.l.Error("не удалось распарсить JSON")
		return errors.Join(errors.New("не удалось распарсить данные из ответа"), err)
	}

	wi.current = respData.Curr

	cacheKey := fmt.Sprintf("weather:%f:%f", lat, long)
	if wi.cache != nil {
		wi.cache.Set(cacheKey, respData.Curr.Temp, wi.cacheTTL)
	}

	return nil
}

func (wi *weatherInfo) GetTemperature(lat, long float64) models.TempInfo {
	cacheKey := fmt.Sprintf("weather:%f:%f", lat, long)

	if wi.cache != nil {
		if cached, found := wi.cache.Get(cacheKey); found {
			if temp, ok := cached.(float32); ok {
				wi.l.Info("✅ Данные из кэша")
				return models.TempInfo{Temp: temp}
			}
		}
	}

	wi.l.Info("🔄 Загружаем из API")
	if err := wi.getWeatherInfo(lat, long); err != nil {
		wi.l.Error("ошибка при загрузке данных: " + err.Error())
		return models.TempInfo{Temp: 0}
	}

	return models.TempInfo{Temp: wi.current.Temp}
}
