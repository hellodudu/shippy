.PHONY: proto
proto:
	protoc -I=../proto --go_out=plugins=micro:../proto/ ../proto/consignment/consignment.proto

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -o shippy-service-consignment
	docker build -t shippy-service-consignment .

.PHONY: run
run:
	docker run -p 50051:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \
	 -e MICRO_REGISTRY=mdns \
	 shippy-service-consignment
