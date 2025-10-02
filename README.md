# Swifty Gasprom — Backend + Frontend

Проект — прототип игрового симулятора «Финансовый путь жизни», разработанный для хакатона.  
Состоит из **backend** на Go (Gin + PostgreSQL + JWT) и **frontend** на Flutter (mobile + web).

---

## 🚀 Технологии

### Backend
- Go
- Gin (HTTP API)
- PostgreSQL
- JWT (аутентификация)
- golang-migrate (миграции БД)
- logrus + lumberjack (логирование)

### Frontend
- Flutter (mobile + web)
- Web build (`flutter build web` → `build/web`)

---

## 📂 Структура репозитория
- `backend/cmd/` — точка входа сервера
- `backend/internal/handler` — контроллеры (auth, game, bank и др.)
- `backend/internal/service` — бизнес-логика
- `backend/internal/repository/postgres` — работа с БД + миграции
- `frontend/gazprom_web` — Flutter web
- `frontend/gazprom_mobile` — Flutter mobile

---

## ⚙️ Настройка Backend

1. Склонировать репозиторий:
```bash
git clone https://github.com/Hirogava/swifty-gasprom.git
cd swifty-gasprom
```

2. Создать файл `.env` в корне backend на основе `.env.example`.

3. Запустить сервер:
```bash
cd backend
go run ./cmd
```

Бэкенд будет доступен на `http://localhost:8080/api/v1`.

---

## 📱 Запуск Frontend

### Web
```bash
cd frontend/gazprom_web
flutter pub get
flutter build web
# результат в build/web
```

### Mobile
```bash
cd frontend/gazprom_mobile
flutter pub get
flutter run
```

---

## 🔑 Основные эндпоинты API
- `POST /api/v1/login` — вход (JWT access + refresh)
- `POST /api/v1/refresh` — обновление токена
- `POST /api/v1/logout` — выход
- `GET /api/v1/game/...` — игровые механики
- `GET /api/v1/bank/...` — банковские данные

---

## 🐳 Docker
Docker-конфигурация будет в отдельной ветке (с `Dockerfile` и `docker-compose.yml`).


