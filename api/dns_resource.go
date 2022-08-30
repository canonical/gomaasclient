package api

import (
	"github.com/maas/gomaasclient/entity"
)

type DNSResource interface {
	Get(id int) (*entity.DNSResource, error)
	Update(id int, params *entity.DNSResourceParams) (*entity.DNSResource, error)
	Delete(id int) error
}
