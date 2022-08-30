package api

import (
	"github.com/maas/gomaasclient/entity"
)

type DNSResources interface {
	Get() ([]entity.DNSResource, error)
	Create(params *entity.DNSResourceParams) (*entity.DNSResource, error)
}
