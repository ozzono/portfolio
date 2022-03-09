# ==============================================================================
# Go migrate postgresql ports service

migrate-version:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ports_service?sslmode=disable" version

migrate-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ports_service?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ports_service?sslmode=disable" -verbose down

migrate-drop:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/ports_service?sslmode=disable" -verbose drop -f

migrate-fresh:
	make migrate-drop && make migrate-up

# Download go packages
tidy:
	go mod tidy
	go mod download

run:
	docker-compose up -d --build --remove-orphans && docker exec ports_golang sh /ports/migrate.sh

stop:
	docker-compose down
