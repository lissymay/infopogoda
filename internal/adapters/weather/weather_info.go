package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/lissymay/infopogoda.git/internal/domain/models"
)

const apiURL = "https://api.open-meteo.com/v1/forecast"

type Logger interface {
	Info(string)
	Debug(string)
	Error(string, error)
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

	resp, err := http.Get(url)
	if err != nil {
		wi.l.Error("failed to get weather data", err)
		customErr := errors.New("failed to get data from openmeteo")
		return errors.Join(customErr, err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			wi.l.Error("failed to close response body", err)
		}
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		wi.l.Error("failed to read response body", err)
		customErr := errors.New("failed to read data from response")
		return errors.Join(customErr, err)
	}

	if err := json.Unmarshal(data, &respData); err != nil {
		wi.l.Error("failed to unmarshal JSON", err)
		customErr := errors.New("failed to unmarshal data from response")
		return errors.Join(customErr, err)
	}

	wi.current = respData.Curr
	wi.isLoaded = true
	return nil
}

func (wi *weatherInfo) GetTemperature(lat, long float64) (models.TempInfo, error) {
	err := wi.getWeatherInfo(lat, long)
	return models.TempInfo{
		Temp: wi.current.Temp,
	}, err
}
