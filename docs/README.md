### Документация API и архитектуры (Swifty Gasprom — backend)

Ниже приведены детали архитектуры, переменных окружения, моделей, эндпоинтов и схем БД. Все маршруты сервера начинаются с префикса `/api/v1`.

## Архитектура

- Точка входа: `cmd/main.go`
  - Создаёт менеджер БД: `postgres.NewManager("postgres", os.Getenv("DB_CONNECT_STRING"))`
  - Применяет миграции: `manager.Migrate()`
  - Поднимает HTTP-сервер Gin: `r.Run(os.Getenv("SERVER_PORT"))`

- Роутинг: `internal/transport/http/router.go`
  - Регистрирует хендлеры доменов: `game`, `bank`, `auth`, `analytics`

- Middleware: `internal/handler/middleware/middleware.go`
  - `AuthMiddleware()` — проверяет заголовок `Authorization: Bearer <token>`, валидирует JWT, извлекает `userID` из клеймов и кладёт в контекст Gin (`c.Set("userID", ...)`).

- Аутентификация: `internal/handler/auth/*`, `internal/service/auth/tokens.go`, `internal/models/auth/*`
  - `JWT_SECRET` из ENV
  - Refresh-токен хранится в БД (таблица `session`)

- Игровая логика: `internal/handler/game/*`, `internal/service/game/*`, модели в `internal/models/game/*`

- Банк: `internal/handler/bank/*`, модели в `internal/models/bank/*`

- Миграции: `internal/repository/migrations/*.sql`

- Логирование: `internal/config/logger/log.go` (Logrus + Lumberjack)

- ENV loader: `internal/config/environment/env.go` (по желанию можно вызывать для загрузки `.env`).

## Переменные окружения

- `DB_CONNECT_STRING` — строка подключения к PostgreSQL, пример: `postgres://user:password@localhost:5432/dbname?sslmode=disable`
- `SERVER_PORT` — порт HTTP сервера, пример: `:8080`
- `JWT_SECRET` — секрет для подписи JWT (обязателен)
- `LOG_LEVEL` — `debug|info|warn|error` (по умолчанию `info`)
- `LOG_TO_CONSOLE` — `true|false` выводить логи в консоль помимо файла

`.env` можно загрузить через `environment.LoadEnvFile(".env")` (вызывается вручную в `main`, по умолчанию не подключён).

## Модели данных

### Аутентификация

```go
type BankUser struct {
    BankUserID string `json:"bank_user_id" binding:"required"`
}

type User struct {
    ID string `json:"id" binding:"required"`
    BankUserID string `json:"bank_user_id" binding:"required"`
    Token Tokens `json:"token" binding:"required"`
}

type RefreshToken struct {
    ID        string `json:"id"`
    Token     string `json:"token"`
    ExpiredAt time.Time `json:"expired_at"`
}

type Tokens struct {
    AccessToken  string `json:"access_token" binding:"required"`
    RefreshToken string `json:"refresh_token" binding:"omitempty"`
}
```

### Банк

```go
type BankProduct struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Url string `json:"url"`
    Type CardType `json:"type"`
}

type BankBonus struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Url string `json:"url"`
    Type BankBonusType `json:"type"`
}

type UpdateBonusStatusRequest struct {
    ID int `json:"id" binding:"required"`
    UserID string `json:"user_id" binding:"required"`
    Status BankBonusStatus `json:"status" binding:"required"`
}

// Типы
type CardType string
const (
    DebitCard CardType = "debit_card"
    CreditCard CardType = "credit_card"
)

type BankBonusType string
const (
    Cashback BankBonusType = "cashback"
    Discount BankBonusType = "discount"
)

type BankBonusStatus string
const (
    Claimed BankBonusStatus = "claimed"
    Pending BankBonusStatus = "pending"
    Expired BankBonusStatus = "expired"
)
```

### Игра

