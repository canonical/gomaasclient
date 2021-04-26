package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
	"github.com/ionutbalutoiu/gomaasclient/api/endpoint/subnet"
)

// Subnet represents the MaaS Subnet endpoint
type Subnet interface {
	Delete(id int) error
	Get(id int) (*endpoint.Subnet, error)
	GetIPAddresses(id int, WithUsername, WithSummary bool) ([]subnet.IPAddress, error)
	GetReservedIPRanges(id int) ([]subnet.ReservedIPRange, error)
	GetStatistics(id int, IncludeRanges, IncludeSuggestions bool) (*subnet.Statistics, error)
	GetUnreservedIPRanges(id int) ([]subnet.IPRange, error)
	Put(id int, params *endpoint.SubnetParams) (*endpoint.Subnet, error)
}
