//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// BCacheCacheSet Contains functionality for manipulating the BCacheCacheSet entity.
type BCacheCacheSet struct {
	APIClient APIClient
}

func (b *BCacheCacheSet) client(systemID string, id int) *APIClient {
	return b.APIClient.
		SubClient("nodes").SubClient(systemID).
		SubClient("bcache-cache-set").SubClient(fmt.Sprintf("%v", id))
}

// Get BCacheCacheSet details.
func (b *BCacheCacheSet) Get(ctx context.Context, systemID string, id int) (*entity.BCacheCacheSet, error) {
	bCacheCacheSet := new(entity.BCacheCacheSet)
	err := b.client(systemID, id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, bCacheCacheSet)
	})

	return bCacheCacheSet, err
}

// Update BCacheCacheSet.
func (b *BCacheCacheSet) Update(ctx context.Context, systemID string, id int, params *entity.BCacheCacheSetParams) (*entity.BCacheCacheSet, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	bCacheCacheSet := new(entity.BCacheCacheSet)
	err = b.client(systemID, id).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, bCacheCacheSet)
	})

	return bCacheCacheSet, err
}

// Delete BCacheCacheSet.
func (b *BCacheCacheSet) Delete(ctx context.Context, systemID string, id int) error {
	return b.client(systemID, id).Delete(ctx)
}
