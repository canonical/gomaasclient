package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// SSLKey is an interface defining API behaviour for SSLKey objects
type SSLKey interface {
	Get(ctx context.Context, id int) (*entity.SSLKey, error)
	Delete(ctx context.Context, id int) error
}
