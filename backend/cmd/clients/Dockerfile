FROM golang:alpine AS builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest && go build -o main ./cmd/clients/main.go

FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache postgresql-client

COPY --from=builder /app/config ./config
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose

COPY ./cmd/clients/entrypoint.sh .
RUN chmod +x ./entrypoint.sh ./main

ENV CONFIG_PATH=config/clients/production.yaml
ENTRYPOINT ["./entrypoint.sh"]
