package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// DNSResourceRecord is an interface defining API behaviour for
// DNS resource records
type DNSResourceRecord interface {
	Get(ctx context.Context, id int) (*entity.DNSResourceRecord, error)
	Update(ctx context.Context, id int, params *entity.DNSResourceRecordParams) (*entity.DNSResourceRecord, error)
	Delete(ctx context.Context, id int) error
}
