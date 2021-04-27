package maas

import (
	"encoding/json"
	"sync"

	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
)

// NewPod converts a MAAS API JSON response into a Golang representation
func NewPod(data []byte) (m *endpoint.Pod, err error) {
	m = &endpoint.Pod{}
	err = json.Unmarshal(data, m)
	return
}

// PodFetcher is the interface that API clients must implement.
type PodFetcher interface {
	Get(int) ([]byte, error)
	Update(int, *endpoint.PodParams) ([]byte, error)
	Delete(int) error
	Compose(int, *endpoint.PodMachineParams) ([]byte, error)
	Refresh(int) ([]byte, error)
}

// NewPodManager creates a new PodManager
func NewPodManager(id int, client PodFetcher) (*PodManager, error) {
	res, err := client.Get(id)
	if err != nil {
		return nil, err
	}
	pod, err := NewPod(res)
	if err != nil {
		return nil, err
	}
	return &PodManager{
		state:  []endpoint.Pod{*pod},
		client: client,
		mutex:  sync.RWMutex{},
	}, nil
}

// PodManager contains functionality for manipulating the Pod endpoint.
// A PodManager represents a single pod, as does the API endpoint.
type PodManager struct {
	state  []endpoint.Pod
	client PodFetcher
	mutex  sync.RWMutex
}

// Current returns the most current known state of the pod.
func (m *PodManager) Current() *endpoint.Pod {
	return &m.state[len(m.state)-1]
}

// ID returns the pod's ID.
func (m *PodManager) ID() int {
	return m.Current().ID
}

// Calls the update Pod API
func (m *PodManager) Update(params *endpoint.PodParams) (*endpoint.Pod, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	res, err := m.client.Update(m.ID(), params)
	if err != nil {
		return nil, err
	}

	pod, err := NewPod(res)
	if err != nil {
		return nil, err
	}

	m.state = append(m.state, *pod)
	return pod, nil
}

// Calls the delete Pod API
func (m *PodManager) Delete() error {
	return m.client.Delete(m.ID())
}

// Compose calls the compose operation on the API.
func (m *PodManager) Compose(params *endpoint.PodMachineParams) (*endpoint.Machine, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	res, err := m.client.Compose(m.ID(), params)
	if err != nil {
		return nil, err
	}

	ma, err := NewMachine(res)
	if err != nil {
		return nil, err
	}

	return ma, err
}

// Refresh calls the refresh operation on the API.
func (m *PodManager) Refresh() (*endpoint.Pod, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	res, err := m.client.Refresh(m.ID())
	if err != nil {
		return nil, err
	}

	pod, err := NewPod(res)
	if err != nil {
		return nil, err
	}

	m.state = append(m.state, *pod)
	return pod, nil
}
