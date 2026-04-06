# Major - Security Configuration Scanner

Утилита для анализа конфигурационных файлов веб-приложений и выявления потенциально опасных настроек.

## 📁 Структура проекта

```
.
├── cmd/                          # Точка входа в приложение
│   └── main.go                   # Главный файл (CLI + HTTP сервер)
├── internal/                     # Внутренняя логика
│   ├── cli/                      # CLI интерфейс
│   │   └── check/                # Обработчик CLI команд
│   ├── config/                   # Загрузка конфигурации приложения
│   ├── handler/                  # HTTP хендлеры (REST API)
│   │   ├── check/                # Проверка конфигов
│   │   ├── health/               # Health check
│   │   ├── issues/               # CRUD угроз
│   │   ├── rules/                # CRUD правил
│   │   ├── severities/           # CRUD уровней опасности
│   │   └── middleware.go         # CORS, Auth, Logger
│   ├── models/                   # DTO и модели данных
│   ├── repository/               # Слой доступа к БД (PostgreSQL)
│   │   ├── issues/               # Работа с угрозами
│   │   ├── rules/                # Работа с правилами
│   │   ├── severities/           # Работа с уровнями
│   │   └── sql/                  # Инициализация БД
│   └── service/                  # Бизнес-логика
│       ├── check/                # Проверка конфигов через CEL
│       ├── issues/               # Управление угрозами
│       ├── rules/                # Управление правилами
│       └── severities/           # Управление уровнями
├── pkg/                          # Переиспользуемые пакеты
│   ├── celchecker/               # Обёртка над Google CEL
│   ├── customErrors/             # Кастомные ошибки (404, 409, 500...)
│   ├── logger/                   # Логирование
│   ├── parser/                   # Парсер JSON/YAML конфигов
│   └── server/                   # HTTP сервер
├── migrations/                   # Миграции БД
│   ├── 000001_init_db.up.sql    # Создание таблиц
│   └── 000001_init_db.down.sql  # Откат
├── deployments/                  # Docker композ для локальной разработки
│   └── local/
│       ├── docker-compose.yml    # PostgreSQL + приложение
│       └── default_major_data.sql # Тестовые данные
├── testConfigs/                  # Примеры конфигов для тестирования
│   ├── config_bad_1.json        # Конфиг с проблемами (debug + пароль + MD5)
│   ├── config_bad_2.yml         # Конфиг с проблемами (0.0.0.0 + TLS)
│   └── config_good.json         # Безопасный конфиг
├── docs/                         # Swagger документация
├── Makefile                      # Автоматизация команд
└── go.mod                        # Зависимости
```

## Быстрый старт

### Требования
- Go 1.25.5
- Docker & Docker Compose
- Make

### 1. Запуск через make

```bash
make docker-build-up-local
```

При запущенной БД можно также:

```bash
make build & make run # запуск с установленным конфигом (http сервер по умолчанию)
make build & make cli-run FILEPATH=testConfigs/config_bad_1.json  # Принудительный запуск в консольном режиме   
```

### 2. Накатить миграции на БД 

```bash
make migration-db-up
./deployments/local/db_fill.sh # Заполнить БД тестовыми данными 
```

### 4. Запуск приложения

#### Режим CLI (консольная утилита)

```bash
# Проверка файла
CONSOLE_MODE=true go run cmd/main.go testConfigs/config_bad_1.json

# Без выхода с ошибкой
CONSOLE_MODE=true go run cmd/main.go -s testConfigs/config_bad_1.json

# Чтение из STDIN
CONSOLE_MODE=true cat testConfigs/config_bad_1.json | go run cmd/main.go --stdin
```

#### Режим HTTP сервера

console_mode также переключает режимы запуска
Установи в `configs/.configs.json`:
```json
{
  "debug": true,
  "console_mode": false, 
  ...
}
```

Затем:
```bash
go run cmd/main.go
```

Сервер запустится на порту, указанном в конфиге (по умолчанию `:8181`).

Swagger документация: `http://localhost:8181/swagger/index.html`

### Makefile команды

```bash
make run                      # Запуск приложения
make run-cli                  # Запуск приложения принудительно в консольной форме
make build                    # Сборка бинарника
make docker-up-local          # Запуск PostgreSQL в Docker
make docker-down-local        # Остановка PostgreSQL
make migration-create NAME=name  # Создать новую миграцию
make migration-db-up          # Накатить миграции
make migration-db-down        # Откатить миграции
make gen-doc                  # Сгенерировать Swagger документацию
```

