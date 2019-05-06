.PHONY: proto
proto:
	make -C consignment-service proto
	make -C vessel-service proto
	make -C user-service proto

.PHONY: build
build:
	make -C consignment-service build
	make -C vessel-service build
	make -C user-service build
	make -C consignment-cli build

.PHONY: compose_build
compose_build:
	docker-compose build consignment-service
	docker-compose build vessel-service
	docker-compose build user-service
