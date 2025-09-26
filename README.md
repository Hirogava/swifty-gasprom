### Swifty Gasprom — backend

Бэкенд на Go с REST API для симулятора финансового/карьерного прогресса игрока. Проект использует `Gin` как HTTP-фреймворк и PostgreSQL с миграциями через `golang-migrate`.

- Язык/версия Go: указан в `go.mod` — 1.24.0
- Фреймворк: `github.com/gin-gonic/gin`
- База данных: PostgreSQL
- Миграции: `github.com/golang-migrate/migrate/v4`
- Аутентификация: JWT (`github.com/golang-jwt/jwt/v5`), refresh-токены в БД
- Логирование: `logrus` + ротация логов `lumberjack`

## Быстрый старт

1) Установите переменные окружения (см. ниже раздел ENV). Пример для Windows PowerShell:
```powershell
$env:DB_CONNECT_STRING = "postgres://user:password@localhost:5432/dbname?sslmode=disable"
$env:SERVER_PORT = ":8080"
$env:JWT_SECRET = "supersecret"
$env:LOG_LEVEL = "info" # debug|info|warn|error
$env:LOG_TO_CONSOLE = "true"
```

2) Запустите сервер:
```powershell
go run ./cmd
```

3) Сервер поднимется на порту из `SERVER_PORT`. Все маршруты начинаются с `/api/v1`.

Миграции БД выполняются автоматически при старте (см. `internal/repository/postgres/migrator.go`).

## Структура проекта (высокоуровнево)

- `cmd/main.go` — точка входа: создаёт менеджер БД, применяет миграции, поднимает HTTP-сервер
- `internal/transport/http/router.go` — конфигурация роутера, регистрация хендлеров
- `internal/handler/**` — обработчики доменов: `auth`, `game`, `bank`, `analytics`, общие `middleware`
- `internal/repository/postgres/**` — менеджер подключения, миграции, доступ к данным
- `internal/models/**` — модели запросов/ответов и доменных типов
- `internal/service/**` — прикладная логика (игровая, токены)
- `internal/config/**` — загрузка `.env`, инициализация логгера

Подробности по API, моделям и схемам БД см. в `docs/README.md`.

## Переменные окружения (ENV)

- `DB_CONNECT_STRING` — строка подключения к PostgreSQL, например: `postgres://user:pass@host:5432/db?sslmode=disable`
- `SERVER_PORT` — порт для HTTP-сервера, например: `:8080`
- `JWT_SECRET` — секрет для подписи JWT
- `LOG_LEVEL` — уровень логирования: `debug|info|warn|error` (по умолчанию `info`)
- `LOG_TO_CONSOLE` — если `true`, дублирует логи в консоль помимо файла

Загрузка из `.env` поддерживается утилитой в `internal/config/environment/env.go` (файл `.env` в `.gitignore`). При необходимости вызовите загрузку в `main` (сейчас не вызывается).

## Сборка и запуск

- Запуск в dev-режиме: `go run ./cmd`
- Сборка бинаря: `go build -o bin/server ./cmd`
- Переменные окружения обязательны для успешного старта (подключение к БД и JWT).

## Краткий обзор API

- `POST /api/v1/login` — вход по `bank_user_id`, выдаёт `access_token` и `refresh_token`
- `POST /api/v1/refresh` — обновить access-токен (требует Bearer-токен)
- `POST /api/v1/logout` — завершить сессию (требует Bearer-токен)
- Раздел Game (все под Bearer): старт игры, прогресс, вакансии, рынок жизни, риск-сценарии и т.д.
- Раздел Bank (все под Bearer): работа с банковскими продуктами и бонусами

Полная спецификация: `docs/README.md`.

## Логирование

`internal/config/logger/log.go` настраивает `logrus` в JSON-формате и ротацию файлов. Управление через `LOG_LEVEL` и `LOG_TO_CONSOLE`.

## Миграции

SQL-миграции лежат в `internal/repository/migrations`. При старте `cmd/main.go` вызывает `manager.Migrate()`.