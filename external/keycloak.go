package external

import (
	"context"
	"keycloak-sample/params"

	"github.com/Nerzal/gocloak/v13"
)

const (
	keycloakHost = "http://localhost:3000"
	realmName    = "myrealm"
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
		client: gocloak.NewClient(keycloakHost),
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
		// FirstName: gocloak.StringP(p.FirstName),
		// LastName:  gocloak.StringP(p.LastName),
		// Email:     gocloak.StringP(p.Email),
		Username: gocloak.StringP(p.Username),
		Enabled:  gocloak.BoolP(true),
	}

	// sign up
	if _, err := k.client.CreateUser(ctx, token.AccessToken, realmName, req); err != nil {
		return err
	}
	return nil
}
