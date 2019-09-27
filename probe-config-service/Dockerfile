FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o probe-config-service


FROM debian:latest

RUN mkdir /app
WORKDIR /app

ADD config.csv /app/config.csv
COPY --from=builder /app/probe-config-service .

CMD ["./probe-config-service"]