```go
type UserProgress struct {
    ID int64 `json:"id"`
    UserID string `json:"user_id"`
    Money float64 `json:"money"`
    Happiness int `json:"happiness"`
    Month int `json:"month"`
    Status UserStatus `json:"status"`
    Natural_expenses int `json:"natural_expenses"`
}

type UserGameInfo struct {
    UserProgress UserProgress
    News News
    Career Career
}

type News struct {
    ID int64 `json:"id"`
    Title string `json:"title"`
    Text string `json:"text"`
    Type NewsType `json:"type"`
}

type Career struct {
    ID int64 `json:"id"`
    Grade int `json:"grade"`
    Salary float64 `json:"salary"`
    Name string `json:"name"`
    CareerType CareerType `json:"career_type"`
}

type Event struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    EventText string `json:"event_text"`
    CapitalPercent int `json:"capital_percent"`
    ToAgree int `json:"to_agree"`
    Refuse int `json:"refuse"`
}

type LifeMarketItem struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    Cost float64 `json:"cost"`
    Category LifeMarketCategory `json:"category"`
    HappinessEffect int `json:"happiness_effect"`
}

type Vacancy struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    Type CareerType `json:"type"`
    CareerLevel int `json:"career_level"`
    CareerGrade int `json:"career_grade"`
    Salary float64 `json:"salary"`
    HappinessEffect int `json:"happiness_effect"`
    MonthsToGrade int `json:"months_to_grade"`
}

type Risk struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    Title string `json:"title"`
    RiskType RiskType `json:"risk_type"`
    CryptoCategory int `json:"crypto_category"`
    Price float64 `json:"price"`
    WinChance int `json:"win_chance"`
    LoseChance int `json:"lose_chance"`
}

type RiskItem struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    PriceHistory []struct {
        Price float64 `json:"price"`
        Month int `json:"month"`
    } `json:"price_history"`
    CurrentPrice float64 `json:"current_price"`
    WinChance int `json:"win_chance"`
    LoseChance int `json:"lose_chance"`
}

type RiskItems struct {
    Crypto []Risk `json:"crypto"`
    Stocks []Risk `json:"stocks"`
    Bets []Risk `json:"bets"`
    QuestionableProjects []Risk `json:"questionable_projects"`
}

type Month struct {
    Money float64 `json:"money"`
    Happiness int `json:"happiness"`
    Month int `json:"month"`
    News []News `json:"news"`
}

// Запросы
type LifeMarketItemRequest struct {
    ID        int64   `json:"id" binding:"required"`
    Cost      float64 `json:"cost" binding:"required"`
    Happiness int     `json:"happiness" binding:"required"`
}

type GetVacancyRequest struct {
    Name string `json:"name" binding:"required"`
}

type SetVacancyRequest struct {
    ID        int64   `json:"id" binding:"required"`
    Level     int     `json:"level" binding:"required"`
    Grade     int     `json:"grade" binding:"required"`
    MinSalary float64 `json:"min_salary" binding:"required"`
    MaxSalary float64 `json:"max_salary" binding:"required"`
}

type UpdatePlayerVacancyRequest struct {
    ID    int64      `json:"id" binding:"required"`
    Level int        `json:"level" binding:"required"`
    Grade int        `json:"grade" binding:"required"`
    Type  CareerType `json:"type" binding:"required"`
}

type Envelope struct {
    Type RiskType `json:"type" binding:"required"`
    Data []byte `json:"data" binding:"omitempty"`
}

type CryptoRequest struct {
    Price float64 `json:"price" binding:"required"`
    CryptoID int64 `json:"crypto_id" binding:"required"`
}

type StocksRequest struct {
    Price float64 `json:"price" binding:"required"`
    StockID int64 `json:"stock_id" binding:"required"`
}

type BetsRequest struct {
    Price float64 `json:"price" binding:"required"`
    BetID int64 `json:"bet_id" binding:"required"`
}

type QuestionableProjectsRequest struct {
    Price float64 `json:"price" binding:"required"`
    QPID int64 `json:"qp_id" binding:"required"`
}

// Типы
type UserStatus string
const (
    UserStatusActive UserStatus = "active"
    UserStatusFinished UserStatus = "finished"
    UserStatusBurnout UserStatus = "burnout"
    UserStatusBankrupt UserStatus = "bankrupt"
)

type LifeMarketCategory string
const (
    EntertainmentAndRecreation LifeMarketCategory = "entertainment_and_recreation"
    AppliancesAndGadgets LifeMarketCategory = "appliances_and_gadgets"
    GiftsAndSocialInteraction LifeMarketCategory = "gifts_and_social_interaction"
    EducationAndSelfDevelopment LifeMarketCategory = "education_and_self_development"
    HealthAndCare LifeMarketCategory = "health_and_care"
    EverydayJoys LifeMarketCategory = "everyday_joys"
)

type CareerType string
const (
    ITAndTechnology CareerType = "it_and_technology"
    EngineeringAndManufacturing CareerType = "engineering_and_manufacturing"
    MedicineAndHealthcare CareerType = "medicine_and_healthcare"
    MarketingAndSales CareerType = "marketing_and_sales"
    WorkingProfessions CareerType = "working_professions"
)

type RiskType string
const (
    Bets RiskType = "bets"
    QuestionableProjects RiskType = "questionable_projects"
    Crypto RiskType = "crypto"
    Stocks RiskType = "stocks"
)

type NewsType string
const (
    Economic NewsType = "economic"
    Political NewsType = "political"
    Corporate NewsType = "corporate"
    Useless NewsType = "useless"
)
```

