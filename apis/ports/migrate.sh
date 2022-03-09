#!/bin/bash
migrate -path migrations -database "postgresql://postgres:postgres@ports_postgresql:5432/ports_service?sslmode=disable" -verbose drop -f
migrate -path migrations -database "postgresql://postgres:postgres@ports_postgresql:5432/ports_service?sslmode=disable" -verbose up
