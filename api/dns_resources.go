package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type DNSResources interface {
	Get() ([]entity.DNSResource, error)
	Create(params *entity.DNSResourceParams) (*entity.DNSResource, error)
}
