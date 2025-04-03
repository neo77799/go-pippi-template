# ビルドステージ
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# gooseをインストール
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o omikuji-app ./cmd

# 実行ステージ
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/omikuji-app .
COPY --from=builder /app/db/migrate ./db/migrate
COPY --from=builder /go/bin/goose /usr/local/bin/goose

RUN apk add --no-cache tzdata

EXPOSE 8080

COPY ./scripts/entrypoint.sh .
RUN chmod +x ./entrypoint.sh

CMD ["./entrypoint.sh"]
