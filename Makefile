.PHONY: run run-redis help

run:
	go run ./cmd/linux/cli/main.go

run-redis:
	go run ./cmd/linux/cli/main.go -config config/config-redis.yaml

help:
	go run ./cmd/linux/cli/main.go -help

# Для работы с Docker
.PHONY: redis-start redis-stop redis-status

redis-start:
	docker run -d --name redis-cache -p 6379:6379 redis:latest

redis-stop:
	docker stop redis-cache
	docker rm redis-cache

redis-status:
	docker ps | grep redis-cache