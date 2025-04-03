#!/bin/sh

# マイグレーション実行
echo "Running database migrations..."
goose -dir ./db/migrate postgres "host=${DB_HOST} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable" up

# アプリケーション起動
echo "Starting application..."
exec ./omikuji-app
