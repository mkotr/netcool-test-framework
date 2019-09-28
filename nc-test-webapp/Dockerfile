FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o nc-test-webapp


FROM debian:latest

RUN mkdir /app
WORKDIR /app

ADD views /app/views 
COPY --from=builder /app/nc-test-webapp .

CMD ["./nc-test-webapp"]
