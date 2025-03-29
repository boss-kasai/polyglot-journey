# Goでの開発におけるメモ

## 忘れがちなコマンド

テスト実行コード

```:go
go test -v ./tests/integration/...
```

キャッシュの削除

```:go
go clean -testcache
```

## Makefileの実行コマンド

```bash
make all          # すべてのチェック＋ビルド
make format       # コード整形とimportの最適化
make vet          # 静的検査（構文・タグ・戻り値など）
make lint         # 総合的なコードチェック
make staticcheck  # より深い静的解析（未使用・非推奨コードなど）
make tidy         # モジュール依存を整理
make build        # バイナリを生成（bin/app）
make test         # 単体・統合テストを実行
make clean        # ビルド成果物を削除
```
