# 実行方法

1. [こちらのページ](https://www.keycloak.org/getting-started/getting-started-docker) で Keycloak サービスを立ち上げ
2. `main.go` を実行

# API

- GET /

  - ログインしている場合：hello world の文字がレスポンス
  - ログインしていない場合：401 がかえる

- POST /sign_up

```json
{
  "username": "foobar"
}
```

- POST /login

```json
{
  "username": "foobar",
  "password": "password"
}
```