## Эндпоинты

Префикс: `/api/v1`. Все защищённые эндпоинты требуют заголовок `Authorization: Bearer <ACCESS_TOKEN>`.

### Аутентификация

**POST /api/v1/login** (без авторизации)
- **Входные данные:**
  ```json
  {
    "bank_user_id": "string" // обязательное поле
  }
  ```
- **Успешный ответ (200):**
  ```json
  {
    "user": {
      "id": "string",
      "bank_user_id": "string", 
      "token": {
        "access_token": "string",
        "refresh_token": "string"
      }
    }
  }
  ```
- **Ошибки:** 400 (неверные данные), 500 (внутренняя ошибка)

**POST /api/v1/refresh** (требует Bearer)
- **Входные данные:**
  ```json
  {
    "access_token": "string", // опционально
    "refresh_token": "string" // опционально
  }
  ```
- **Успешный ответ (200):**
  ```json
  {
    "access_token": "string",
    "refresh_token": "string"
  }
  ```
- **Ошибки:** 400 (неверные данные), 401 (токен истёк), 500

**POST /api/v1/logout** (требует Bearer)
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "logout": "success"
  }
  ```
- **Ошибки:** 401 (токен истёк), 500

### Игровые эндпоинты (все требуют Bearer)

**GET /api/v1/start**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "game": {
      "id": 0,
      "user_id": "string",
      "money": 0.0,
      "happiness": 100,
      "month": 1,
      "status": "active",
      "natural_expenses": 0
    }
  }
  ```
- **Ошибки:** 500

**GET /api/v1/progress**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "game": {
      "UserProgress": {
        "id": 0,
        "user_id": "string",
        "money": 0.0,
        "happiness": 100,
        "month": 1,
        "status": "active",
        "natural_expenses": 0
      },
      "News": {
        "id": 0,
        "title": "string",
        "text": "string",
        "type": "economic"
      },
      "Career": {
        "id": 0,
        "grade": 0,
        "salary": 0.0,
        "name": "string",
        "career_type": "it_and_technology"
      }
    }
  }
  ```
- **Ошибки:** 404 (игра не найдена), 500

**POST /api/v1/life-market/:id**
- **Параметры URL:** `id` (int) - ID предмета
- **Входные данные:**
  ```json
  {
    "id": 0,
    "cost": 0.0,
    "happiness": 0
  }
  ```
- **Успешный ответ (200):**
  ```json
  {
    "status": "success"
  }
  ```
- **Ошибки:** 400 (недостаточно денег/неверные данные), 500

**GET /api/v1/life-market**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "items": [
      {
        "id": 0,
        "name": "string",
        "cost": 0.0,
        "category": "entertainment_and_recreation",
        "happiness_effect": 0
      }
    ]
  }
  ```
