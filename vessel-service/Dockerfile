FROM alpine:3.8

RUN mkdir /app

WORKDIR /app

COPY shippy-service-vessel /app/shippy-service-vessel 
COPY vessels.json /app/vessels.json

ENTRYPOINT ["./shippy-service-vessel"]
