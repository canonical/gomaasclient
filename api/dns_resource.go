package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// DNSResource is an interface defining API behaviour
// for DNS resources
type DNSResource interface {
	Get(ctx context.Context, id int) (*entity.DNSResource, error)
	Update(ctx context.Context, id int, params *entity.DNSResourceParams) (*entity.DNSResource, error)
	Delete(ctx context.Context, id int) error
}
