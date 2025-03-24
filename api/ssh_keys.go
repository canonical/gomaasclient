package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// SSHKeys is an interface for listing and creating SSHKey objects
type SSHKeys interface {
	Get(ctx context.Context) ([]entity.SSHKey, error)
	Create(ctx context.Context, key string) (*entity.SSHKey, error)
	Import(ctx context.Context, keysource string) ([]entity.SSHKey, error)
}
