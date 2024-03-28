package horreum

import (
	"context"
	"fmt"
	"net/url"

	gocloak "github.com/Nerzal/gocloak/v13"
	"github.com/hyperfoil/horreum-client-golang/pkg/raw_client/models"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

var _ authentication.AccessTokenProvider = KeycloakAccessProvider{}       // Verify that T implements I.
var _ authentication.AccessTokenProvider = (*KeycloakAccessProvider)(nil) // Verify that *T implements I.

type KeycloakAccessProvider struct {
	config   models.KeycloakConfigable
	username string
	password string
	client   *gocloak.GoCloak
}

// NewKeycloakAccessProvider create a new KeycloakAccessProvider instance
func NewKeycloakAccessProvider(config models.KeycloakConfigable, username string, password string) (*KeycloakAccessProvider, error) {
	if config.GetUrl() == nil {
		return nil, fmt.Errorf("missing keycloak url")
	}

	return &KeycloakAccessProvider{
		config:   config,
		username: username,
		password: password,
		client:   gocloak.NewClient(*config.GetUrl()),
	}, nil
}

// GetAllowedHostsValidator implements authentication.AccessTokenProvider.
func (k KeycloakAccessProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	ahv, err := authentication.NewAllowedHostsValidatorErrorCheck([]string{})
	if err != nil {
		return nil
	}

	return ahv
}

// GetAuthorizationToken implements authentication.AccessTokenProvider.
func (k KeycloakAccessProvider) GetAuthorizationToken(context context.Context, url *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	realm := ""
	if k.config.GetRealm() != nil {
		realm = *k.config.GetRealm()
	}
	clientId := ""
	if k.config.GetClientId() != nil {
		clientId = *k.config.GetClientId()
	}

	token, err := k.client.Login(context, clientId, "", realm, k.username, k.password)

	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}
