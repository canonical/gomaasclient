package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Account is an interface for managing user account
type Account interface {
	CreateAuthorisationToken(ctx context.Context, name string) (*entity.AuthorisationToken, error)
	DeleteAuthorisationToken(ctx context.Context, key string) error
	ListAuthorisationTokens(ctx context.Context) ([]entity.AuthorisationTokenListItem, error)
	UpdateTokenName(ctx context.Context, name, token string) error
}
