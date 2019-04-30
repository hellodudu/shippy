.PHONY: proto
proto:
	protoc -I=../proto --go_out=plugins=micro:../proto ../proto/vessel.proto

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -o shippy-service-vessel
	docker build -t shippy-service-vessel .

.PHONY: run
run:
	docker run -p 50052:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \
	 -e MICRO_REGISTRY=mdns \
	 shippy-service-vessel
