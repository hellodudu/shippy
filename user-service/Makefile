.PHONY: proto
proto:
	protoc -I=../proto --go_out=plugins=micro:../proto ../proto/user/user.proto

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -o shippy-service-user

.PHONY: run
run:
	docker run -p 50053:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \
	 -e MICRO_REGISTRY=mdns \
	 -e GO_MICRO_BROKER \
	 -e DB_HOST=root:@(localhost:3306)/db_shippy
	 shippy-service-user
