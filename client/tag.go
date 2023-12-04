package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Tag implements api.Tag
type Tag struct {
	ApiClient ApiClient
}

func (t *Tag) client(name string) ApiClient {
	return t.ApiClient.GetSubObject("tags").GetSubObject(name)
}

// Get fetches a given tag by name
func (t *Tag) Get(name string) (tag *entity.Tag, err error) {
	tag = new(entity.Tag)
	err = t.client(name).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, tag)
	})
	return
}

// Update updates a given tag
func (t *Tag) Update(name string, tagParams *entity.TagParams) (tag *entity.Tag, err error) {
	qsp, err := query.Values(tagParams)
	if err != nil {
		return
	}
	tag = new(entity.Tag)
	err = t.client(name).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, tag)
	})
	return
}

// Delete deletes a given tag
func (t *Tag) Delete(name string) error {
	return t.client(name).Delete()
}

// GetMachines fetches a list of machines with a given tag
func (t *Tag) GetMachines(name string) (machines []entity.Machine, err error) {
	err = t.client(name).Get("machines", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &machines)
	})
	return
}

// AddMachines adds a set of Machines to a given tag
func (t *Tag) AddMachines(name string, machineIds []string) error {
	return t.updateNodes(name, machineIds, "add")
}

// RemoveMachines removes a set of Machines
func (t *Tag) RemoveMachines(name string, machineIds []string) error {
	return t.updateNodes(name, machineIds, "remove")
}

func (t *Tag) updateNodes(name string, machineIds []string, op string) error {
	qsp := url.Values{}
	for _, id := range machineIds {
		qsp.Add(op, id)
	}
	return t.client(name).Post("update_nodes", qsp, func(data []byte) error { return nil })
}
