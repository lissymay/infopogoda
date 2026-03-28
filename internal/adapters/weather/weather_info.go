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
	isLoaded bool
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

	wi.l.Debug(fmt.Sprintf("URL сгенерирован: %s", url))

	resp, err := http.Get(url)
	if err != nil {
		wi.l.Error("не удалось получить данные о погоде")
		return errors.Join(errors.New("не удалось получить данные от openmeteo"), err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			wi.l.Error("не удалось закрыть тело ответа: " + err.Error())
		}
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		wi.l.Error("не удалось прочитать тело ответа")
		return errors.Join(errors.New("не удалось прочитать данные из ответа"), err)
	}

	wi.l.Debug(fmt.Sprintf("Данные успешно прочитаны, размер: %d байт", len(data)))

	if err := json.Unmarshal(data, &respData); err != nil {
		wi.l.Error("не удалось распарсить JSON")
		return errors.Join(errors.New("не удалось распарсить данные из ответа"), err)
	}

	wi.current = respData.Curr
	wi.isLoaded = true

	// Сохраняем в кэш
	cacheKey := fmt.Sprintf("weather:%f:%f", lat, long)
	if wi.cache != nil {
		wi.cache.Set(cacheKey, respData.Curr, wi.cacheTTL)
		wi.l.Debug(fmt.Sprintf("Данные сохранены в кэш с ключом: %s", cacheKey))
	}

	return nil
}

func (wi *weatherInfo) GetTemperature(lat, long float64) models.TempInfo {
	cacheKey := fmt.Sprintf("weather:%f:%f", lat, long)

	// Пытаемся получить из кэша
	if wi.cache != nil {
		if cached, found := wi.cache.Get(cacheKey); found {
			if curr, ok := cached.(current); ok {
				wi.l.Info(fmt.Sprintf("✅ Данные получены из кэша для ключа: %s", cacheKey))
				return models.TempInfo{
					Temp: curr.Temp,
				}
			}
		}
	}

	// Если в кэше нет, загружаем из API
	wi.l.Info(fmt.Sprintf("🔄 Кэш не найден для ключа: %s, загружаем из API", cacheKey))
	if err := wi.getWeatherInfo(lat, long); err != nil {
		wi.l.Error("ошибка при загрузке данных о погоде: " + err.Error())
		return models.TempInfo{Temp: 0}
	}

	return models.TempInfo{
		Temp: wi.current.Temp,
	}
}
