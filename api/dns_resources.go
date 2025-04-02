package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// DNSResources is an interface for listing and creating
// DNSResource records
type DNSResources interface {
	Get(ctx context.Context, params *entity.DNSResourcesParams) ([]entity.DNSResource, error)
	Create(ctx context.Context, params *entity.DNSResourceParams) (*entity.DNSResource, error)
}
