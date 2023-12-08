package entity

// VMHost represents the MAAS VM host endpoint.
type VMHost struct {
	Zone      Zone           `json:"zone,omitempty"`
	Pool      ResourcePool   `json:"pool,omitempty"`
	Used      VMHostResource `json:"used,omitempty"`
	Available VMHostResource `json:"available,omitempty"`
	Total     VMHostResource `json:"total,omitempty"`
	Host      struct {
		SystemID   string `json:"system_id,omitempty"`
		Incomplete bool   `json:"__incomplete__,omitempty"` //nolint:tagliatelle // MAAS returns this field intentionally
	} `json:"host,omitempty"`
	StoragePools          []VMHostStoragePool `json:"storage_pools,omitempty"`
	Architectures         []string            `json:"architectures,omitempty"`
	Tags                  []string            `json:"tags,omitempty"`
	Capabilities          []string            `json:"capabilities,omitempty"`
	Name                  string              `json:"name,omitempty"`
	Type                  string              `json:"type,omitempty"`
	DefaultMACVLANMode    string              `json:"default_macvlan_mode,omitempty"`
	ResourceURI           string              `json:"resource_uri,omitempty"`
	CPUOverCommitRatio    float64             `json:"cpu_over_commit_ratio,omitempty"`
	MemoryOverCommitRatio float64             `json:"memory_over_commit_ratio,omitempty"`
	ID                    int                 `json:"id,omitempty"`
}

// VMHostStoragePool represents the "storage_pools" object in a VMHost.
// This type should not be used directly.
type VMHostStoragePool struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	Path      string `json:"path,omitempty"`
	Total     int64  `json:"total,omitempty"`
	Used      int64  `json:"used,omitempty"`
	Available int64  `json:"available,omitempty"`
	Default   bool   `json:"default,omitempty"`
}

// VMHostResource represents the "used", "available", and "total" objects in a VMHost
// This type should not be used directly.
type VMHostResource struct {
	Cores        int   `json:"cores,omitempty"`
	Memory       int64 `json:"memory,omitempty"`
	LocalStorage int64 `json:"local_storage,omitempty"`
}

// VMHostParams enumerates the VMHost configuration options.
type VMHostParams struct {
	Type                  string  `url:"type,omitempty"`
	PowerAddress          string  `url:"power_address,omitempty"`
	PowerUser             string  `url:"power_user,omitempty"`
	PowerPass             string  `url:"power_pass,omitempty"`
	Name                  string  `url:"name,omitempty"`
	Zone                  string  `url:"zone,omitempty"`
	Pool                  string  `url:"pool,omitempty"`
	DefaultStoragePool    string  `url:"default_storage_pool,omitempty"`
	DefaultMacvlanMode    string  `url:"default_macvlan_mode,omitempty"`
	CPUOverCommitRatio    float64 `url:"cpu_over_commit_ratio,omitempty"`
	MemoryOverCommitRatio float64 `url:"memory_over_commit_ratio,omitempty"`
	Tags                  string  `url:"tags,omitempty"`
}

// VMHostMachineParams enumerates the VMHost machine configuration options.
type VMHostMachineParams struct {
	Cores           int    `url:"cores,omitempty"`
	PinnedCores     int    `url:"pinned_cores,omitempty"`
	Memory          int64  `url:"memory,omitempty"`
	HugepagesBacked bool   `url:"hugepages_backed,omitempty"`
	Architecture    string `url:"architecture,omitempty"`
	Storage         string `url:"storage,omitempty"`
	Interfaces      string `url:"interfaces,omitempty"`
	Hostname        string `url:"hostname,omitempty"`
}
