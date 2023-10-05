# 実行方法

1. [こちらのページ](https://www.keycloak.org/getting-started/getting-started-docker) で Keycloak サービスを立ち上げ
2. `main.go` を実行
3. `http://localhost:8080/sign_up` に POST する

リクエスト例

```
{
  "username": "foobar"
}
```
