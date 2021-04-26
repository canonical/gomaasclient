package maas

import (
	"encoding/json"
	"sync"

	"github.com/ionutbalutoiu/gomaasclient/maas/entity"
)

// NewMachine converts a MAAS API JSON response into a Golang representation
func NewMachine(data []byte) (m *entity.Machine, err error) {
	m = &entity.Machine{}
	err = json.Unmarshal(data, m)
	return
}

// MachineManager contains functionality for manipulating the Machine endpoint.
// A MachineManager represents a single machine, as does the API endpoint.
type MachineManager struct {
	state  []entity.Machine
	client MachineFetcher
	mutex  sync.RWMutex
}

// NewMachineManager creates a new MachineManager.
// It attempts to fetch the current state of the machine with the given systemID,
// and will return the API client's error if it is not successful. It will also return
// an error from NewMachine if the response cannot successfully be parsed as a Machine.
func NewMachineManager(systemID string, client MachineFetcher) (*MachineManager, error) {
	res, err := client.Get(systemID)
	if err != nil {
		return nil, err
	}
	m, err := NewMachine(res)
	if err != nil {
		return nil, err
	}
	return &MachineManager{
		state:  []entity.Machine{*m},
		client: client,
		mutex:  sync.RWMutex{},
	}, nil
}

// Current returns the most current known state of the machine.
func (m *MachineManager) Current() *entity.Machine {
	return &m.state[len(m.state)-1]
}

func (m *MachineManager) append(current *entity.Machine) {
	m.state = append(m.state, *current)
}

func (m *MachineManager) appendBytes(data []byte) error {
	next, err := NewMachine(data)
	if err == nil {
		m.append(next)
	}
	return err
}

// SystemID returns the machine's systemID.
func (m *MachineManager) SystemID() string {
	return m.Current().SystemID
}

// Commission calls the commission operation on the API.
func (m *MachineManager) Commission(params MachineCommissionParams) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	res, err := m.client.Commission(m.SystemID(), params)
	if err == nil {
		err = m.appendBytes(res)
	}
	return err
}

// Deploy calls the deploy operation on the API.
func (m *MachineManager) Deploy(params *MachineDeployParams) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	res, err := m.client.Deploy(m.SystemID(), params)
	if err == nil {
		err = m.appendBytes(res)
	}
	return err
}

// Lock calls the lock operation on the API.
func (m *MachineManager) Lock(comment string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	res, err := m.client.Lock(m.SystemID(), comment)
	if err == nil {
		err = m.appendBytes(res)
	}
	return err
}

// Update fetches and returns the current state of the machine.
func (m *MachineManager) Update() (ma *entity.Machine, err error) {
	ma, err = m.update()
	if err == nil {
		m.append(ma)
	}
	return
}

func (m *MachineManager) update() (ma *entity.Machine, err error) {
	var res []byte
	res, err = m.client.Get(m.SystemID())
	if err == nil {
		ma, err = NewMachine(res)
	}
	return
}

// MachineFetcher is the interface that API clients must implement.
type MachineFetcher interface {
	Get(string) ([]byte, error)
	Commission(string, MachineCommissionParams) ([]byte, error)
	Deploy(string, *MachineDeployParams) ([]byte, error)
	Lock(string, string) ([]byte, error)
}

// MachineCommissionParams enumerates the parameters for the commission operation
type MachineCommissionParams struct {
	EnableSSH            int
	SkipBMCConfig        int
	SkipNetworking       int
	SkipStorage          int
	CommissioningScripts string
	TestingScripts       string
}

// MachineDeployParams enumerates the parameters for the deploy operation
type MachineDeployParams struct {
	UserData     string
	DistroSeries string
	HWEKernel    string
	AgentName    string
	Comment      string
	BridgeFD     int
	BridgeAll    bool
	BridgeSTP    bool
	InstallRackD bool `json:"install_rackd"`
	InstallKVM   bool `json:"install_kvm"`
}
