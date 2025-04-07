package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Version is an interface for listing MAAS version details
type Version interface {
	Get(ctx context.Context) (*entity.Version, error)
}
