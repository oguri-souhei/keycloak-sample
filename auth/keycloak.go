package auth

import (
	"context"
	"errors"
	"keycloak-sample/params"

	"github.com/Nerzal/gocloak/v13"
)

const (
	KeycloakHost = "http://localhost:3000"
	RealmName    = "myrealm"
	clientID     = "myclient"
	clientSecret = ""
)

var client *KeycloakClient

type KeycloakClient struct {
	// 認証情報を格納
	token *gocloak.JWT

	// keycloak client
	client *gocloak.GoCloak
}

func GetKeycloakClient() *KeycloakClient {
	// 初期化済みならそのまま返す
	if client != nil {
		return client
	}

	client = &KeycloakClient{
		client: gocloak.NewClient(KeycloakHost),
	}

	return client
}

func (k *KeycloakClient) SignUp(ctx context.Context, p params.SignUp) error {
	adminToken, err := k.client.LoginAdmin(ctx, "admin", "admin", "master")
	if err != nil {
		panic(err)
	}

	// create params
	req := gocloak.User{
		Username: gocloak.StringP(p.Username),
		Enabled:  gocloak.BoolP(true),
	}

	// sign up
	if _, err := k.client.CreateUser(ctx, adminToken.AccessToken, RealmName, req); err != nil {
		return err
	}

	return nil
}

func (k *KeycloakClient) Login(ctx context.Context, p params.Login) error {
	token, err := k.client.Login(ctx, clientID, clientSecret, RealmName, p.Username, p.Password)
	if err != nil {
		return err
	}

	k.token = token
	return nil
}

// 認証情報が正しいかどうかを確認する
func (k *KeycloakClient) ValidateToken(ctx context.Context) error {
	if k.token == nil {
		return errors.New("unauthorized: token is nil")
	}

	userInfo, err := k.client.GetUserInfo(ctx, k.token.AccessToken, RealmName)
	if err != nil {
		return err
	}

	if userInfo == nil {
		return errors.New("unauthorized")
	}
	return nil
}
