// Package client contains the implementation of CRUD operations on MAAS resources.
package client

import (
	"crypto/tls"
	"net/http"
	"net/url"

	gomaasapi "github.com/juju/gomaasapi/v2"
	"github.com/maas/gomaasclient/api"
)

// GetTLSClient creates a Client configured with TLS
func GetTLSClient(apiURL string, apiKey string, apiVersion string, tlsConfig *tls.Config) (*Client, error) {
	apiClient, err := getAPIClient(apiURL, apiKey, apiVersion, tlsConfig)
	if err != nil {
		return nil, err
	}

	return constructClient(apiClient), nil
}

// GetClient creates a client
func GetClient(apiURL string, apiKey string, apiVersion string) (*Client, error) {
	apiClient, err := getAPIClient(apiURL, apiKey, apiVersion, nil)
	if err != nil {
		return nil, err
	}

	return constructClient(apiClient), nil
}

func constructClient(apiClient *APIClient) *Client {
	client := Client{
		Device:                &Device{APIClient: *apiClient},
		Devices:               &Devices{APIClient: *apiClient},
		Domain:                &Domain{APIClient: *apiClient},
		Domains:               &Domains{APIClient: *apiClient},
		DNSResource:           &DNSResource{APIClient: *apiClient},
		DNSResources:          &DNSResources{APIClient: *apiClient},
		DNSResourceRecord:     &DNSResourceRecord{APIClient: *apiClient},
		DNSResourceRecords:    &DNSResourceRecords{APIClient: *apiClient},
		Fabric:                &Fabric{APIClient: *apiClient},
		Fabrics:               &Fabrics{APIClient: *apiClient},
		VLAN:                  &VLAN{APIClient: *apiClient},
		VLANs:                 &VLANs{APIClient: *apiClient},
		Space:                 &Space{APIClient: *apiClient},
		Spaces:                &Spaces{APIClient: *apiClient},
		Machine:               &Machine{APIClient: *apiClient},
		Machines:              &Machines{APIClient: *apiClient},
		VMHost:                &VMHost{APIClient: *apiClient},
		VMHosts:               &VMHosts{APIClient: *apiClient},
		NetworkInterface:      &NetworkInterface{APIClient: *apiClient},
		NetworkInterfaces:     &NetworkInterfaces{APIClient: *apiClient},
		RAID:                  &RAID{APIClient: *apiClient},
		RAIDs:                 &RAIDs{APIClient: *apiClient},
		Subnet:                &Subnet{APIClient: *apiClient},
		Subnets:               &Subnets{APIClient: *apiClient},
		IPRange:               &IPRange{APIClient: *apiClient},
		IPRanges:              &IPRanges{APIClient: *apiClient},
		IPAddresses:           &IPAddresses{APIClient: *apiClient},
		Tag:                   &Tag{APIClient: *apiClient},
		Tags:                  &Tags{APIClient: *apiClient},
		BlockDevice:           &BlockDevice{APIClient: *apiClient},
		BlockDevices:          &BlockDevices{APIClient: *apiClient},
		BlockDevicePartition:  &BlockDevicePartition{APIClient: *apiClient},
		BlockDevicePartitions: &BlockDevicePartitions{APIClient: *apiClient},
		User:                  &User{APIClient: *apiClient},
		Users:                 &Users{APIClient: *apiClient},
		ResourcePool:          &ResourcePool{APIClient: *apiClient},
		ResourcePools:         &ResourcePools{APIClient: *apiClient},
		MAASServer:            &MAASServer{APIClient: *apiClient},
	}

	return &client
}

// Client is an object providing API interactions
// with a configured MAAS installation
type Client struct {
	Device                api.Device
	Devices               api.Devices
	DNSResource           api.DNSResource
	DNSResources          api.DNSResources
	DNSResourceRecord     api.DNSResourceRecord
	DNSResourceRecords    api.DNSResourceRecords
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
	RAID                  api.RAID
	RAIDs                 api.RAIDs
	Subnet                api.Subnet
	Subnets               api.Subnets
	IPRange               api.IPRange
	IPRanges              api.IPRanges
	IPAddresses           api.IPAddresses
	Tag                   api.Tag
	Tags                  api.Tags
	BlockDevice           api.BlockDevice
	BlockDevices          api.BlockDevices
	BlockDevicePartition  api.BlockDevicePartition
	BlockDevicePartitions api.BlockDevicePartitions
	User                  api.User
	Users                 api.Users
	ResourcePool          api.ResourcePool
	ResourcePools         api.ResourcePools
	MAASServer            api.MAASServer
}

// GetAPIClient returns a MAAS API client.
//
// Deprecated: The gomaasapi client will be no longer exposed.
// Instead, please use GetClient which instantiate all MAAS resources endpoints with gomaasapi client.
func GetAPIClient(apiURL string, apiKey string, apiVersion string) (*APIClient, error) {
	return getAPIClient(apiURL, apiKey, apiVersion, nil)
}

func getAPIClient(apiURL string, apiKey string, apiVersion string, tlsConfig *tls.Config) (*APIClient, error) {
	versionedURL := gomaasapi.AddAPIVersionToURL(apiURL, apiVersion)

	authClient, err := gomaasapi.NewAuthenticatedClient(versionedURL, apiKey)
	if err != nil {
		return nil, err
	}

	if tlsConfig != nil {
		tr := &http.Transport{TLSClientConfig: tlsConfig}
		httpClient := &http.Client{Transport: tr}
		authClient.HTTPClient = httpClient
	}

	return &APIClient{*authClient, gomaasapi.NewMAAS(*authClient)}, nil
}

type APIClient struct {
	AuthClient gomaasapi.Client
	*gomaasapi.MAASObject
}

func (c APIClient) Get(op string, params url.Values, f func([]byte) error) error {
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

func (c APIClient) GetSubObject(name string) APIClient {
	mc := c.MAASObject.GetSubObject(name)
	return APIClient{c.AuthClient, &mc}
}

func (c APIClient) Post(op string, params url.Values, f func([]byte) error) error {
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

func (c APIClient) Put(params url.Values, f func([]byte) error) error {
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
