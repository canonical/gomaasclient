package endpoint

// Machines represents the Machines endpoint
type Machines []Machine

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
