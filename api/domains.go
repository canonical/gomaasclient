package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Domains is an interface for listing and creaing
// Domain records
type Domains interface {
	Get(ctx context.Context) ([]entity.Domain, error)
	Create(ctx context.Context, params *entity.DomainParams) (*entity.Domain, error)
	SetSerial(ctx context.Context, serial int) error
}
