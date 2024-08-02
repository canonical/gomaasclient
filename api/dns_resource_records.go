package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// DNSResourceRecords is an interface for listing and creaing
// DNSResourceRecord records
type DNSResourceRecords interface {
	Get(params *entity.DNSResourceRecordsParams) ([]entity.DNSResourceRecord, error)
	Create(params *entity.DNSResourceRecordParams) (*entity.DNSResourceRecord, error)
}
