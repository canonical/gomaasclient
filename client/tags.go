package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Tags implements api.Tags
type Tags struct {
	APIClient APIClient
}

func (t *Tags) client() APIClient {
	return t.APIClient.GetSubObject("tags")
}

// Get fetches a list of Tag objects
func (t *Tags) Get() (tags []entity.Tag, err error) {
	err = t.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &tags)
	})

	return
}

// Create creates a new Tag
func (t *Tags) Create(tagParams *entity.TagParams) (tag *entity.Tag, err error) {
	qsp, err := query.Values(tagParams)
	if err != nil {
		return
	}

	tag = new(entity.Tag)
	err = t.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, tag)
	})

	return
}
