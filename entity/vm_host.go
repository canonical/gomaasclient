package entity

// VMHost represents the MAAS VM host endpoint.
type VMHost struct {
	Zone               Zone                `json:"zone,omitempty"`
	Pool               ResourcePool        `json:"pool,omitempty"`
	ResourceURI        string              `json:"resource_uri,omitempty"`
	DefaultMACVLANMode string              `json:"default_macvlan_mode,omitempty"`
	Type               string              `json:"type,omitempty"`
	Name               string              `json:"name,omitempty"`
	Tags               []string            `json:"tags,omitempty"`
	Architectures      []string            `json:"architectures,omitempty"`
	StoragePools       []VMHostStoragePool `json:"storage_pools,omitempty"`
	Capabilities       []string            `json:"capabilities,omitempty"`
	Host               struct {
		SystemID   string `json:"system_id,omitempty"`
		Incomplete bool   `json:"__incomplete__,omitempty"` //nolint:tagliatelle // MAAS returns this field intentionally
	} `json:"host,omitempty"`
	Total                 VMHostResource `json:"total,omitempty"`
	Available             VMHostResource `json:"available,omitempty"`
	Used                  VMHostResource `json:"used,omitempty"`
	CPUOverCommitRatio    float64        `json:"cpu_over_commit_ratio,omitempty"`
	MemoryOverCommitRatio float64        `json:"memory_over_commit_ratio,omitempty"`
	ID                    int            `json:"id,omitempty"`
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
	Pool                  string  `url:"pool,omitempty"`
	Type                  string  `url:"type,omitempty"`
	PowerUser             string  `url:"power_user,omitempty"`
	PowerPass             string  `url:"power_pass,omitempty"`
	Name                  string  `url:"name,omitempty"`
	Zone                  string  `url:"zone,omitempty"`
	PowerAddress          string  `url:"power_address,omitempty"`
	DefaultStoragePool    string  `url:"default_storage_pool,omitempty"`
	Key                   string  `url:"key,omitempty"`
	Tags                  string  `url:"tags,omitempty"`
	DefaultMacvlanMode    string  `url:"default_macvlan_mode,omitempty"`
	Certificate           string  `url:"certificate,omitempty"`
	MemoryOverCommitRatio float64 `url:"memory_over_commit_ratio,omitempty"`
	CPUOverCommitRatio    float64 `url:"cpu_over_commit_ratio,omitempty"`
}

// VMHostMachineParams enumerates the VMHost machine configuration options.
type VMHostMachineParams struct {
	Architecture    string `url:"architecture,omitempty"`
	Storage         string `url:"storage,omitempty"`
	Interfaces      string `url:"interfaces,omitempty"`
	Hostname        string `url:"hostname,omitempty"`
	Cores           int    `url:"cores,omitempty"`
	PinnedCores     int    `url:"pinned_cores,omitempty"`
	Memory          int64  `url:"memory,omitempty"`
	HugepagesBacked bool   `url:"hugepages_backed,omitempty"`
}
