package cli

import (
	"testing"

	"github.com/lissymay/infopogoda.git/internal/domain/models"
	"github.com/lissymay/infopogoda.git/pkg/config"
)

// Мок для WeatherInfo
type mockWeatherInfo struct {
	temp float32
}

func (m *mockWeatherInfo) GetTemperature(lat, long float64) models.TempInfo {
	return models.TempInfo{Temp: m.temp}
}

// Мок для логгера
type testLogger struct{}

func (l *testLogger) Info(msg string)  {}
func (l *testLogger) Debug(msg string) {}
func (l *testLogger) Error(msg string) {}

func TestCliApp_Run(t *testing.T) {
	logger := &testLogger{}
	mockWI := &mockWeatherInfo{temp: 22.5}

	cfg := config.Config{
		L: config.Location{
			Lat:  53.6688,
			Long: 23.8223,
		},
	}

	app := New(logger, mockWI, cfg)

	// Просто проверяем, что Run не паникует
	err := app.Run()
	if err != nil {
		t.Errorf("Run() вернул ошибку: %v", err)
	}
}
