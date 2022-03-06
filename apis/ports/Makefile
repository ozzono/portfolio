# ==============================================================================
# Go migrate postgresql ports service

migrate-version:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ports_service?sslmode=disable" version

migrate-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ports_service?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ports_service?sslmode=disable" -verbose down


# Download go packages
tidy:
	go mod tidy
	go mod download

run:
	go run cmd/main.go

db:
	docker-compose up -d

