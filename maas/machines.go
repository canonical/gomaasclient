package maas

import (
	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
)

// MachinesFetcher is the interface that API Clients must implement
type MachinesFetcher interface {
	Allocate(params *endpoint.MachinesAllocateParams) ([]byte, error)
	Release(systemID []string, comment string) error
}

// NewMachineManager creates a new MachinesManager
func NewMachinesManager(client MachinesFetcher) *MachinesManager {
	return &MachinesManager{client: client}
}

// MachinesManager provides locking and management capabilities for Machines
type MachinesManager struct {
	client MachinesFetcher
}

// Allocate calls the allocate operation
func (m *MachinesManager) Allocate(params *endpoint.MachinesAllocateParams) (ma *endpoint.Machine, err error) {
	var res []byte
	res, err = m.client.Allocate(params)
	if err == nil {
		ma, err = NewMachine(res)
	}
	return
}

// Release calls the release operation.
func (m *MachinesManager) Release(systemIDs []string, comment string) error {
	return m.client.Release(systemIDs, comment)
}
