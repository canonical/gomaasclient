package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// DNSResources is an interface for listing and creating
// DNSResource records
type DNSResources interface {
	Get(params *entity.DNSResourcesParams) ([]entity.DNSResource, error)
	Create(params *entity.DNSResourceParams) (*entity.DNSResource, error)
}
