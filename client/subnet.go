package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/canonical/gomaasclient/entity/subnet"
	"github.com/google/go-querystring/query"
)

// Subnet implements api.Subnet
type Subnet struct {
	APIClient APIClient
}

func (s *Subnet) client(id int) *APIClient {
	return s.APIClient.SubClient("subnets").SubClient(fmt.Sprintf("%v", id))
}

// Delete deletes a given Subnet
func (s *Subnet) Delete(ctx context.Context, id int) error {
	return s.client(id).Delete(ctx)
}

// Get fetches a given Subnet
func (s *Subnet) Get(ctx context.Context, id int) (*entity.Subnet, error) {
	subnet := new(entity.Subnet)
	err := s.client(id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &subnet)
	})

	return subnet, err
}

// GetIPAddresses fetches the allocated IPAddresses in the given subnet
func (s *Subnet) GetIPAddresses(ctx context.Context, id int) ([]subnet.IPAddress, error) {
	qsp := url.Values{}
	qsp.Set("with_username", "1")
	qsp.Set("with_summary", "1")

	subnetIPAddresses := make([]subnet.IPAddress, 0)
	err := s.client(id).Get(ctx, "ip_addresses", qsp, func(data []byte) error {
		return json.Unmarshal(data, &subnetIPAddresses)
	})

	return subnetIPAddresses, err
}

// GetReservedIPRanges fetches the reserved IPRange objects for the given Subnet
func (s *Subnet) GetReservedIPRanges(ctx context.Context, id int) ([]subnet.ReservedIPRange, error) {
	subnetReservedIPRanges := make([]subnet.ReservedIPRange, 0)
	err := s.client(id).Get(ctx, "reserved_ip_ranges", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &subnetReservedIPRanges)
	})

	return subnetReservedIPRanges, err
}

// GetStatistics gets the stats of the given subnet
func (s *Subnet) GetStatistics(ctx context.Context, id int) (*subnet.Statistics, error) {
	stats := new(subnet.Statistics)
	err := s.client(id).Get(ctx, "statistics", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, stats)
	})

	return stats, err
}

// GetUnreservedIPRanges gets the IPRange objects for allocation use
func (s *Subnet) GetUnreservedIPRanges(ctx context.Context, id int) ([]subnet.IPRange, error) {
	ipRanges := make([]subnet.IPRange, 0)
	err := s.client(id).Get(ctx, "unreserved_ip_ranges", url.Values{}, func(data []byte) error {
		log.Printf("%s\n", data)
		return json.Unmarshal(data, &ipRanges)
	})

	return ipRanges, err
}

// Update updates the given Subnet
func (s *Subnet) Update(ctx context.Context, id int, params *entity.SubnetParams) (*entity.Subnet, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	subnet := new(entity.Subnet)
	err = s.client(id).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, subnet)
	})

	return subnet, err
}
