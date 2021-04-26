package maas

import (
	"github.com/ionutbalutoiu/gomaasclient/maas/entity"
)

// Machines represents the Machines endpoint
type Machines []entity.Machine

// MachinesManager provides locking and management capabilities for Machines
type MachinesManager struct {
	client MachinesFetcher
}

// NewMachineManager creates a new MachinesManager
func NewMachinesManager(client MachinesFetcher) *MachinesManager {
	return &MachinesManager{client: client}
}

// Allocate calls the allocate operation
func (m *MachinesManager) Allocate(params *MachinesAllocateParams) (ma *entity.Machine, err error) {
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

// MachinesParams enumerates the options for the GET operation
type MachinesParams struct {
	Hostname   []string
	MACAddress []string
	ID         []string
	Domain     string
	Zone       string
	AgentName  string
	Pool       []string
}

// MachinesAllocateParams enumerates the options for the allocate operation.
type MachinesAllocateParams struct {
	Tags             []string `json:"tags"`
	NotTags          []string `json:"not_tags"`
	NotInZone        []string `json:"not_in_zone"`
	NotInPool        []string `json:"not_in_pool"`
	Subnets          []string `json:"subnets"`
	NotSubnets       []string `json:"not_subnets"`
	Storage          []string `json:"storage"`
	Fabrics          []string `json:"fabrics"`
	NotFabrics       []string `json:"not_fabrics"`
	FabricClasses    []string `json:"fabric_classes"`
	NotFabricClasses []string `json:"not_fabric_classes"`
	Name             string   `json:"name"`
	SystemID         string   `json:"system_id"`
	Arch             string   `json:"arch"`
	Zone             string   `json:"zone"`
	Pool             string   `json:"pool"`
	Pod              string   `json:"pod"`
	NotPod           string   `json:"not_pod"`
	PodType          string   `json:"pod_type"`
	NotPodType       string   `json:"not_pod_type"`
	Interfaces       string   `json:"interfaces"`
	AgentName        string   `json:"agent_name"`
	Comment          string   `json:"comment"`
	CPUCount         int      `json:"cpu_count"`
	Mem              int      `json:"mem"`
	BridgeFD         int      `json:"bridge_fd"`
	BridgeAll        bool     `json:"bridge_all"`
	BridgeSTP        bool     `json:"bridge_stp"`
	DryRun           bool     `json:"dry_run"`
	Verbose          bool     `json:"verbose"`
}

// MachinesFetcher is the interface that API Clients must implement
type MachinesFetcher interface {
	Allocate(params *MachinesAllocateParams) ([]byte, error)
	Release(systemID []string, comment string) error
}
