package api

import (
	"github.com/ionutbalutoiu/gomaasclient/api/params"
	"github.com/ionutbalutoiu/gomaasclient/maas/entity"
)

// NetworkInterface represents the MaaS Server Interface endpoint
type NetworkInterface interface {
	Delete(systemID string, id int) error
	Get(systemID string, id int) (*entity.NetworkInterface, error)
	AddTag(systemID string, id int, tag string) (*entity.NetworkInterface, error)
	Disconnect(systemID string, id int) (*entity.NetworkInterface, error)
	LinkSubnet(systemID string, id int, params *params.NetworkInterfaceLink) (*entity.NetworkInterface, error)
	RemoveTag(systemID string, id int, tag string) (*entity.NetworkInterface, error)
	SetDefaultGateway(systemID string, id, linkID int) (*entity.NetworkInterface, error)
	UnlinkSubnet(systemID string, id, linkID int) (*entity.NetworkInterface, error)
	Put(systemID string, id int, params interface{}) (*entity.NetworkInterface, error)
}
