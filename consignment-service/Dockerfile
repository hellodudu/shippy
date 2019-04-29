FROM alpine:3.8

RUN mkdir /app

WORKDIR /app

ADD consignment-service /app/consignment-service

# 运行服务端
ENTRYPOINT ["./consignment-service"]
