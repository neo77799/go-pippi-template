# Docker操作
docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

run-test:
	cd backend && go test -v ./...

lint:
	cd backend && golangci-lint run

run-all:
	docker-compose down && docker-compose up --build

# 環境別実行コマンド
dev:
	docker-compose --env-file backend/.env.local up -d

prod:
	docker-compose --env-file backend/.env.prod up -d

test:
	docker-compose --env-file backend/.env.test up -d

# 開発用ショートカット
dev-shortcut: docker-build docker-up
