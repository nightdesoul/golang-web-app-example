# Golang simple web app example

Простое Golang web приложение. Используется для обучения практикам DevOps.

## Запуск Unit тестов

```bash
go test
```

## Сборка под Linux

```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-extldflags '-static'" -o service
```