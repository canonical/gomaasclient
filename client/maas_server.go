package client

import (
	"net/url"
)

// MAASServer contains functionality for manipulating the MAASServer entity.
type MAASServer struct {
	APIClient APIClient
}

func (m *MAASServer) client() APIClient {
	return m.APIClient.GetSubObject("maas")
}

// Get MAAS server configuration value.
func (m *MAASServer) Get(name string) ([]byte, error) {
	qsp := url.Values{}
	qsp.Set("name", name)

	value := new([]byte)
	err := m.client().Get("get_config", qsp, func(data []byte) error {
		*value = data
		return nil
	})

	return *value, err
}

// Post (Set) MAAS server configuration value.
func (m *MAASServer) Post(name, value string) error {
	qsp := url.Values{}
	qsp.Set("name", name)
	qsp.Set("value", value)

	err := m.client().Post("set_config", qsp, func(data []byte) error {
		return nil
	})

	return err
}
