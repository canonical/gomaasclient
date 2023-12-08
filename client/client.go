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
	apiClient, err := getApiClient(apiURL, apiKey, apiVersion, tlsConfig)
	if err != nil {
		return nil, err
	}

	return constructClient(apiClient), nil
}

// GetClient creates a client
func GetClient(apiURL string, apiKey string, apiVersion string) (*Client, error) {
	apiClient, err := getApiClient(apiURL, apiKey, apiVersion, nil)
	if err != nil {
		return nil, err
	}

	return constructClient(apiClient), nil
}

func constructClient(apiClient *ApiClient) *Client {
	client := Client{
		Device:                &Device{ApiClient: *apiClient},
		Devices:               &Devices{ApiClient: *apiClient},
		Domain:                &Domain{ApiClient: *apiClient},
		Domains:               &Domains{ApiClient: *apiClient},
		DNSResource:           &DNSResource{ApiClient: *apiClient},
		DNSResources:          &DNSResources{ApiClient: *apiClient},
		DNSResourceRecord:     &DNSResourceRecord{ApiClient: *apiClient},
		DNSResourceRecords:    &DNSResourceRecords{ApiClient: *apiClient},
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
		RAID:                  &RAID{ApiClient: *apiClient},
		RAIDs:                 &RAIDs{ApiClient: *apiClient},
		Subnet:                &Subnet{ApiClient: *apiClient},
		Subnets:               &Subnets{ApiClient: *apiClient},
		IPRange:               &IPRange{ApiClient: *apiClient},
		IPRanges:              &IPRanges{ApiClient: *apiClient},
		IPAddresses:           &IPAddresses{ApiClient: *apiClient},
		Tag:                   &Tag{ApiClient: *apiClient},
		Tags:                  &Tags{ApiClient: *apiClient},
		BlockDevice:           &BlockDevice{ApiClient: *apiClient},
		BlockDevices:          &BlockDevices{ApiClient: *apiClient},
		BlockDevicePartition:  &BlockDevicePartition{ApiClient: *apiClient},
		BlockDevicePartitions: &BlockDevicePartitions{ApiClient: *apiClient},
		User:                  &User{ApiClient: *apiClient},
		Users:                 &Users{ApiClient: *apiClient},
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
}

// Deprecated: The gomaasapi client will be no longer exposed.
// Instead, please use GetClient which instantiate all MAAS resources endpoints with gomaasapi client.
func GetApiClient(apiURL string, apiKey string, apiVersion string) (*ApiClient, error) {
	return getApiClient(apiURL, apiKey, apiVersion, nil)
}

func getApiClient(apiURL string, apiKey string, apiVersion string, tlsConfig *tls.Config) (*ApiClient, error) {
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
