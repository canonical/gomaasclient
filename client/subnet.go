package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/ionutbalutoiu/gomaasclient/entity"
	"github.com/ionutbalutoiu/gomaasclient/entity/subnet"
)

type Subnet struct {
	ApiClient ApiClient
}

func (s *Subnet) client(id int) ApiClient {
	return s.ApiClient.GetSubObject("subnets").GetSubObject(fmt.Sprintf("%v", id))
}

func (s *Subnet) Delete(id int) error {
	return s.client(id).Delete()
}

func (s *Subnet) Get(id int) (subnet *entity.Subnet, err error) {
	subnet = new(entity.Subnet)
	err = s.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &subnet)
	})
	return
}

func (s *Subnet) GetIPAddresses(id int) (subnetIPAddresses []subnet.IPAddress, err error) {
	qsp := url.Values{}
	qsp.Set("with_username", "1")
	qsp.Set("with_summary", "1")
	err = s.client(id).Get("ip_addresses", qsp, func(data []byte) error {
		return json.Unmarshal(data, &subnetIPAddresses)
	})
	return
}

func (s *Subnet) GetReservedIPRanges(id int) (subnetReservedIPRanges []subnet.ReservedIPRange, err error) {
	err = s.client(id).Get("reserved_ip_ranges", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &subnetReservedIPRanges)
	})
	return
}

func (s *Subnet) GetStatistics(id int) (stats *subnet.Statistics, err error) {
	stats = new(subnet.Statistics)
	err = s.client(id).Get("statistics", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, stats)
	})
	return
}

func (s *Subnet) GetUnreservedIPRanges(id int) (ipRanges []subnet.IPRange, err error) {
	err = s.client(id).Get("unreserved_ip_ranges", url.Values{}, func(data []byte) error {
		log.Printf("%s\n", data)
		return json.Unmarshal(data, &ipRanges)
	})
	return
}

func (s *Subnet) Update(id int, params *entity.SubnetParams) (subnet *entity.Subnet, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	subnet = new(entity.Subnet)
	err = s.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, subnet)
	})
	return
}
