# リポジトリ層

このディレクトリにはデータベース操作を抽象化したリポジトリ実装が含まれています。GORMを使用してデータアクセスを行います。

## 実装パターン

1. **CRUD操作**:
   - Create: `db.Create()`
   - Read: `db.First()`, `db.Find()`
   - Update: `db.Save()`, `db.Update()`
   - Delete: `db.Delete()`

2. **トランザクション管理**:
```go
err := r.db.Transaction(func(tx *gorm.DB) error {
    // トランザクション内の処理
    if err := tx.Create(&data1).Error; err != nil {
        return err
    }
    return tx.Update(&data2).Error
})
```

3. **複雑なクエリ**:
   - スコープを使用してクエリを組み立て
   - プリロードで関連データを取得
   - ページネーション実装

## ベストプラクティス
- インターフェースでリポジトリを抽象化
- トランザクション境界を明確に
- エラーハンドリングを徹底
- テスト用にモック実装を準備
