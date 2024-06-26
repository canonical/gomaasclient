package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// DNSResource is an interface defining API behaviour
// for DNS resources
type DNSResource interface {
	Get(id int) (*entity.DNSResource, error)
	Update(id int, params *entity.DNSResourceParams) (*entity.DNSResource, error)
	Delete(id int) error
}
