package api

import "github.com/canonical/gomaasclient/entity"

// Account is an interface for managing user account
type Account interface {
	CreateAuthorisationToken(name string) (*entity.AuthorisationToken, error)
	DeleteAuthorisationToken(key string) error
	ListAuthorisationTokens() ([]entity.AuthorisationTokenListItem, error)
	UpdateTokenName(name, token string) error
}
