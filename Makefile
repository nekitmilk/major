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