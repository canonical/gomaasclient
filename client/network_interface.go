package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type NetworkInterface struct {
	ApiClient ApiClient
}

func (n *NetworkInterface) client(systemID string, id int) ApiClient {
	return n.ApiClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("interfaces").GetSubObject(fmt.Sprintf("%v", id))
}

func (n *NetworkInterface) callPost(systemID string, id int, qsp url.Values, op string) (networkInterface *entity.NetworkInterface, err error) {
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID, id).Post(op, qsp, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}

func (n *NetworkInterface) Get(systemID string, id int) (networkInterface *entity.NetworkInterface, err error) {
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID, id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}

func (n *NetworkInterface) Update(systemID string, id int, params interface{}) (networkInterface *entity.NetworkInterface, err error) {
	networkInterface = new(entity.NetworkInterface)
	err = n.client(systemID, id).Put(ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, networkInterface)
	})
	return
}

func (n *NetworkInterface) Delete(systemID string, id int) error {
	return n.client(systemID, id).Delete()
}

func (n *NetworkInterface) Disconnect(systemID string, id int) (networkInterface *entity.NetworkInterface, err error) {
	return n.callPost(systemID, id, url.Values{}, "disconnect")
}

func (n *NetworkInterface) AddTag(systemID string, id int, tag string) (*entity.NetworkInterface, error) {
	qsp := url.Values{}
	qsp.Add("tag", tag)
	return n.callPost(systemID, id, qsp, "add_tag")
}

func (n *NetworkInterface) RemoveTag(systemID string, id int, tag string) (*entity.NetworkInterface, error) {
	qsp := url.Values{}
	qsp.Add("tag", tag)
	return n.callPost(systemID, id, qsp, "remove_tag")
}

func (n *NetworkInterface) LinkSubnet(systemID string, id int, params *entity.NetworkInterfaceLinkParams) (*entity.NetworkInterface, error) {
	return n.callPost(systemID, id, ToQSP(params), "link_subnet")
}

func (n *NetworkInterface) UnlinkSubnet(systemID string, id int, linkID int) (networkInterface *entity.NetworkInterface, err error) {
	return n.callPost(systemID, id, url.Values{}, "unlink_subnet")
}

func (n *NetworkInterface) SetDefaultGateway(systemID string, id int, linkID int) (networkInterface *entity.NetworkInterface, err error) {
	qsp := url.Values{}
	qsp.Add("link_id", fmt.Sprintf("%v", linkID))
	return n.callPost(systemID, id, qsp, "set_default_gateway")
}
