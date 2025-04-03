# バックエンドプロジェクト

Go (Gin) + GORM + PostgreSQL を使用したバックエンドAPIサーバーです。

## ディレクトリ構成

### cmd/
アプリケーションのエントリーポイント
- [詳細](cmd/README.md)

### config/
設定ファイル管理
- [詳細](config/README.md)

### db/
データベース関連
- [マイグレーション](db/migrate/README.md)

### di/
依存性注入(DI)設定
- [詳細](di/README.md)

### internal/
コアロジック
- api/: APIハンドラ
- models/: データモデル
- repository/: DBアクセス層
- service/: ビジネスロジック

### scripts/
開発用スクリプト
- [詳細](scripts/README.md)

## 環境構築

```bash
# 依存関係インストール
make setup

# 開発サーバー起動
make run-dev

# テスト実行
make test
```

## APIドキュメント
Swagger UIで確認可能:
```
http://localhost:8080/swagger/index.html
```

## デプロイ方法
```bash
# Dockerイメージビルド
make docker-build

# Railwayにデプロイ
railway up
```

## 技術スタック
- Go 1.20+
- Gin
- GORM
- PostgreSQL
- Goose (マイグレーション)
- Wire (DI)
