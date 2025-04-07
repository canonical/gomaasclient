package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// SSHKey is an interface defining API behaviour for SSHKey objects
type SSHKey interface {
	Get(ctx context.Context, id int) (*entity.SSHKey, error)
	Delete(ctx context.Context, id int) error
}
