package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// SSLKeys is an interface for listing and creating SSLKey objects
type SSLKeys interface {
	Get(ctx context.Context) ([]entity.SSLKey, error)
	Create(ctx context.Context, key string) (*entity.SSLKey, error)
}
