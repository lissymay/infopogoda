package main

import (
	"os"
	"time"

	"github.com/lissymay/infopogoda.git/internal/adapters/weather"
	"github.com/lissymay/infopogoda.git/internal/pkg/app/cli"
	"github.com/lissymay/infopogoda.git/internal/pkg/cache"
	"github.com/lissymay/infopogoda.git/internal/pkg/flags"
	"github.com/lissymay/infopogoda.git/pkg/config"
	"github.com/lissymay/infopogoda.git/pkg/logger"
)

func main() {
	// Парсим аргументы командной строки
	arguments := flags.Parse()

	// Открываем конфигурационный файл
	r, err := os.Open(arguments.Path)
	if err != nil {
		panic(err)
	}

	// Парсим конфигурацию
	c, err := config.Parse(r)
	if err != nil {
		panic(err)
	}

	log := logger.New()

	// Инициализируем кэш в зависимости от типа
	var cacheInstance cache.Cache
	cacheTTL := time.Duration(c.C.TTL) * time.Second

	switch c.C.Type {
	case "memory":
		cacheInstance = cache.NewMemoryCache(cacheTTL)
		log.Info("✅ Используется in-memory кэш")
	case "redis":
		redisCache, err := cache.NewRedisCache(cacheTTL, c.C.RedisAddr, c.C.RedisPassword, c.C.RedisDB)
		if err != nil {
			log.Error("❌ Ошибка подключения к Redis: " + err.Error())
			log.Info("Переключаемся на in-memory кэш")
			cacheInstance = cache.NewMemoryCache(cacheTTL)
		} else {
			cacheInstance = redisCache
			log.Info("✅ Используется Redis кэш (адрес: " + c.C.RedisAddr + ")")
			defer redisCache.Close()
		}
	default:
		cacheInstance = cache.NewMemoryCache(cacheTTL)
		log.Info("⚠️  Кэш не настроен, используется in-memory")
	}

	wi := getProvider(c, log, cacheInstance, cacheTTL)

	app := cli.New(log, wi, c)

	err = app.Run()
	if err != nil {
		log.Error("Ошибка выполнения приложения: " + err.Error())
		os.Exit(1)
	}

	log.Info("Приложение завершено успешно")
	os.Exit(0)
}

func getProvider(c config.Config, l cli.Logger, cacheInstance cache.Cache, ttl time.Duration) cli.WeatherInfo {
	var wi cli.WeatherInfo
	switch c.P.Type {
	case "open-meteo":
		wi = weather.New(l, cacheInstance, ttl)
	default:
		wi = weather.New(l, cacheInstance, ttl)
	}
	return wi
}
