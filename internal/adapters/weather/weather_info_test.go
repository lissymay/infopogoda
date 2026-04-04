package weather

import (
	"testing"
	"time"
)

// Мок для логгера
type mockLogger struct {
	infoMessages  []string
	debugMessages []string
	errorMessages []string
}

func (m *mockLogger) Info(msg string) {
	m.infoMessages = append(m.infoMessages, msg)
}

func (m *mockLogger) Debug(msg string) {
	m.debugMessages = append(m.debugMessages, msg)
}

func (m *mockLogger) Error(msg string) {
	m.errorMessages = append(m.errorMessages, msg)
}

// Мок для кэша
type mockCache struct {
	data map[string]interface{}
}

func newMockCache() *mockCache {
	return &mockCache{
		data: make(map[string]interface{}),
	}
}

func (m *mockCache) Get(key string) (interface{}, bool) {
	val, ok := m.data[key]
	return val, ok
}

func (m *mockCache) Set(key string, value interface{}, ttl time.Duration) {
	m.data[key] = value
}

func (m *mockCache) Delete(key string) {
	delete(m.data, key)
}

func (m *mockCache) Clear() {
	m.data = make(map[string]interface{})
}

func TestGetTemperature_FromCache(t *testing.T) {
	logger := &mockLogger{}
	cache := newMockCache()

	lat := 53.6688
	long := 23.8223
	cacheKey := "weather:53.668800:23.822300"

	// Предварительно сохраняем в кэш
	cache.Set(cacheKey, float32(25.5), 0)

	wi := New(logger, cache, 5*time.Second)

	tempInfo := wi.GetTemperature(lat, long)

	if tempInfo.Temp != 25.5 {
		t.Errorf("GetTemperature() = %v, ожидалось 25.5", tempInfo.Temp)
	}
}

func TestGetTemperature_FromAPI(t *testing.T) {
	logger := &mockLogger{}
	cache := newMockCache()

	lat := 53.6688
	long := 23.8223

	wi := New(logger, cache, 5*time.Second)

	tempInfo := wi.GetTemperature(lat, long)

	// Проверяем, что температура вернулась (не 0)
	if tempInfo.Temp == 0 {
		t.Errorf("GetTemperature() вернул 0, ожидалась реальная температура")
	}

	// Проверяем, что данные сохранились в кэш
	cacheKey := "weather:53.668800:23.822300"
	_, found := cache.Get(cacheKey)
	if !found {
		t.Errorf("Данные не сохранились в кэш после запроса к API")
	}
}

func TestGetTemperature_NoCache(t *testing.T) {
	logger := &mockLogger{}

	wi := New(logger, nil, 5*time.Second) // cache = nil

	lat := 53.6688
	long := 23.8223

	tempInfo := wi.GetTemperature(lat, long)

	// Должно работать даже без кэша
	if tempInfo.Temp == 0 {
		t.Errorf("GetTemperature() без кэша вернул 0")
	}
}
