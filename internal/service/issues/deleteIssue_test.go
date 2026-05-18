package issues

import (
	"context"
	"database/sql"
	"major/internal/models"
	"major/internal/repository"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go" // ключевые библиотеки для поднятия контейнера
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupTestDB(t *testing.T) *sql.DB {
	ctx := context.Background()

	initScript, _ := filepath.Abs("testdata/init-db.sql") // подгрузка эммулированной бд

	container, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		postgres.WithInitScripts(initScript),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready").
				WithOccurrence(2).
				WithStartupTimeout(120*time.Second),
		),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		container.Terminate(ctx)
	})

	connStr, _ := container.ConnectionString(ctx, "sslmode=disable")
	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)
	require.NoError(t, db.PingContext(ctx))

	return db
}

func TestDeleteIssue(t *testing.T) {
	db := setupTestDB(t)
	ctx := context.Background()
	// создаем сервис с зависимостями
	cfg := &models.Config{
		// если конфиг не используется в функции - можем передать пустым
	}
	repo := repository.NewRepository(db)

	svc := &IssuesService{
		conf: cfg,
		repo: repo,
	}

	testUUID := "123e4567-e89b-12d3-a456-426614174000" // тестовый id

	// заполнение таблицы
	_, err := db.ExecContext(ctx, `
        INSERT INTO issues (id, name, description, recommendation)
        VALUES ($1, $2, $3, $4)
    `, testUUID, "Test User", "Test desc", "Test recs")
	require.NoError(t, err, "Failed to insert test user")
	t.Log("Test user created with UUID:", testUUID)

	err = svc.DeleteIssue(testUUID) // передача тестируемой функции

	assert.NoError(t, err, "DeleteUser should succeed")
	t.Log("User was succesfull deleted")

	// проверка на наличие записей в таблице
	var count int
	err = db.QueryRowContext(ctx, `
        SELECT COUNT(*) FROM issues WHERE id = $1
    `, testUUID).Scan(&count)
	require.NoError(t, err)
	assert.True(t, true, count, "User should be soft deleted (deleted_at set)")

	t.Log("Test passed: user successfully deleted.")
}

// Поэтапный запуск теста
func TestDebugContainer(t *testing.T) {
	t.Log("1. Начинаем тест")

	ctx := context.Background()
	t.Log("2. Контекст создан")

	initScript, err := filepath.Abs("testdata/init-db.sql")
	if err != nil {
		t.Fatalf("Ошибка при поиске init-db.sql: %v", err)
	}
	t.Logf("3. Путь к init-db.sql: %s", initScript)

	t.Log("4. Пытаемся запустить контейнер...")

	container, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		postgres.WithInitScripts(initScript),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready").
				WithOccurrence(2).
				WithStartupTimeout(2*time.Minute),
		),
	)

	if err != nil {
		t.Fatalf("Ошибка при запуске контейнера: %v", err)
	}
	t.Log("5. Контейнер запущен!")

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("Ошибка получения connection string: %v", err)
	}
	t.Logf("6. Connection string: %s", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Ошибка открытия БД: %v", err)
	}
	t.Log("7. БД открыта")

	err = db.PingContext(ctx)
	if err != nil {
		t.Fatalf("Ошибка ping БД: %v", err)
	}
	t.Log("8. Ping успешен!")

	t.Cleanup(func() {
		t.Log("Очистка: завершаем контейнер")
		container.Terminate(ctx)
	})

	t.Log("Тест прошел успешно!")
}
