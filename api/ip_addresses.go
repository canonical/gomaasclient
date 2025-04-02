package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// IPAddresses is an interface defining API behaviour for IP addresses
type IPAddresses interface {
	Get(ctx context.Context, params *entity.IPAddressesParams) ([]entity.IPAddress, error)
	Release(ctx context.Context, params *entity.IPAddressesParams) error
	Reserve(ctx context.Context, params *entity.IPAddressesParams) (*entity.IPAddress, error)
}