## Как это работает

### Архитектура

Проект построен на **чистой архитектуре** с разделением на слои:

1. **Handler** (HTTP/CLI) → принимает запросы
2. **Service** → бизнес-логика
3. **Repository** → работа с БД
4. **Models** → DTO и сущности

### Модель данных

```
Severity (уровень опасности)
├── id (UUID)
├── name (LOW/MEDIUM/HIGH)
├── level (1/2/3)
└── description

Issue (угроза/проблема)
├── id (UUID)
├── name
├── description
├── recommendation
└── severity_id → Severity.id

Rule (правило выявления)
├── id (UUID)
├── name
├── description
├── issue_id → Issue.id
├── expression (CEL выражение)
└── enabled (bool)
```

### Логика проверки

1. Пользователь отправляет **JSON/YAML конфиг** (через CLI или REST API)
2. Конфиг парсится в `map[string]interface{}`
3. Загружаются все **включенные правила** из БД
4. Каждое правило содержит **CEL-выражение**, например:
    - `config.log.level == "debug" || config.debug == true`
    - `has(config.password) && !config.password.startsWith("$")`
    - `config.host == "0.0.0.0" || config.bind == "0.0.0.0"`
5. CEL движок ([Google CEL](https://github.com/google/cel-go)) вычисляет выражение на конфиге
6. Если выражение вернуло `true` → правило сработало
7. Возвращается список **угроз** (Issue) с описанием и рекомендацией

### Пример правила в БД

```sql
INSERT INTO rules (id, name, description, issue_id, expression) VALUES (
    'f1a2b3c4...',
    'Debug mode check',
    'Проверяет, включён ли debug-режим',
    (SELECT id FROM issues WHERE name = 'Debug mode enabled'),
    'config.log.level == "debug" || config.debug == true'
);
```

### Добавление нового правила

1. Создать **Issue** (угрозу) через API `POST /api/v1/issues`
2. Создать **Rule** (правило) через API `POST /api/v1/rules` с CEL-выражением
3. Правило начнёт работать автоматически (при `enabled: true`)

## 📊 API Endpoints

| Метод | Путь | Описание |
|-------|------|----------|
| POST | `/api/v1/check` | Проверить конфиг на安全问题 |
| POST | `/api/v1/severities` | Создать уровень опасности |
| GET | `/api/v1/severities` | Получить все уровни |
| DELETE | `/api/v1/severities/:id` | Удалить уровень |
| POST | `/api/v1/issues` | Создать угрозу |
| GET | `/api/v1/issues` | Получить все угрозы |
| DELETE | `/api/v1/issues/:id` | Удалить угрозу |
| POST | `/api/v1/rules` | Создать правило |
| GET | `/api/v1/rules` | Получить все правила |
| GET | `/api/v1/rules/full` | Получить правила с деталями |
| GET | `/api/v1/rules/full/:id` | Получить правило по ID |
| DELETE | `/api/v1/rules/:id` | Удалить правило |

### Пример запроса к `/api/v1/check`

```bash
curl -X POST http://localhost:8181/api/v1/check \
  -H "Content-Type: application/json" \
  -d '{
    "config": "{\"debug\":true,\"database\":{\"password\":\"secret123\"},\"storage\":{\"digest_algorithm\":\"MD5\"}}"
  }'
```

## 🧪 Тестовые конфиги

- `testConfigs/config_bad_1.json` → 3 проблемы (LOW + 2xHIGH)
- `testConfigs/config_bad_2.yml` → 2 проблемы (MEDIUM + HIGH)
- `testConfigs/config_good.json` → 0 проблем

## Дополнительно

- **CEL выражения**: поддерживают `has()`, `in`, `matches()`, `startsWith()`
- **Обработка ошибок**: кастомные ошибки с маппингом в HTTP статусы (404, 409, 500)
- **Логирование**: структурированное через `logrus` с request_id
- **Документация**: Swagger автоматически генерируется через `swag init`

## Сборка бинарника

```bash
make build
./major.out testConfigs/config_bad_1.json
```

## Полный запуск в Docker

```bash
make docker-build-up-local   # Собрать и запустить всё
make migration-db-up         # Накатить миграции
# Заполнить данными через db_fill.sh
```

