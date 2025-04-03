# Go + Next.js フルスタック開発テンプレート

このリポジトリは個人開発用のテンプレートリポジトリで、Go(Gin)バックエンドとNext.js(App Router)フロントエンドのフルスタック構成となっています。

## 技術スタック

### フロントエンド
- Next.js (App Router)
- TypeScript
- React Query (TanStack Query)
- NextAuth.js (Auth.js)
- Tailwind CSS
- shadcn/ui

### バックエンド
- Go (Gin)
- GORM (ORM)
- PostgreSQL
- Goose (DBマイグレーション)

### 開発ツール
- Docker / docker-compose
- Makefile
- ESLint / Prettier
- golangci-lint (Go用Linter)

## 開発環境構築

1. 必要なツールのインストール:
   - Docker
   - Go (1.20以上)
   - Node.js (18以上)

2. 環境変数の設定と依存関係のインストール:
   ```bash
   make dev-init
   ```
   このコマンドで以下が実行されます:
   - backend/frontendの.env.exampleを.env.localにコピー
   - Goモジュールの依存関係をインストール
   - npmパッケージをインストール

3. Docker環境の構築:
   ```bash
   make docker-build
   make docker-up
   ```

## アプリケーションの実行

```bash
# 全サービスビルド＆起動
make dev

# 個別コマンド
make docker-build  # コンテナビルド
make docker-up    # コンテナ起動 
make docker-down  # コンテナ停止
make docker-test  # テスト実行
```

## APIドキュメント

### おみくじAPI

`GET /api/omikuji`

レスポンス例:
```json
{
  "result": "大吉",
}
```

## テストの実行

```bash
# バックエンドテスト
make docker-test

# バックエンドLint実行
make lint
```