- **Ошибки:** 404 (предметы не найдены), 500

**GET /api/v1/life-market/:id**
- **Параметры URL:** `id` (int) - ID предмета
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "item": {
      "id": 0,
      "name": "string",
      "cost": 0.0,
      "category": "entertainment_and_recreation",
      "happiness_effect": 0
    }
  }
  ```
- **Ошибки:** 400 (неверный ID), 404 (предмет не найден), 500

**GET /api/v1/event**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "event": {
      "id": 0,
      "name": "string",
      "event_text": "string",
      "capital_percent": 0,
      "to_agree": 0,
      "refuse": 0
    }
  }
  ```
- **Ошибки:** 500

**GET /api/v1/vacancies**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "vacancies": [
      {
        "id": 0,
        "name": "string",
        "type": "it_and_technology",
        "career_level": 0,
        "career_grade": 0,
        "salary": 0.0,
        "happiness_effect": 0,
        "months_to_grade": 0
      }
    ]
  }
  ```
- **Ошибки:** 500

**GET /api/v1/vacancy**
- **Входные данные:**
  ```json
  {
    "name": "string"
  }
  ```
- **Успешный ответ (200):**
  ```json
  {
    "vacancy": {
      "id": 0,
      "name": "string",
      "type": "it_and_technology",
      "career_level": 0,
      "career_grade": 0,
      "salary": 0.0,
      "happiness_effect": 0,
      "months_to_grade": 0
    }
  }
  ```
- **Ошибки:** 400 (неверные данные), 404 (вакансия не найдена), 500

**GET /api/v1/player/vacancy**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "vacancy": {
      "id": 0,
      "name": "string",
      "type": "it_and_technology",
      "career_level": 0,
      "career_grade": 0,
      "salary": 0.0,
      "happiness_effect": 0,
      "months_to_grade": 0
    }
  }
  ```
- **Ошибки:** 404 (вакансия не найдена), 500

**POST /api/v1/player/vacancy**
- **Входные данные:**
  ```json
  {
    "id": 0,
    "level": 0,
    "grade": 0,
    "min_salary": 0.0,
    "max_salary": 0.0
  }
  ```
- **Успешный ответ (200):**
  ```json
  {
    "status": "success"
  }
  ```
- **Ошибки:** 400 (неверные данные), 500

**PUT /api/v1/player/vacancy/:name**
- **Параметры URL:** `name` (string) - имя вакансии
- **Входные данные:**
  ```json
  {
    "id": 0,
    "level": 0,
    "grade": 0,
    "type": "it_and_technology"
  }
  ```
- **Успешный ответ (200):**
  ```json
  {
    "vacancy": {
      "id": 0,
      "name": "string",
      "type": "it_and_technology",
      "career_level": 0,
      "career_grade": 0,
      "salary": 0.0,
      "happiness_effect": 0,
      "months_to_grade": 0
    }
  }
  ```
- **Ошибки:** 400 (неверные данные), 404 (вакансия не найдена), 500

**DELETE /api/v1/player/vacancy/:id**
- **Параметры URL:** `id` (int) - ID вакансии
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "status": "success"
  }
  ```
- **Ошибки:** 400 (неверный ID), 404 (вакансия не найдена), 500

**PUT /api/v1/player/vacancy/grade/:id**
- **Параметры URL:** `id` (int) - ID вакансии
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "newSalary": 0.0,
    "penaltyFactor": 0
  }
  ```
- **Ошибки:** 400 (неверный ID), 500

