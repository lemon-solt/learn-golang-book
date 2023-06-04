```
# user作成
createuser --interactive gouser

# db作成
createdb -O gouser gwp

# testdbにログイン
psql gwp
```