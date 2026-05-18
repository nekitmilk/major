BINARY_NAME=major

run:
	go run cmd/main.go

cli-run:
	CONSOLE_MODE=true go run cmd/main.go $(FILEPATH)

build:
	go build -o $(BINARY_NAME).out ./cmd/main.go

docker-up-local:
	sudo docker compose -f deployments/local/docker-compose.yml up -d

docker-build-up-local:
	sudo docker compose -f deployments/local/docker-compose.yml up -d --build

docker-down-local:
	sudo docker compose -f deployments/local/docker-compose.yml down

migration-create:
	migrate create -ext sql -dir migrations/ -seq $(NAME_MIGRATE)

migration-db-up:
	migrate -path migrations/ -database "postgresql://admin:admin@localhost:6432/major?sslmode=disable" -verbose up

migration-db-down:
	migrate -path migrations/ -database "postgresql://admin:admin@localhost:6432/major?sslmode=disable" -verbose down

gen-doc:
	swag init -d ./cmd,internal -o ./docs

.PHONY: lint
lint:
	@echo "Running golangci-lint..."
	golangci-lint run ./cmd/... ./internal/... ./pkg/...

.PHONY: lint-fix
lint-fix:
	@echo "Running golangci-lint with --fix..."
	golangci-lint run --fix ./cmd/... ./internal/... ./pkg/...

fuzz-all:
	@echo "Running all fuzz tests..."
	@go list ./... | grep -v vendor | while read pkg; do \
		if go test -list=Fuzz $$pkg 2>/dev/null | grep -q "^Fuzz"; then \
			echo "=== Fuzzing $$pkg ==="; \
			go test -fuzz=. -fuzztime=30s $$pkg; \
			echo ""; \
#		else \
#			echo "=== Skipping $$pkg (no fuzz tests) ==="; \
		fi \
	done

fuzz-parser:
	go test -fuzz=. -fuzztime=30s ./pkg/parser

fuzz-issues:
	go test -fuzz=. -fuzztime=30s ./internal/service/issues