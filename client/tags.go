package client

import (
	"encoding/json"
	"net/url"

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
	tag = new(entity.Tag)
	err = t.client().Post("", ToQSP(tagParams), func(data []byte) error {
		return json.Unmarshal(data, tag)
	})
	return
}
