# コマンドのメモ

```
docker exec -it rust_pg_db psql -U user -d my_rust_db
# 中でSQLを貼り付けて実行
```

SQLのテーブル作成
```
CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    postal_code TEXT,
    prefecture TEXT,
    address TEXT,
    contact_name TEXT,
    phone TEXT
);

```