//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Tags implements api.Tags
type Tags struct {
	APIClient APIClient
}

func (t *Tags) client() APIClient {
	return t.APIClient.GetSubObject("tags")
}

// Get fetches a list of Tag objects
func (t *Tags) Get() ([]entity.Tag, error) {
	tags := make([]entity.Tag, 0)
	err := t.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &tags)
	})

	return tags, err
}

// Create creates a new Tag
func (t *Tags) Create(tagParams *entity.TagParams) (*entity.Tag, error) {
	qsp, err := query.Values(tagParams)
	if err != nil {
		return nil, err
	}

	tag := new(entity.Tag)
	err = t.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, tag)
	})

	return tag, err
}
