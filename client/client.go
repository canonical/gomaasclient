package client

import (
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/api"
	"github.com/juju/gomaasapi"
)

func GetClient(apiURL string, apiKey string, apiVersion string) (*Client, error) {
	apiClient, err := GetApiClient(apiURL, apiKey, apiVersion)
	if err != nil {
		return nil, err
	}
	client := Client{
		Domain:                &Domain{ApiClient: *apiClient},
		Domains:               &Domains{ApiClient: *apiClient},
		Fabric:                &Fabric{ApiClient: *apiClient},
		Fabrics:               &Fabrics{ApiClient: *apiClient},
		VLAN:                  &VLAN{ApiClient: *apiClient},
		VLANs:                 &VLANs{ApiClient: *apiClient},
		Space:                 &Space{ApiClient: *apiClient},
		Spaces:                &Spaces{ApiClient: *apiClient},
		Machine:               &Machine{ApiClient: *apiClient},
		Machines:              &Machines{ApiClient: *apiClient},
		VMHost:                &VMHost{ApiClient: *apiClient},
		VMHosts:               &VMHosts{ApiClient: *apiClient},
		NetworkInterface:      &NetworkInterface{ApiClient: *apiClient},
		NetworkInterfaces:     &NetworkInterfaces{ApiClient: *apiClient},
		Subnet:                &Subnet{ApiClient: *apiClient},
		Subnets:               &Subnets{ApiClient: *apiClient},
		IPRange:               &IPRange{ApiClient: *apiClient},
		IPRanges:              &IPRanges{ApiClient: *apiClient},
		Tag:                   &Tag{ApiClient: *apiClient},
		Tags:                  &Tags{ApiClient: *apiClient},
		BlockDevice:           &BlockDevice{ApiClient: *apiClient},
		BlockDevices:          &BlockDevices{ApiClient: *apiClient},
		BlockDevicePartition:  &BlockDevicePartition{ApiClient: *apiClient},
		BlockDevicePartitions: &BlockDevicePartitions{ApiClient: *apiClient},
	}
	return &client, nil
}

type Client struct {
	Domain                api.Domain
	Domains               api.Domains
	Fabric                api.Fabric
	Fabrics               api.Fabrics
	VLAN                  api.VLAN
	VLANs                 api.VLANs
	Space                 api.Space
	Spaces                api.Spaces
	Machine               api.Machine
	Machines              api.Machines
	VMHost                api.VMHost
	VMHosts               api.VMHosts
	NetworkInterface      api.NetworkInterface
	NetworkInterfaces     api.NetworkInterfaces
	Subnet                api.Subnet
	Subnets               api.Subnets
	IPRange               api.IPRange
	IPRanges              api.IPRanges
	Tag                   api.Tag
	Tags                  api.Tags
	BlockDevice           api.BlockDevice
	BlockDevices          api.BlockDevices
	BlockDevicePartition  api.BlockDevicePartition
	BlockDevicePartitions api.BlockDevicePartitions
}

func GetApiClient(apiURL string, apiKey string, apiVersion string) (*ApiClient, error) {
	versionedURL := gomaasapi.AddAPIVersionToURL(apiURL, apiVersion)
	authClient, err := gomaasapi.NewAuthenticatedClient(versionedURL, apiKey)
	if err != nil {
		return nil, err
	}
	return &ApiClient{*authClient, gomaasapi.NewMAAS(*authClient)}, nil
}

type ApiClient struct {
	AuthClient gomaasapi.Client
	*gomaasapi.MAASObject
}

func (c ApiClient) Get(op string, params url.Values, f func([]byte) error) error {
	res, err := c.CallGet(op, params)
	if err != nil {
		return err
	}
	data, err := res.GetBytes()
	if err != nil {
		return err
	}
	return f(data)
}

func (c ApiClient) GetSubObject(name string) ApiClient {
	mc := c.MAASObject.GetSubObject(name)
	return ApiClient{c.AuthClient, &mc}
}

func (c ApiClient) Post(op string, params url.Values, f func([]byte) error) error {
	res, err := c.CallPost(op, params)
	if err != nil {
		return err
	}
	data, err := res.GetBytes()
	if err != nil {
		return err
	}
	return f(data)
}

func (c ApiClient) Put(params url.Values, f func([]byte) error) error {
	res, err := c.Update(params)
	if err != nil {
		return err
	}
	data, err := res.MarshalJSON()
	if err != nil {
		return err
	}
	return f(data)
}
