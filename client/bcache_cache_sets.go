//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/canonical/gomaasclient/entity"
)

// BCacheCacheSets contains functionality for manipulating the BCacheCacheSets entity.
type BCacheCacheSets struct {
	APIClient APIClient
}

func (b *BCacheCacheSets) client(systemID string) *APIClient {
	return b.APIClient.SubClient("nodes").SubClient(systemID).SubClient("bcache-cache-sets")
}

// Get BCacheCacheSets of a machine.
func (b *BCacheCacheSets) Get(ctx context.Context, systemID string) ([]entity.BCacheCacheSet, error) {
	bCacheCacheSets := make([]entity.BCacheCacheSet, 0)
	err := b.client(systemID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &bCacheCacheSets)
	})

	return bCacheCacheSets, err
}

// Create a BCacheCacheSet of a machine.
func (b *BCacheCacheSets) Create(ctx context.Context, systemID string, params *entity.BCacheCacheSetParams) (*entity.BCacheCacheSet, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	bCacheCacheSet := new(entity.BCacheCacheSet)
	err = b.client(systemID).Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, bCacheCacheSet)
	})

	return bCacheCacheSet, err
}
