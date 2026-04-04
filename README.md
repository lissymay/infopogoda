# Weather Info CLI

[![CI](https://github.com/lissymay/infopogoda.git/actions/workflows/ci.yml/badge.svg)](https://github.com/lissymay/infopogoda.git/actions/workflows/ci.yml)
[![Tests](https://github.com/lissymay/infopogoda.git/actions/workflows/test.yml/badge.svg)](https://github.com/lissymay/infopogoda.git/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lissymay/infopogoda.git)](https://goreportcard.com/report/github.com/lissymay/infopogoda.git)

## CLI приложение для получения погоды

### Возможности
- Получение текущей температуры по координатам
- Кэширование (in-memory и Redis)
- Конфигурация через YAML
- Поддержка аргументов командной строки

### Запуск
```bash
make run          # In-memory кэш
make run-redis    # Redis кэш
make test         # Запуск тестов