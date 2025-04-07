package client

import (
	"context"
	"encoding/json"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// NodeResults implements api.NodeResults
type NodeResults struct {
	APIClient APIClient
}

func (c *NodeResults) client() *APIClient {
	return c.APIClient.SubClient("installation-results")
}

// Get node results
func (c *NodeResults) Get(ctx context.Context, params *entity.NodeResultParams) ([]entity.NodeResult, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	results := make([]entity.NodeResult, 0)
	err = c.client().Get(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, &results)
	})

	return results, err
}
