package api

import (
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type DNSResource interface {
	Get(id int) (*entity.DNSResource, error)
	Update(id int, params *entity.DNSResourceParams) (*entity.DNSResource, error)
	Delete(id int) error
}
