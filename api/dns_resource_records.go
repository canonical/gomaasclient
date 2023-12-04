package api

import (
	"github.com/maas/gomaasclient/entity"
)

// DNSResourceRecords is an interface for listing and creaing
// DNSResourceRecord records
type DNSResourceRecords interface {
	Get() ([]entity.DNSResourceRecord, error)
	Create(params *entity.DNSResourceRecordParams) (*entity.DNSResourceRecord, error)
}
