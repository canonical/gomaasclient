package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type DNSResourceRecords interface {
	Get() ([]entity.DNSResourceRecord, error)
	Create(params *entity.DNSResourceRecordParams) (*entity.DNSResourceRecord, error)
}
