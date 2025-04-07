package api

import (
	"context"

	"github.com/canonical/gomaasclient/entity"
)

// DNSResourceRecords is an interface for listing and creaing
// DNSResourceRecord records
type DNSResourceRecords interface {
	Get(ctx context.Context, params *entity.DNSResourceRecordsParams) ([]entity.DNSResourceRecord, error)
	Create(ctx context.Context, params *entity.DNSResourceRecordParams) (*entity.DNSResourceRecord, error)
}
