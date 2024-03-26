package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// NetworkInterface represents the MAAS Server Interface endpoint
type NetworkInterface interface {
	Get(systemID string, id int) (*entity.NetworkInterface, error)
	Update(systemID string, id int, params *entity.NetworkInterfaceUpdateParams) (*entity.NetworkInterface, error)
	Delete(systemID string, id int) error
	Disconnect(systemID string, id int) (*entity.NetworkInterface, error)
	AddTag(systemID string, id int, tag string) (*entity.NetworkInterface, error)
	RemoveTag(systemID string, id int, tag string) (*entity.NetworkInterface, error)
	LinkSubnet(systemID string, id int, params *entity.NetworkInterfaceLinkParams) (*entity.NetworkInterface, error)
	UnlinkSubnet(systemID string, id int, linkID int) (*entity.NetworkInterface, error)
	SetDefaultGateway(systemID string, id int, linkID int) (*entity.NetworkInterface, error)
}
