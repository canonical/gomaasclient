package api

import (
	"github.com/maas/gomaasclient/entity"
)

type DNSResourceRecord interface {
	Get(id int) (*entity.DNSResourceRecord, error)
	Update(id int, params *entity.DNSResourceRecordParams) (*entity.DNSResourceRecord, error)
	Delete(id int) error
}
