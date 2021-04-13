build:
	docker-compose build gateway

run:
	docker-compose up gateway

stop:
	docker-compose stop

test:
	go test -v  ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out