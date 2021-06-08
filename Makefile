build:
	make gp
	docker-compose build gateway
run:
	docker-compose up gateway
ps:
	docker ps
stop:
	docker-compose stop
test:
	go test -v  ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
gp:
	protoc --go_out=./models api/proto/v9.proto
	protoc --go_out=./models api/proto/v10.proto
	protoc --go_out=./models api/proto/slp_net_protocol.proto