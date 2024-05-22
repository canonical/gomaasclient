package entity

// Node represents the MAAS Node endpoint.
type Node Machine

// NUMANode represents the MAAS numa node
// referred from https://github.com/maas/maas/blob/deab73792a4fe839a2e84a926a6c728d510fc9ad/src/maasserver/api/machines.py#L110
type NUMANode struct {
	HugepagesSet []NUMANodeHugepages `json:"hugepages_set,omitempty"`
	Cores        []int               `json:"cores,omitempty"`
	Index        int                 `json:"index,omitempty"`
	Memory       int                 `json:"memory,omitempty"`
}

// NUMANodeHugepages represents the MAAS numa node hugepages
// referred from https://github.com/maas/maas/blob/deab73792a4fe839a2e84a926a6c728d510fc9ad/src/maasserver/api/machines.py#L108
type NUMANodeHugepages struct {
	PageSize int `json:"page_size,omitempty"`
	Total    int `json:"total,omitempty"`
}

// NodeDetails represent the MAAS node details
type NodeDetails struct {
	LLDP string `json:"lldp,omitempty"`
	LSHW string `json:"lshw,omitempty"`
}

// NodePowerState represent current node power state
type NodePowerState struct {
	State string `json:"state,omitempty"`
}

// NodeGetParams enumerates the parameters for the get nodes operation
type NodeGetParams struct {
	ID               []string `url:"id,omitempty"`
	NotID            []string `url:"not_id,omitempty"`
	Hostname         []string `url:"hostname,omitempty"`
	NotHostname      []string `url:"not_hostname,omitempty"`
	MACAddress       []string `url:"mac_address,omitempty"`
	Domain           []string `url:"domain,omitempty"`
	NotDomain        []string `url:"not_domain,omitempty"`
	AgentName        []string `url:"agent_name,omitempty"`
	NotAgentName     []string `url:"not_agent_name,omitempty"`
	Status           []string `url:"status,omitempty"`
	NotStatus        []string `url:"not_status,omitempty"`
	SimpleStatus     []string `url:"simple_status,omitempty"`
	NotSimpleStatus  []string `url:"not_simple_status,omitempty"`
	Arch             []string `url:"arch,omitempty"`
	NotArch          []string `url:"not_arch,omitempty"`
	Tags             []string `url:"tags,omitempty"`
	NotTags          []string `url:"not_tags,omitempty"`
	Fabrics          []string `url:"fabrics,omitempty"`
	NotFabrics       []string `url:"not_fabrics,omitempty"`
	FabricClasses    []string `url:"fabric_classes,omitempty"`
	NotFabricClasses []string `url:"not_fabric_classes,omitempty"`
	Subnets          []string `url:"subnets,omitempty"`
	NotSubnets       []string `url:"not_subnets,omitempty"`
	LinkSpeed        []int    `url:"link_speed,omitempty"`
	VLANs            []string `url:"vlans,omitempty"`
	NotVLANs         []string `url:"not_vlans,omitempty"`
	Zone             []string `url:"zone,omitempty"`
	NotInZone        []string `url:"not_in_zone,omitempty"`
	Pool             []string `url:"pool,omitempty"`
	NotInPool        []string `url:"not_in_pool,omitempty"`
	Storage          string   `url:"storage,omitempty"`
	Interfaces       string   `url:"devices,omitempty"`
	Devices          string   `url:"devices,omitempty"`
	CPUCount         []int    `url:"cpu_count,omitempty"`
	CPUSpeed         []int    `url:"cpu_speed,omitempty"`
	Mem              []int64  `url:"mem,omitempty"`
	Pod              []string `url:"pod,omitempty"`
	NotPod           []string `url:"not_pod,omitempty"`
	PodType          []string `url:"pod_type,omitempty"`
	NotPodType       []string `url:"not_pod_type,omitempty"`
	Owner            []string `url:"owner,omitempty"`
	NotOwner         []string `url:"not_owner,omitempty"`
	PowerState       []string `url:"power_state,omitempty"`
	NotPowerState    []string `url:"not_power_state,omitempty"`
}

// NodePowerOnParams enumerates the parameters for the node power on operation
// UserData should be Base64-encoded data
type NodePowerOnParams struct {
	Comment  string `url:"comment,omitempty"`
	UserData string `url:"user_data,omitempty"`
}

// NodePowerOffParams enumerates the parameters for the node power off operation
type NodePowerOffParams struct {
	Comment  string `url:"comment,omitempty"`
	StopMode string `url:"stop_mode,omitempty"`
}
