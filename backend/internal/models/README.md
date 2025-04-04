# データモデル層

このディレクトリにはデータベースのテーブル構造に対応するモデル定義が含まれています。GORMを使用したモデル定義を行います。

## モデル定義のポイント

1. **基本フィールド**:
   - `gorm.Model`を埋め込むことで、ID, CreatedAt, UpdatedAt, DeletedAtを自動追加
   - フィールドごとにGORMタグでデータベース属性を指定

2. **GORMタグ**:
   - `size`: 文字列フィールドの最大長
   - `not null`: NULLを許可しない
   - `unique`: ユニーク制約
   - `index`: インデックス作成
   - `default`: デフォルト値

3. **JSONタグ**:
   - APIレスポンス時のフィールド名を指定
   - `omitempty`: 空値の場合フィールドを除外

4. **リレーション定義**:
   - `belongs_to`, `has_many`などの関連を定義可能
   - 外部キー制約を設定

## バリデーション
- モデルレベルでのバリデーションを実装
- カスタムバリデーション関数を定義可能
- ビジネスルールに基づくチェックを実施
