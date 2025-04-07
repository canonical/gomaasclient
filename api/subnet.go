package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
	"github.com/canonical/gomaasclient/entity/subnet"
)

// Subnet represents the MAAS Subnet endpoint
type Subnet interface {
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*entity.Subnet, error)
	GetIPAddresses(ctx context.Context, id int) ([]subnet.IPAddress, error)
	GetReservedIPRanges(ctx context.Context, id int) ([]subnet.ReservedIPRange, error)
	GetStatistics(ctx context.Context, id int) (*subnet.Statistics, error)
	GetUnreservedIPRanges(ctx context.Context, id int) ([]subnet.IPRange, error)
	Update(ctx context.Context, id int, params *entity.SubnetParams) (*entity.Subnet, error)
}
