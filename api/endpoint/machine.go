package endpoint

import (
	"net"

	"github.com/ionutbalutoiu/gomaasclient/api/endpoint/node"
)

// Machine represents the MaaS Machine endpoint.
type Machine struct {
	BootInterface   NetworkInterface `json:"boot_interface,omitempty"`
	Pod             Pod              `json:"pod,omitempty"`
	BootDisk        BlockDevice      `json:"boot_disk,omitempty"`
	Domain          Domain           `json:"domain,omitempty"`
	DefaultGateways struct {
		IPv4 struct {
			GatewayIP net.IP `json:"gateway_ip,omitempty"`
			LinkID    int    `json:"link_id,omitempty"`
		} `json:"ipv4,omitempty"`
		IPv6 struct {
			GatewayIP net.IP `json:"gateway_ip,omitempty"`
			LinkID    int    `json:"link_id,omitempty"`
		} `json:"ipv6,omitempty"`
	} `json:"default_gateways,omitempty"`
	Pool                         ResourcePool        `json:"pool,omitempty"`
	Zone                         Zone                `json:"zone,omitempty"`
	TagNames                     []string            `json:"tag_names,omitempty"`
	IPAddresses                  []net.IP            `json:"ip_addresses,omitempty"`
	BlockDeviceSet               []BlockDevice       `json:"blockdevice_set,omitempty"`
	CacheSets                    []string            `json:"cache_sets,omitempty"`
	VolumeGroups                 []string            `json:"volume_groups,omitempty"`
	InterfaceSet                 []NetworkInterface  `json:"interface_set,omitempty"`
	BCaches                      []string            `json:"bcaches,omitempty"`
	RAIDs                        []string            `json:"raids,omitempty"`
	SpecialFilesystems           []string            `json:"special_filesystems,omitempty"`
	ServiceSet                   []MachineServiceSet `json:"service_set,omitempty"`
	PhysicalBlockDeviceSet       []BlockDevice       `json:"physicalblockdevice_set,omitempty"`
	ISCSIBlockDeviceSet          []BlockDevice       `json:"iscsiblockdevice_set,omitempty"`
	VirtualBlockDeviceSet        []BlockDevice       `json:"virtualblockdevice_set,omitempty"`
	FQDN                         string              `json:"fqdn,omitempty"`
	DistroSeries                 string              `json:"distro_series,omitempty"`
	MinHWEKernel                 string              `json:"min_hwe_kernel,omitempty"`
	CommissioningStatusName      string              `json:"commissioning_status_name,omitempty"`
	SystemID                     string              `json:"system_id,omitempty"`
	NodeTypeName                 string              `json:"node_type_name,omitempty"`
	StorageTestStatusName        string              `json:"storage_test_status_name,omitempty"`
	Owner                        string              `json:"owner,omitempty"`
	HWEKernel                    string              `json:"hwe_kernel,omitempty"`
	TestingStatusName            string              `json:"testing_status_name,omitempty"`
	Version                      string              `json:"version,omitempty"`
	Architecture                 string              `json:"architecture,omitempty"`
	PowerState                   string              `json:"power_state,omitempty"`
	MemoryTestStatusName         string              `json:"memory_test_status_name,omitempty"`
	PowerType                    string              `json:"power_type,omitempty"`
	OwnerData                    interface{}         `json:"owner_data,omitempty"`
	Hostname                     string              `json:"hostname,omitempty"`
	Description                  string              `json:"description,omitempty"`
	StatusAction                 string              `json:"status_action,omitempty"`
	StatusMessage                string              `json:"status_message,omitempty"`
	StatusName                   string              `json:"status_name,omitempty"`
	OSystem                      string              `json:"osystem,omitempty"`
	CPUTestStatusName            string              `json:"cpu_test_status_name,omitempty"`
	OtherTestStatusName          string              `json:"other_test_status_name,omitempty"`
	ResourceURI                  string              `json:"resource_uri,omitempty"`
	Memory                       int                 `json:"memory,omitempty"`
	NodeType                     int                 `json:"node_type,omitempty"`
	CurrentCommissioningResultID int                 `json:"current_commissioning_result_id,omitempty"`
	CPUTestStatus                int                 `json:"cpu_test_status,omitempty"`
	AddressTTL                   int                 `json:"address_ttl,omitempty"`
	Storage                      float64             `json:"storage,omitempty"`
	HardwareInfo                 map[string]string   `json:"hardware_info,omitempty"`
	CPUCount                     int                 `json:"cpu_count,omitempty"`
	Status                       node.Status         `json:"status,omitempty"`
	CurrentInstallationResultID  int                 `json:"current_installation_result_id,omitempty"`
	CurrentTestingResultID       int                 `json:"current_testing_result_id,omitempty"`
	CommissioningStatus          int                 `json:"commissioning_status,omitempty"`
	OtherTestStatus              int                 `json:"other_test_status,omitempty"`
	TestingStatus                int                 `json:"testing_status,omitempty"`
	StorageTestStatus            int                 `json:"storage_test_status,omitempty"`
	SwapSize                     int                 `json:"swap_size,omitempty"`
	MemoryTestStatus             int                 `json:"memory_test_status,omitempty"`
	CPUSpeed                     int                 `json:"cpu_speed,omitempty"`
	DisableIPv4                  bool                `json:"disable_ipv4,omitempty"`
	Netboot                      bool                `json:"netboot,omitempty"`
	Locked                       bool                `json:"locked,omitempty"`
}

