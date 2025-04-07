package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

type Discoveries interface {
	Get(ctx context.Context) ([]entity.Discovery, error)
	GetByUnknownIP(ctx context.Context) ([]entity.Discovery, error)
	GetByUnknownMAC(ctx context.Context) ([]entity.Discovery, error)
	GetByUnknownIpAndMAC(ctx context.Context) ([]entity.Discovery, error)
	Clear(ctx context.Context, params *entity.DiscoveryClearParams) error
	ClearByMACAndIP(ctx context.Context, mac, ip string) error
}
