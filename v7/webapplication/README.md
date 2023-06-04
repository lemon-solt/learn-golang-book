# postgres
```

-- user一覧
SELECT * FROM pg_user;

-- user作成
createuser --interactive gwp

-- db作成
createdb -O gwp gwp

-- パスワード変更
ALTER USER gwp WITH PASSWORD 'gwp';

```