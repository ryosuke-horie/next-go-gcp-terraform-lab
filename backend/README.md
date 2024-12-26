# Backendメモ

躓いた部分を中心に記載

## sqlファイルをローカルPostgresファイルに食わせる方法

```bash
cat sql/create-tasks-table.sql | docker compose exec -T db  psql -U postgres -d xodb
```

## xoでモデルを自動生成する

```bash
xo schema postgres://postgres:password@localhost:5432/xodb?sslmode=disable
```
