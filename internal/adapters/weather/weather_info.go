package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/lissymay/infopogoda.git/internal/domain/models" // используйте ваш модуль
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
	current  current
	isLoaded bool
}

func New(l Logger) *weatherInfo {
	return &weatherInfo{
		l: l,
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
	return nil
}

func (wi *weatherInfo) GetTemperature(lat, long float64) models.TempInfo {
	if !wi.isLoaded {
		if err := wi.getWeatherInfo(lat, long); err != nil {
			wi.l.Error("ошибка при загрузке данных о погоде: " + err.Error())
		}
	}
	return models.TempInfo{
		Temp: wi.current.Temp,
	}
}
