package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// Subnets represents the MAAS Subnets endpoint
type Subnets interface {
	Get(ctx context.Context) ([]entity.Subnet, error)
	Create(ctx context.Context, params *entity.SubnetParams) (*entity.Subnet, error)
}
