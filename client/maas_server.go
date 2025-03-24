package client

import (
	"context"
	"net/url"
)

// MAASServer contains functionality for manipulating the MAASServer entity.
type MAASServer struct {
	APIClient APIClient
}

func (m *MAASServer) client() *APIClient {
	return m.APIClient.SubClient("maas")
}

// Get MAAS server configuration value.
func (m *MAASServer) Get(ctx context.Context, name string) ([]byte, error) {
	qsp := url.Values{}
	qsp.Set("name", name)

	value := new([]byte)
	err := m.client().Get(ctx, "get_config", qsp, func(data []byte) error {
		*value = data
		return nil
	})

	return *value, err
}

// Post (Set) MAAS server configuration value.
func (m *MAASServer) Post(ctx context.Context, name, value string) error {
	qsp := url.Values{}
	qsp.Set("name", name)
	qsp.Set("value", value)

	err := m.client().Post(ctx, "set_config", qsp, func(data []byte) error {
		return nil
	})

	return err
}
