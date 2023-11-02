## メモ

### echo インストール
```bash
$ cd myapp
$ go mod init myapp
$ go get github.com/labstack/echo/v4
```

### 起動方法

```bash
$ go run server.go
```

.devcontainerのdocker-compose.ymlにポート定義あるが、現時点では localhost:18080 をたたけばいい
