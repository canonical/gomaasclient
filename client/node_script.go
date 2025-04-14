package client

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// NodeScript implements api.NodeScript
type NodeScript struct {
	APIClient APIClient
}

func (ns *NodeScript) client(name string) *APIClient {
	return ns.APIClient.SubClient("scripts").SubClientWithoutSlash(name)
}

// Get fetches a nodeScript with a given system_id
func (ns *NodeScript) Get(ctx context.Context, name string, includeScript bool) (*entity.NodeScript, error) {
	nodeScript := new(entity.NodeScript)
	qsp := url.Values{}
	qsp.Set("include_script", strconv.FormatBool(includeScript))
	err := ns.client(name).Get(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}

// Update updates a given NodeScript
func (ns *NodeScript) Update(ctx context.Context, name string, nodeScriptParams *entity.NodeScriptParams, script []byte) (*entity.NodeScript, error) {
	qsp, err := query.Values(nodeScriptParams)
	if err != nil {
		return nil, err
	}

	files := map[string][]byte{}
	if script != nil {
		files["script"] = script
	}

	nodeScript := new(entity.NodeScript)
	err = ns.client(name).PutFiles(ctx, qsp, files, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}

// Delete deletes a given NodeScript
func (ns *NodeScript) Delete(ctx context.Context, name string) error {
	return ns.client(name).Delete(ctx)
}

// AddTag adds a tag to a given NodeScript
func (ns *NodeScript) AddTag(ctx context.Context, name string, tag string) (*entity.NodeScript, error) {
	nodeScript := new(entity.NodeScript)
	qsp := url.Values{}
	qsp.Set("tag", tag)
	err := ns.client(name).Post(ctx, "add_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}

// RemoveTag removes a tag from a given NodeScript
func (ns *NodeScript) RemoveTag(ctx context.Context, name string, tag string) (*entity.NodeScript, error) {
	nodeScript := new(entity.NodeScript)
	qsp := url.Values{}
	qsp.Set("tag", tag)
	err := ns.client(name).Post(ctx, "remove_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}

// Download given NodeScript
func (ns *NodeScript) Download(ctx context.Context, name string, revision int) ([]byte, error) {
	qsp := url.Values{}
	qsp.Set("revision", strconv.Itoa(revision))

	value := new([]byte)
	err := ns.client(name).Get(ctx, "download", qsp, func(data []byte) error {
		*value = data
		return nil
	})

	return *value, err
}

// Revert a version of a given NodeScript
func (ns *NodeScript) Revert(ctx context.Context, name string, to int) (*entity.NodeScript, error) {
	nodeScript := new(entity.NodeScript)
	qsp := url.Values{}
	qsp.Set("to", strconv.Itoa(to))
	err := ns.client(name).Post(ctx, "revert", qsp, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}
