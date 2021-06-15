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
		Fabrics:           &Fabrics{ApiClient: *apiClient},
		VLANs:             &VLANs{ApiClient: *apiClient},
		Machine:           &Machine{ApiClient: *apiClient},
		Machines:          &Machines{ApiClient: *apiClient},
		VMHost:            &VMHost{ApiClient: *apiClient},
		VMHosts:           &VMHosts{ApiClient: *apiClient},
		NetworkInterface:  &NetworkInterface{ApiClient: *apiClient},
		NetworkInterfaces: &NetworkInterfaces{ApiClient: *apiClient},
		Subnets:           &Subnets{ApiClient: *apiClient},
		Tag:               &Tag{ApiClient: *apiClient},
		Tags:              &Tags{ApiClient: *apiClient},
		BlockDevice:       &BlockDevice{ApiClient: *apiClient},
		BlockDevices:      &BlockDevices{ApiClient: *apiClient},
	}
	return &client, nil
}

type Client struct {
	Fabrics           api.Fabrics
	VLANs             api.VLANs
	Machine           api.Machine
	Machines          api.Machines
	VMHost            api.VMHost
	VMHosts           api.VMHosts
	NetworkInterface  api.NetworkInterface
	NetworkInterfaces api.NetworkInterfaces
	Subnets           api.Subnets
	Tag               api.Tag
	Tags              api.Tags
	BlockDevice       api.BlockDevice
	BlockDevices      api.BlockDevices
}

func GetApiClient(apiURL string, apiKey string, apiVersion string) (*ApiClient, error) {
	versionedURL := gomaasapi.AddAPIVersionToURL(apiURL, apiVersion)
	authClient, err := gomaasapi.NewAuthenticatedClient(versionedURL, apiKey)
	if err != nil {
		return nil, err
	}
	return &ApiClient{gomaasapi.NewMAAS(*authClient)}, nil
}

type ApiClient struct {
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
	return ApiClient{&mc}
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
