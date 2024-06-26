package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
	gomaasapi "github.com/juju/gomaasapi/v2"
)

// NodeScript implements api.NodeScript
type NodeScript struct {
	APIClient APIClient
}

func (ns *NodeScript) client(name string) (*APIClient, error) {
	uri := ns.APIClient.URI()
	newURL := url.URL{Path: fmt.Sprintf("scripts/%s", name)}
	resURL := uri.ResolveReference(&newURL)
	input := map[string]interface{}{"resource_uri": resURL.String()}

	jsonObj, err := gomaasapi.JSONObjectFromStruct(ns.APIClient.AuthClient, input)
	if err != nil {
		return nil, err
	}

	maasObj, err := jsonObj.GetMAASObject()
	if err != nil {
		return nil, err
	}

	return &APIClient{ns.APIClient.AuthClient, &maasObj}, nil
}

// Get fetches a nodeScript with a given system_id
func (ns *NodeScript) Get(name string, includeScript bool) (*entity.NodeScript, error) {
	client, err := ns.client(name)
	if err != nil {
		return nil, err
	}

	nodeScript := new(entity.NodeScript)
	qsp := url.Values{}
	qsp.Set("include_script", strconv.FormatBool(includeScript))
	err = client.Get("", qsp, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}

// Update updates a given NodeScript
func (ns *NodeScript) Update(name string, nodeScriptParams *entity.NodeScriptParams, script []byte) (*entity.NodeScript, error) {
	client, err := ns.client(name)
	if err != nil {
		return nil, err
	}

	qsp, err := query.Values(nodeScriptParams)
	if err != nil {
		return nil, err
	}

	files := map[string][]byte{}
	if script != nil {
		files["script"] = script
	}

	nodeScript := new(entity.NodeScript)
	err = client.PutFiles(qsp, files, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}

// Delete deletes a given NodeScript
func (ns *NodeScript) Delete(name string) error {
	client, err := ns.client(name)
	if err != nil {
		return err
	}

	return client.Delete()
}

// AddTag adds a tag to a given NodeScript
func (ns *NodeScript) AddTag(name string, tag string) (*entity.NodeScript, error) {
	client, err := ns.client(name)
	if err != nil {
		return nil, err
	}

	nodeScript := new(entity.NodeScript)
	qsp := url.Values{}
	qsp.Set("tag", tag)
	err = client.Post("add_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}

// RemoveTag removes a tag from a given NodeScript
func (ns *NodeScript) RemoveTag(name string, tag string) (*entity.NodeScript, error) {
	client, err := ns.client(name)
	if err != nil {
		return nil, err
	}

	nodeScript := new(entity.NodeScript)
	qsp := url.Values{}
	qsp.Set("tag", tag)
	err = client.Post("remove_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}

// Download given NodeScript
func (ns *NodeScript) Download(name string, revision int) ([]byte, error) {
	client, err := ns.client(name)
	if err != nil {
		return nil, err
	}

	qsp := url.Values{}
	qsp.Set("revision", strconv.Itoa(revision))

	value := new([]byte)
	err = client.Get("download", qsp, func(data []byte) error {
		*value = data
		return nil
	})

	return *value, err
}

// Revert a version of a given NodeScript
func (ns *NodeScript) Revert(name string, to int) (*entity.NodeScript, error) {
	client, err := ns.client(name)
	if err != nil {
		return nil, err
	}

	nodeScript := new(entity.NodeScript)
	qsp := url.Values{}
	qsp.Set("to", strconv.Itoa(to))
	err = client.Post("revert", qsp, func(data []byte) error {
		return json.Unmarshal(data, nodeScript)
	})

	return nodeScript, err
}
