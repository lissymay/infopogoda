package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/lissymay/infopogoda.git/pkg/logger"
)

type cliApp struct {
	log logger.Logger
}

func New(log logger.Logger) *cliApp {
	return &cliApp{log: log}
}

func (c *cliApp) Run() error {
	c.log.Info("Формируем параметры запроса к open-meteo")

	type Current struct {
		Temp float32 `json:"temperature_2m"`
	}

	type Response struct {
		Curr Current `json:"current"`
	}

	var response Response

	params := fmt.Sprintf(
		"latitude=%f&longitude=%f&current=temperature_2m",
		53.6688,
		23.8223,
	)

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s", params)

	c.log.Info("Отправляем HTTP запрос к open-meteo")

	resp, err := http.Get(url)
	if err != nil {
		customErr := errors.New("can't get weather data from openmeteo")
		c.log.Error(customErr.Error())
		return errors.Join(customErr, err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.log.Error("can't close response body: " + err.Error())
		}
	}()

	c.log.Debug("Читаем тело ответа")

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		customErr := errors.New("can't read data from response")
		c.log.Error(customErr.Error())
		return errors.Join(customErr, err)
	}

	c.log.Debug("Парсим JSON")

	if err := json.Unmarshal(data, &response); err != nil {
		customErr := errors.New("can't unmarshal data from response")
		c.log.Error(customErr.Error())
		return errors.Join(customErr, err)
	}

	c.log.Info("Данные успешно получены")

	fmt.Printf(
		"Температура воздуха - %.2f градусов цельсия\n",
		response.Curr.Temp,
	)

	return nil
}
