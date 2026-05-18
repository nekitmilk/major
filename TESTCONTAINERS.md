# Интеграционное Тестирование методом виртуального контейнера
Она позволяет поднимать как одиночные контейнеры для ваших задач или весь проект целиком.

## Цель

Упростить этап интеграционного тестирования сервисов/функций для корректной работы которых требуется поднятый контейнер той или иной утилиты (например, База Данных).

## Инструмент интеграционного тестирования методом контейнера 

В проекте используется:
- **github.com/testcontainers/testcontainers-go** — основной пакет для языка Go, необходимый для первичного запуска виртуального контейнера.
- **github.com/testcontainers/testcontainers-go/modules/postgres** — подпакет, драйвер для корректной работы контейнера с утилитой PostgeSQL. (имя пакета может отличаться в зависимости от требуемой Базы Данных)
- **github.com/testcontainers/testcontainers-go/wait** — подпакет, необходим для создания тайминга между запуском самого контейнера и его сервиса. Например, чтобы тест не начинал работу, пока база данных (или другой сервис) внутри контейнера действительно не запустится и не будет готов к приему запросов - не будет начат запуск.


**Установка:**
```bash
go get github.com/testcontainers/testcontainers-go
go get github.com/testcontainers/testcontainers-go/modules/postgres
```

**Запуск:**
```bash
go test -v -timeout 5m -run TestDeleteIssue
```

## Компоненты системы

Интеграционный тест объединяет следующие компоненты системы:
- PostgreSQL Container (Image: docker.io/postgres:16-alpine) - База данных для тестов
- Container Runtime - Запуск и управление контейнерами



## Конфигурация

Параметры настройки на примере зафиксированы в файле 
`internal/service/issues/deleteIssue_test.go` и включают:

- postgres.WithDatabase("testdb") - имя ДБ
- postgres.WithUsername("testuser") - Пользователь
- postgres.WithPassword("testpass") - Пароль
- postgres.WithInitScripts(initScript) - init-скрипт (путь к миграции БД)
- testcontainers.WithWaitStrategy(...) - Стратегия ожидания (первичные настройки системы)

- **Содержимое проекта**
- **Реализация**: тестирование сервисной функции 'DeleteIssue' `service/issues/DeleteIssue_test.go`
- **Миграция**: специальная миграция для проверки функции `service/issues/testdata/init-db.sql`


## Отчеты по результатам

Формат вывода — консольный (текстовый) с указанием:
- Этапа проверки функции
- Выявления ошибки при работе
- Краткого описания проблемы

**Пример отчета:**
```text
deleteIssue_test.go:73: Test user created with UUID: 123e4567-e89b-12d3-a456-426614174000
deleteIssue_test.go:78: User was succesfull deleted
deleteIssue_test.go:88: Test passed: user successfully deleted.
--- PASS: TestDeleteIssue (5.24s)
PASS
ok      major/internal/service/issues   5.643s
```
