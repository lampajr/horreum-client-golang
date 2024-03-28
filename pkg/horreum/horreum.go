package horreum

import (
	"context"
	"fmt"

	"github.com/hyperfoil/horreum/pkg/raw_client"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	http "github.com/microsoft/kiota-http-go"

	_ "embed"
)

//go:embed version.txt
var version string

type HorreumClient struct {
	baseUrl  string
	username *string
	password *string

	RawClient   *raw_client.HorreumRawClient
	AuthProvder authentication.AuthenticationProvider
}

func setupAuthProvider(baseUrl string, username string, password string) (authentication.AccessTokenProvider, error) {
	anonymousProvider := &authentication.AnonymousAuthenticationProvider{}
	anonymousAdapter, err := http.NewNetHttpRequestAdapter(anonymousProvider)
	if err != nil {
		return nil, fmt.Errorf("error creating anonymous adapter: %w", err)
	}

	anonymousAdapter.SetBaseUrl(baseUrl)
	client := raw_client.NewHorreumRawClient(anonymousAdapter)

	config, err := client.Api().Config().Keycloak().Get(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error retrieving keycloak configuration: %w", err)
	}

	return NewKeycloakAccessProvider(config, username, password)
}

func NewHorreumClient(baseUrl string, username *string, password *string) (*HorreumClient, error) {
	client := &HorreumClient{
		baseUrl:  baseUrl,
		username: username,
		password: password,
	}

	if username != nil && password != nil {
		provider, err := setupAuthProvider(baseUrl, *username, *password)
		if err != nil {
			return nil, fmt.Errorf("error setting up keycloak provider: %w", err)
		}
		client.AuthProvder = authentication.NewBaseBearerTokenAuthenticationProvider(provider)
	} else if password != nil {
		return nil, fmt.Errorf("providing password without username, have you missed something?")
	} else {
		// anonymous authentication
		client.AuthProvder = &authentication.AnonymousAuthenticationProvider{}
	}

	// Disable gzip compression
	// To avoid "HTTP 400 Bad Request Illegal character ((CTRL-CHAR, code 31)): only regular white space (\r, \n, \t) is allowed between tokens"
	compressOpt := http.NewCompressionOptions(false)
	middlewares, err := http.GetDefaultMiddlewaresWithOptions(&compressOpt)
	if err != nil {
		return nil, fmt.Errorf("error creating default middlewares with compression disabled")
	}
	httpClient := http.GetDefaultClient(middlewares...)
	adapter, err := http.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(client.AuthProvder, nil, nil, httpClient)
	if err != nil {
		return nil, fmt.Errorf("error creating client adapter: %w", err)
	}
	adapter.SetBaseUrl(baseUrl)
	client.RawClient = raw_client.NewHorreumRawClient(adapter)

	return client, nil
}

func (h *HorreumClient) Version() string {
	return version
}
