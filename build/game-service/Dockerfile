FROM docker:19.03.8 AS builder
FROM golang:1.16-alpine

WORKDIR app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY . ./

RUN go mod download

RUN go build -o /game-service ./cmd/game-service

CMD ["/game-service"]