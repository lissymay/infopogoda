.PHONY: run help build clean

run:
	go run ./cmd/linux/cli/main.go

run-config:
	go run ./cmd/linux/cli/main.go -config $(CONFIG)

help:
	go run ./cmd/linux/cli/main.go -help

build:
	go build -o bin/weather.exe ./cmd/linux/cli/main.go

clean:
	if exist bin rmdir /s /q bin