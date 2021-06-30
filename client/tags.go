package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Tags struct {
	ApiClient ApiClient
}

func (t *Tags) client() ApiClient {
	return t.ApiClient.GetSubObject("tags")
}

func (t *Tags) Get() (tags []entity.Tag, err error) {
	err = t.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &tags)
	})
	return
}

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
