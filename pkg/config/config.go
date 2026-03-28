package config

import (
	"io"

	"gopkg.in/yaml.v3"
)

type ConfigFile struct {
	C Config `yaml:"service"`
}

type Provider struct {
	Type string `yaml:"type"`
}

type Location struct {
	Lat  float64 `yaml:"lat"`
	Long float64 `yaml:"long"`
}

type CacheConfig struct {
	Type string `yaml:"type"` // "memory" или "redis"
	TTL  int    `yaml:"ttl"`  // время жизни кэша в секундах
	// Для Redis
	RedisAddr     string `yaml:"redis_addr,omitempty"`
	RedisPassword string `yaml:"redis_password,omitempty"`
	RedisDB       int    `yaml:"redis_db,omitempty"`
}

type Config struct {
	P Provider    `yaml:"provider"`
	L Location    `yaml:"location"`
	C CacheConfig `yaml:"cache"`
}

func Parse(r io.Reader) (Config, error) {
	var c ConfigFile
	if err := yaml.NewDecoder(r).Decode(&c); err != nil {
		return Config{}, err
	}
	return c.C, nil
}
