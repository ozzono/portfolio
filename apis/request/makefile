default:
	docker-compose up -d --build
test:
	go clean -testcache && docker-compose up -d --build && sleep 2 &&  go test ./... -v && docker-compose down
stop:
	docker-compose down