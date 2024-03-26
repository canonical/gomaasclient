package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// BootResource implements api.BootResource
type BootResource struct {
	APIClient APIClient
}

func (b *BootResource) client(id int) APIClient {
	return b.APIClient.GetSubObject("boot-resources").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a boot resource with a given id
func (b *BootResource) Get(id int) (*entity.BootResource, error) {
	bootResource := new(entity.BootResource)
	err := b.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, bootResource)
	})

	return bootResource, err
}

// Delete deletes a given boot resource
func (b *BootResource) Delete(id int) error {
	return b.client(id).Delete()
}
