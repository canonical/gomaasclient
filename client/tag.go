package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Tag implements api.Tag
type Tag struct {
	APIClient APIClient
}

func (t *Tag) client(name string) *APIClient {
	return t.APIClient.SubClient("tags").SubClient(name)
}

// Get fetches a given tag by name
func (t *Tag) Get(ctx context.Context, name string) (*entity.Tag, error) {
	tag := new(entity.Tag)
	err := t.client(name).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, tag)
	})

	return tag, err
}

// Update updates a given tag
func (t *Tag) Update(ctx context.Context, name string, tagParams *entity.TagParams) (*entity.Tag, error) {
	qsp, err := query.Values(tagParams)
	if err != nil {
		return nil, err
	}

	tag := new(entity.Tag)
	err = t.client(name).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, tag)
	})

	return tag, err
}

// Delete deletes a given tag
func (t *Tag) Delete(ctx context.Context, name string) error {
	return t.client(name).Delete(ctx)
}

// GetMachines fetches a list of machines with a given tag
func (t *Tag) GetMachines(ctx context.Context, name string) ([]entity.Machine, error) {
	machines := make([]entity.Machine, 0)
	err := t.client(name).Get(ctx, "machines", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &machines)
	})

	return machines, err
}

// AddMachines adds a set of Machines to a given tag
func (t *Tag) AddMachines(ctx context.Context, name string, machineIds []string) error {
	return t.updateNodes(ctx, name, machineIds, "add")
}

// RemoveMachines removes a set of Machines
func (t *Tag) RemoveMachines(ctx context.Context, name string, machineIds []string) error {
	return t.updateNodes(ctx, name, machineIds, "remove")
}

func (t *Tag) updateNodes(ctx context.Context, name string, machineIds []string, op string) error {
	qsp := url.Values{}
	for _, id := range machineIds {
		qsp.Add(op, id)
	}

	return t.client(name).Post(ctx, "update_nodes", qsp, func(data []byte) error { return nil })
}