**POST /api/v1/risk**
- **Входные данные:**
  ```json
  {
    "type": "crypto", // crypto|stocks|bets|questionable_projects
    "data": "base64_encoded_json"
  }
  ```
  В зависимости от типа, `data` содержит:
  - **crypto:** `{"price": 0.0, "crypto_id": 0}`
  - **stocks:** `{"price": 0.0, "stock_id": 0}`
  - **bets:** `{"price": 0.0, "bet_id": 0}`
  - **questionable_projects:** `{"price": 0.0, "qp_id": 0}`
- **Успешный ответ (200):**
  ```json
  {
    "status": "success"
  }
  ```
- **Ошибки:** 400 (недостаточно денег/неверные данные), 404 (объект не найден), 500

**GET /api/v1/risk**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "risk": {
      "crypto": [
        {
          "id": 0,
          "name": "string",
          "title": "string",
          "risk_type": "crypto",
          "crypto_category": 0,
          "price": 0.0,
          "win_chance": 0,
          "lose_chance": 0
        }
      ],
      "stocks": [],
      "bets": [],
      "questionable_projects": []
    }
  }
  ```
- **Ошибки:** 404 (риски не найдены), 500

**GET /api/v1/risk/type/:type**
- **Параметры URL:** `type` (string) - тип риска
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "items": [
      {
        "id": 0,
        "name": "string",
        "title": "string",
        "risk_type": "crypto",
        "crypto_category": 0,
        "price": 0.0,
        "win_chance": 0,
        "lose_chance": 0
      }
    ]
  }
  ```
- **Ошибки:** 404 (риски не найдены), 500

**GET /api/v1/risk/:id/:type**
- **Параметры URL:** `id` (int), `type` (string)
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "item": {
      "id": 0,
      "name": "string",
      "price_history": [
        {
          "price": 0.0,
          "month": 0
        }
      ],
      "current_price": 0.0,
      "win_chance": 0,
      "lose_chance": 0
    }
  }
  ```
- **Ошибки:** 400 (неверный ID), 404 (объект не найден), 500

**GET /api/v1/player/risk**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "items": [
      {
        "id": 0,
        "name": "string",
        "price_history": [
          {
            "price": 0.0,
            "month": 0
          }
        ],
        "current_price": 0.0,
        "win_chance": 0,
        "lose_chance": 0
      }
    ]
  }
  ```
- **Ошибки:** 404 (риски не найдены), 500

**GET /api/v1/player/news**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "news": [
      {
        "id": 0,
        "title": "string",
        "text": "string",
        "type": "economic"
      }
    ]
  }
  ```
- **Ошибки:** 404 (новости не найдены), 500

**GET /api/v1/player/news/:id**
- **Параметры URL:** `id` (int) - ID новости
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "news": {
      "id": 0,
      "title": "string",
      "text": "string",
      "type": "economic"
    }
  }
  ```
- **Ошибки:** 400 (неверный ID), 404 (новость не найдена), 500

**POST /api/v1/month/next**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "newMonth": {
      "money": 0.0,
      "happiness": 100,
      "month": 2,
      "news": [
        {
          "id": 0,
          "title": "string",
          "text": "string",
          "type": "economic"
        }
      ]
    }
  }
  ```
- **Ошибки:** 500

### Банковские эндпоинты (все требуют Bearer)

**POST /api/v1/card/:id**
- **Параметры URL:** `id` (int) - ID банковского продукта
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "card": {
      "id": 0,
      "name": "string",
      "url": "string",
      "type": "debit_card"
    }
  }
  ```
- **Ошибки:** 400 (неверный ID), 404 (продукт не найден), 500

**GET /api/v1/card**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "card": {
      "id": 0,
      "name": "string",
      "url": "string",
      "type": "debit_card"
    }
  }
  ```
- **Ошибки:** 404 (продукт не найден), 500

**POST /api/v1/bonus/:id**
- **Параметры URL:** `id` (int) - ID бонуса
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "bonus": {
      "id": 0,
      "name": "string",
      "url": "string",
      "type": "cashback"
    }
  }
  ```
- **Ошибки:** 400 (неверный ID), 404 (бонус не найден), 500