// MachineServiceSet represents a Machine's "service_set".
// This type should not be used directly.
type MachineServiceSet struct {
	Name       string `json:"name,omitempty"`
	Status     string `json:"status,omitempty"`
	StatusInfo string `json:"status_info,omitempty"`
}

// MachineParams enumerates the parameters for the machine update operation
type MachineParams struct {
	CPUCount      int    `json:"cpu_count"`
	Memory        int    `json:"memory"`
	SwapSize      int    `json:"swap_size"`
	PXEMacAddress string `json:"mac_addresses"`
	Architecture  string `json:"architecture"`
	MinHWEKernel  string `json:"min_hwe_kernel"`
	PowerType     string `json:"power_type"`
	Hostname      string `json:"hostname"`
	Description   string `json:"description"`
	Domain        string `json:"domain"`
	Pool          string `json:"pool"`
	Zone          string `json:"zone"`
	Commission    bool   `json:"commission"`
}

// VirshPowerParams enumerates the parameters for the virsh power type.
type VirshPowerParams struct {
	PowerAddress  string `json:"power_parameters_power_address"`
	PowerID       string `json:"power_parameters_power_id"`
	PowerPassword string `json:"power_parameters_power_pass"`
}

// MachineCommissionParams enumerates the parameters for the commission operation
type MachineCommissionParams struct {
	EnableSSH            int    `json:"enable_ssh"`
	SkipBMCConfig        int    `json:"skip_bmc_config"`
	SkipNetworking       int    `json:"skip_networking"`
	SkipStorage          int    `json:"skip_storage"`
	CommissioningScripts string `json:"commissioning_scripts"`
	TestingScripts       string `json:"testing_scripts"`
}

// MachineDeployParams enumerates the parameters for the deploy operation
type MachineDeployParams struct {
	UserData     string `json:"user_data"`
	DistroSeries string `json:"distro_series"`
	HWEKernel    string `json:"hwe_kernel"`
	AgentName    string `json:"agent_name"`
	Comment      string `json:"comment"`
	BridgeFD     int    `json:"bridge_fd"`
	BridgeAll    bool   `json:"bridge_all"`
	BridgeSTP    bool   `json:"bridge_stp"`
	InstallRackD bool   `json:"install_rackd"`
	InstallKVM   bool   `json:"install_kvm"`
}
