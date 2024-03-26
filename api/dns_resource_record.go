package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// DNSResourceRecord is an interface defining API behaviour for
// DNS resource records
type DNSResourceRecord interface {
	Get(id int) (*entity.DNSResourceRecord, error)
	Update(id int, params *entity.DNSResourceRecordParams) (*entity.DNSResourceRecord, error)
	Delete(id int) error
}
