package external

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
	client *gocloak.GoCloak
}

func GetKeycloakClient() *KeycloakClient {
	if client != nil {
		return client
	}

	client = &KeycloakClient{
		client: gocloak.NewClient(KeycloakHost),
	}

	return client
}

func (k *KeycloakClient) SignUp(ctx context.Context, p params.SignUp) error {
	token, err := k.client.LoginAdmin(ctx, "admin", "admin", "master")
	if err != nil {
		panic(err)
	}

	// create params
	req := gocloak.User{
		Username: gocloak.StringP(p.Username),
		Enabled:  gocloak.BoolP(true),
	}

	// sign up
	if _, err := k.client.CreateUser(ctx, token.AccessToken, RealmName, req); err != nil {
		return err
	}
	return nil
}

func (k *KeycloakClient) Login(ctx context.Context, p params.Login) (*gocloak.JWT, error) {
	return k.client.Login(ctx, clientID, clientSecret, RealmName, p.Username, p.Password)
}

// 認証情報が正しいかどうかを確認する
func (k *KeycloakClient) ValidateToken(ctx context.Context, token *gocloak.JWT) error {
	if token == nil {
		return errors.New("unauthorized: token is nil")
	}

	userInfo, err := k.client.GetUserInfo(ctx, token.AccessToken, RealmName)
	if err != nil {
		return err
	}
	if userInfo == nil {
		return errors.New("unauthorized")
	}
	return nil
}
