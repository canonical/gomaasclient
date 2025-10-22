package client

import (
	"encoding/json"
	"fmt"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// NodeResults implements api.NodeResults
type NodeResults struct {
	APIClient APIClient
}

func (c *NodeResults) client(systemID string) APIClient {
	return c.APIClient.GetSubObject("nodes").GetSubObject(fmt.Sprintf("%v", systemID)).GetSubObject("results")
}

// Get node results
func (c *NodeResults) Get(systemID string, params *entity.NodeResultParams) ([]entity.NodeResult, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	results := make([]entity.NodeResult, 0)
	err = c.client(systemID).Get("", qsp, func(data []byte) error {
		return json.Unmarshal(data, &results)
	})

	return results, err
}
