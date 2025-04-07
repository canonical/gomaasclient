package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// User is an interface for getting and deleting users
type User interface {
	Get(ctx context.Context, userName string) (*entity.User, error)
	Delete(ctx context.Context, userName string) error
}
