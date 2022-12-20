with-docker:
	docker-compose up -d

without-docker:
	docker-compose up -d mongodb;
	cmd/url-shortened >> url-shortened.log 2>&1  &

tidy:
	go mod tidy;
	go mod download;

mongo-logs:
	docker logs -f mongodb

stop-docker-api:
	docker-compose down

stop-api:
	docker-compose down
	pkill url-shortened

api-docker-logs:
	docker logs -f golang

api-logs:
	tail -f url-shortened.log

rebuild:
	rm cmd/url-shortened
	go build -o cmd/url-shortened ./cmd/main.go

rebuild-tests:
	rm -rf handler.test
	rm -rf database.test
	go test ./internal/handler -c
	go test ./internal/database -c

tests:
	./handler.test -test.v
	./database.test -test.v
