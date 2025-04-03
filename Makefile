# 環境初期設定
dev-init:
	@echo "開発環境の初期設定を開始します"
	# .envファイルのコピー
	cp backend/.env.example backend/.env.local
	cp frontend/.env.example frontend/.env.local
	# 依存関係のインストール
	cd backend && go mod download
	cd frontend && npm install
	@echo "開発環境の初期設定が完了しました"

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
