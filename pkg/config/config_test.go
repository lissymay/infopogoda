package config

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	yamlData := `
service:
  provider:
    type: open-meteo
  location:
    lat: 53.6688
    long: 23.8223
  cache:
    type: memory
    ttl: 300
`

	reader := strings.NewReader(yamlData)
	cfg, err := Parse(reader)

	if err != nil {
		t.Errorf("Parse() вернул ошибку: %v", err)
	}

	if cfg.P.Type != "open-meteo" {
		t.Errorf("Provider.Type = %v, ожидалось open-meteo", cfg.P.Type)
	}

	if cfg.L.Lat != 53.6688 {
		t.Errorf("Location.Lat = %v, ожидалось 53.6688", cfg.L.Lat)
	}

	if cfg.L.Long != 23.8223 {
		t.Errorf("Location.Long = %v, ожидалось 23.8223", cfg.L.Long)
	}

	if cfg.C.Type != "memory" {
		t.Errorf("Cache.Type = %v, ожидалось memory", cfg.C.Type)
	}

	if cfg.C.TTL != 300 {
		t.Errorf("Cache.TTL = %v, ожидалось 300", cfg.C.TTL)
	}
}

func TestParse_RedisConfig(t *testing.T) {
	yamlData := `
service:
  provider:
    type: open-meteo
  location:
    lat: 53.6688
    long: 23.8223
  cache:
    type: redis
    ttl: 600
    redis_addr: "localhost:6379"
    redis_password: ""
    redis_db: 0
`

	reader := strings.NewReader(yamlData)
	cfg, err := Parse(reader)

	if err != nil {
		t.Errorf("Parse() вернул ошибку: %v", err)
	}

	if cfg.C.Type != "redis" {
		t.Errorf("Cache.Type = %v, ожидалось redis", cfg.C.Type)
	}

	if cfg.C.TTL != 600 {
		t.Errorf("Cache.TTL = %v, ожидалось 600", cfg.C.TTL)
	}

	if cfg.C.RedisAddr != "localhost:6379" {
		t.Errorf("RedisAddr = %v, ожидалось localhost:6379", cfg.C.RedisAddr)
	}
}

func TestParse_InvalidYAML(t *testing.T) {
	invalidYAML := `
service:
  provider:
    type: open-meteo
  location:
    lat: invalid_number
`

	reader := strings.NewReader(invalidYAML)
	_, err := Parse(reader)

	if err == nil {
		t.Errorf("Parse() с невалидным YAML не вернул ошибку")
	}
}