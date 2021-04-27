package maas

import (
	"encoding/json"

	"github.com/ionutbalutoiu/gomaasclient/api/endpoint"
)

// NewPods converts a MAAS API JSON response with Pods into the Golang representation
func NewPods(data []byte) (m []*endpoint.Pod, err error) {
	m = []*endpoint.Pod{}
	err = json.Unmarshal(data, &m)
	return
}

// PodsFetcher is the interface that lower level API Clients must implement
type PodsFetcher interface {
	Get() ([]byte, error)
	Create(params *endpoint.PodParams) ([]byte, error)
}

// NewPodsManager creates a new PodsManager
func NewPodsManager(client PodsFetcher) *PodsManager {
	return &PodsManager{client: client}
}

// PodsManager management capabilities for Pods
type PodsManager struct {
	client PodsFetcher
}

// Get a listing of all Pods.
func (m *PodsManager) Get() (pods []*endpoint.Pod, err error) {
	res, err := m.client.Get()
	if err != nil {
		return nil, err
	}
	pods, err = NewPods(res)
	return
}

// Create or discover a new Pod.
func (m *PodsManager) Create(params *endpoint.PodParams) (ma *endpoint.Pod, err error) {
	var res []byte
	res, err = m.client.Create(params)
	if err == nil {
		ma, err = NewPod(res)
	}
	return
}
