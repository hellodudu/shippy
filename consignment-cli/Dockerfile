FROM alpine:3.8

RUN mkdir -p /app

WORKDIR /app

COPY consignment.json /app/consignment.json
COPY shippy-cli-consignment /app/shippy-cli-consignment

ENTRYPOINT ["./shippy-cli-consignment"]
