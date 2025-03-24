package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Users is an interface for listing and creating users
type Users interface {
	Get(ctx context.Context) ([]entity.User, error)
	Create(ctx context.Context, params *entity.UserParams) (*entity.User, error)
}