**GET /api/v1/bonus**
- **Входные данные:** нет
- **Успешный ответ (200):**
  ```json
  {
    "bonuses": [
      {
        "id": 0,
        "name": "string",
        "url": "string",
        "type": "cashback"
      }
    ]
  }
  ```
- **Ошибки:** 404 (бонусы не найдены), 500

**PUT /api/v1/bonus/:id**
- **Параметры URL:** `id` (int) - ID бонуса
- **Входные данные:**
  ```json
  {
    "id": 0,
    "user_id": "string",
    "status": "claimed" // claimed|pending|expired
  }
  ```
- **Успешный ответ (200):**
  ```json
  {
    "status": "success"
  }
  ```
- **Ошибки:** 400 (неверные данные), 500

### Аналитика

**Пока эндпоинтов нет** 

## Схемы БД (основные миграции)

Аутентификация — `000001_create_auth_struct.up.sql`:
```sql
CREATE TABLE "user" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "bank_user_id" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "session" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" UUID NOT NULL,
  "token" text UNIQUE NOT NULL,
  "expires_at" timestamp NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
```

Банк — `000002_create_bank_struct.up.sql`:
```sql
CREATE TYPE "statuses" AS ENUM (
  'expired',
  'pending',
  'claimed'
);

CREATE TYPE "bank_product_type" AS ENUM (
  'credit_card',
  'debit_card'
);

CREATE TYPE "bonus_type" AS ENUM (
  'cashback',
  'discount'
);

CREATE TABLE "bank_products" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "url" text NOT NULL,
  "type" bank_product_type NOT NULL
);

CREATE TABLE "user_used_bank_products" (
  "started_at" timestamp NOT NULL DEFAULT (now()),
  "bank_product_id" integer NOT NULL,
  "user_id" UUID NOT NULL
);

CREATE TABLE "bank_bonuses" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "url" text NOT NULL,
  "type" bonus_type NOT NULL
);

CREATE TABLE "user_bonuses" (
  "claimed_at" timestamp NOT NULL DEFAULT (now()),
  "status" statuses NOT NULL,
  "user_id" UUID NOT NULL,
  "bonus_id" integer NOT NULL
);
```

Игра — `000003_create_game_struct.up.sql` (выдержки):
```sql
CREATE TABLE "user_progress" (
  "id" serial PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "money" decimal(10,2) DEFAULT 0,
  "happiness" integer DEFAULT 100,
  "health" integer DEFAULT 100,
  "month" integer DEFAULT 1,
  "status" user_statuses NOT NULL,
  "natural_expenses" integer,
  "last_sync_at" timestamp
);

CREATE TABLE "careers" (
  "id" serial PRIMARY KEY,
  "name" varchar(100),
  "career_type" career_category,
  "career_level" integer,
  "career_grade" integer,
  "min_salary" decimal(10,2),
  "max_salary" decimal(10,2),
  "happiness_penalty_factor" integer,
  "natural_expenses" decimal(10,2),
  "months_to_grade" integer
);

CREATE TABLE "life_market" (
  "id" serial PRIMARY KEY,
  "name" varchar(100),
  "cost" decimal(10,2),
  "category" life_market_category,
  "effect_duration" integer
);
```

Аналитика — `000004_create_analytics_struct.up.sql`:
```sql
CREATE TABLE "game_logs" (
  "id" serial PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "event_id" integer,
  "action" jsonb
);

CREATE TABLE "analytics_snapshots" (
  "id" serial PRIMARY KEY,
  "metric" jsonb NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "game_logs" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
```

## Примечания по безопасности

- Все защищённые эндпоинты требуют корректного JWT Access-токена в заголовке `Authorization`.
- Секрет `JWT_SECRET` храните вне репозитория (в `.env` или секретах CI/CD).

## Запуск миграций вручную

Сервер выполнит `Up()` на старте. Для ручного управления можно использовать `golang-migrate` CLI с путём `internal/repository/migrations`.

## Диагностика и логи

См. `internal/config/logger/log.go`. Уровни логов и вывод регулируются ENV.
