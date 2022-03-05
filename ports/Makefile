# ==============================================================================
# Go migrate postgresql ports service

version_ports_db:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ports_service?sslmode=disable" version

migrate_up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ports_service?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ports_service?sslmode=disable" -verbose down


# Download go packages
tidy:
	go mod tidy
	go mod download

run:
	go run cmd/main.go