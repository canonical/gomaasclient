package api

import "github.com/maas/gomaasclient/entity"

// NodeDevices is an interface for listing Node Devices objects
type NodeDevices interface {
	Get(systemID string, param *entity.NodeDeviceParams) ([]entity.NodeDevice, error)
}
