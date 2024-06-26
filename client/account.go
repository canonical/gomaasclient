package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// Account implements api.Account
type Account struct {
	APIClient APIClient
}

func (a *Account) client() APIClient {
	return a.APIClient.GetSubObject("account")
}

// CreateAuthorisationToken creates authorisation token with provided name
func (a *Account) CreateAuthorisationToken(name string) (*entity.AuthorisationToken, error) {
	qsp := url.Values{}
	qsp.Set("name", name)

	authToken := new(entity.AuthorisationToken)
	err := a.client().Post("create_authorisation_token", qsp, func(data []byte) error {
		return json.Unmarshal(data, authToken)
	})

	return authToken, err
}

// DeleteAuthorisationToken deletes authorisation token with provided key
func (a *Account) DeleteAuthorisationToken(key string) error {
	qsp := url.Values{}
	qsp.Set("token_key", key)
	err := a.client().Post("delete_authorisation_token", qsp, func(data []byte) error { return nil })

	return err
}

// ListAuthorisationTokens Lists authorisation tokens
func (a *Account) ListAuthorisationTokens() ([]entity.AuthorisationTokenListItem, error) {
	authToken := make([]entity.AuthorisationTokenListItem, 0)
	err := a.client().Get("list_authorisation_tokens", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &authToken)
	})

	return authToken, err
}

// UpdateTokenName updates given token with provided name
func (a *Account) UpdateTokenName(name, token string) error {
	qsp := url.Values{}
	qsp.Set("name", name)
	qsp.Set("token", token)
	err := a.client().Post("update_token_name", qsp, func(data []byte) error { return nil })

	return err
}
