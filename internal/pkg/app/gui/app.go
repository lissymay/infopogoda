package gui

import (
	"fmt"

	guisettings "github.com/lissymay/infopogoda.git/internal/domain/gui_settings"
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

type guiApp struct {
	l          Logger
	p          guisettings.AppProvider
	wi         WeatherInfo
	cfg        config.Config
	window     guisettings.Window
	textWidget guisettings.TextWidget
}

func New(l Logger, p guisettings.AppProvider, wi WeatherInfo, cfg config.Config) *guiApp {
	return &guiApp{
		l:   l,
		p:   p,
		wi:  wi,
		cfg: cfg,
	}
}

func (g *guiApp) Run() error {
	// Создаем окно
	window, err := g.p.CreateWindow("Weather Info", guisettings.NewWS(400, 300))
	if err != nil {
		return err
	}
	g.window = window

	// Создаем текстовый виджет
	g.textWidget = g.p.GetTextWidget("Loading...")
	g.window.SetTemperatureWidget(g.textWidget)

	// Получаем температуру
	tempInfo, err := g.wi.GetTemperature(g.cfg.L.Lat, g.cfg.L.Long)
	if err != nil {
		g.l.Error("Failed to get temperature", err)
		g.textWidget.SetText("Error: " + err.Error())
	} else {
		g.textWidget.SetText("Temperature: " + formatTemperature(tempInfo.Temp))
	}

	// Запускаем приложение
	runner := g.p.GetAppRunner()
	runner.Run()

	return nil
}

func formatTemperature(t float32) string {
	return fmt.Sprintf("%.2f°C", t)
}
