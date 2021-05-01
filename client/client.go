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
		Machine:  &Machine{ApiClient: *apiClient},
		Machines: &Machines{ApiClient: *apiClient},
		Pod:      &Pod{ApiClient: *apiClient},
		Pods:     &Pods{ApiClient: *apiClient},
		Tag:      &Tag{ApiClient: *apiClient},
		Tags:     &Tags{ApiClient: *apiClient},
	}
	return &client, nil
}

type Client struct {
	Machine  api.Machine
	Machines api.Machines
	Pod      api.Pod
	Pods     api.Pods
	Tag      api.Tag
	Tags     api.Tags
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
