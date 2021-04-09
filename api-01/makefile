default:
	docker-compose up -d
test:
	docker-compose up -d --build && go test ./... -v
stop:
	docker-compose down