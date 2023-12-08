package entity

import (
	"encoding/json"
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/maas/gomaasclient/entity/node"
)

// Machine represents the MaaS Machine endpoint.
type Machine struct {
	BootInterface   NetworkInterface `json:"boot_interface,omitempty"`
	VMHost          VMHost           `json:"pod,omitempty"`
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
	Pool                         ResourcePool               `json:"pool,omitempty"`
	Zone                         Zone                       `json:"zone,omitempty"`
	TagNames                     []string                   `json:"tag_names,omitempty"`
	IPAddresses                  []net.IP                   `json:"ip_addresses,omitempty"`
	BlockDeviceSet               []BlockDevice              `json:"blockdevice_set,omitempty"`
	CacheSets                    []MachineCacheSet          `json:"cache_sets,omitempty"`
	VolumeGroups                 []MachineVolumeGroup       `json:"volume_groups,omitempty"`
	InterfaceSet                 []NetworkInterface         `json:"interface_set,omitempty"`
	BCaches                      []MachineBCache            `json:"bcaches,omitempty"`
	RAIDs                        []MachineRAID              `json:"raids,omitempty"`
	SpecialFilesystems           []MachineSpecialFilesystem `json:"special_filesystems,omitempty"`
	ServiceSet                   []MachineServiceSet        `json:"service_set,omitempty"`
	PhysicalBlockDeviceSet       []BlockDevice              `json:"physicalblockdevice_set,omitempty"`
	ISCSIBlockDeviceSet          []BlockDevice              `json:"iscsiblockdevice_set,omitempty"`
	VirtualBlockDeviceSet        []BlockDevice              `json:"virtualblockdevice_set,omitempty"`
	FQDN                         string                     `json:"fqdn,omitempty"`
	DistroSeries                 string                     `json:"distro_series,omitempty"`
	MinHWEKernel                 string                     `json:"min_hwe_kernel,omitempty"`
	CommissioningStatusName      string                     `json:"commissioning_status_name,omitempty"`
	SystemID                     string                     `json:"system_id,omitempty"`
	NodeTypeName                 string                     `json:"node_type_name,omitempty"`
	StorageTestStatusName        string                     `json:"storage_test_status_name,omitempty"`
	Owner                        string                     `json:"owner,omitempty"`
	HWEKernel                    string                     `json:"hwe_kernel,omitempty"`
	TestingStatusName            string                     `json:"testing_status_name,omitempty"`
	Version                      string                     `json:"version,omitempty"`
	Architecture                 string                     `json:"architecture,omitempty"`
	PowerState                   string                     `json:"power_state,omitempty"`
	MemoryTestStatusName         string                     `json:"memory_test_status_name,omitempty"`
	PowerType                    string                     `json:"power_type,omitempty"`
	OwnerData                    interface{}                `json:"owner_data,omitempty"`
	Hostname                     string                     `json:"hostname,omitempty"`
	EnableHwSync                 bool                       `json:"enable_hw_sync,omitempty"`
	HwLastSync                   MAASTime                   `json:"last_sync,omitempty"`
	HwSyncInterval               int                        `json:"sync_interval,omitempty"`
	HwNextSync                   MAASTime                   `json:"next_sync,omitempty"`
	Description                  string                     `json:"description,omitempty"`
	StatusAction                 string                     `json:"status_action,omitempty"`
	StatusMessage                string                     `json:"status_message,omitempty"`
	StatusName                   string                     `json:"status_name,omitempty"`
	OSystem                      string                     `json:"osystem,omitempty"`
	CPUTestStatusName            string                     `json:"cpu_test_status_name,omitempty"`
	OtherTestStatusName          string                     `json:"other_test_status_name,omitempty"`
	ResourceURI                  string                     `json:"resource_uri,omitempty"`
	Memory                       int64                      `json:"memory,omitempty"`
	NodeType                     int                        `json:"node_type,omitempty"`
	CurrentCommissioningResultID int                        `json:"current_commissioning_result_id,omitempty"`
	CPUTestStatus                int                        `json:"cpu_test_status,omitempty"`
	AddressTTL                   int                        `json:"address_ttl,omitempty"`
	Storage                      float64                    `json:"storage,omitempty"`
	HardwareInfo                 map[string]string          `json:"hardware_info,omitempty"`
	CPUCount                     int                        `json:"cpu_count,omitempty"`
	Status                       node.Status                `json:"status,omitempty"`
	CurrentInstallationResultID  int                        `json:"current_installation_result_id,omitempty"`
	CurrentTestingResultID       int                        `json:"current_testing_result_id,omitempty"`
	CommissioningStatus          int                        `json:"commissioning_status,omitempty"`
	OtherTestStatus              int                        `json:"other_test_status,omitempty"`
	TestingStatus                int                        `json:"testing_status,omitempty"`
	StorageTestStatus            int                        `json:"storage_test_status,omitempty"`
	SwapSize                     int64                      `json:"swap_size,omitempty"`
	MemoryTestStatus             int                        `json:"memory_test_status,omitempty"`
	CPUSpeed                     int                        `json:"cpu_speed,omitempty"`
	DisableIPv4                  bool                       `json:"disable_ipv4,omitempty"`
	Netboot                      bool                       `json:"netboot,omitempty"`
	Locked                       bool                       `json:"locked,omitempty"`
}

func (m *Machine) UnmarshalJSON(data []byte) error {
	type machine Machine

	if err := json.Unmarshal(data, (*machine)(m)); err != nil {
		return err
	}

	// NOTE: deduce field value for backward compatibility
	// `enable_hw_sync` field is exposed via API in MAAS 3.4+
	// but hardware sync feature was added in 3.2
	if !m.EnableHwSync {
		t := MAASTime{}
		if m.HwSyncInterval != 0 && (m.HwNextSync != t || m.HwLastSync != t) {
			m.EnableHwSync = true
		}
	}

	return nil
}

// MAASTime is a custom time format returned by MAAS API
// which is not RFC3339
type MAASTime time.Time

func (t *MAASTime) UnmarshalJSON(data []byte) error {
	str, err := strconv.Unquote(string(data))
	if errors.Is(err, strconv.ErrSyntax) {
		return nil
	} else if err != nil {
		return err
	}

	temp, err := time.Parse("2006-01-02T15:04:05.999", str)
	if err != nil {
		return err
	}

	*t = MAASTime(temp)

	return nil
}

// String returns MAASTime in RFC3339Nano format
func (t MAASTime) String() string {
	temp := time.Time(t)
	return temp.Format(time.RFC3339Nano)
}

// MachineCacheSet represents a Machine's "cache_sets" list item.
// This type should not be used directly.
type MachineCacheSet struct {
	SystemID   string `json:"system_id,omitempty"`
	ID         int    `json:"id,omitempty"`
	Incomplete bool   `json:"__incomplete__,omitempty"`
}

// MachineVolumeGroup represents a Machine's "volume_groups" list item.
// This type should not be used directly.
type MachineVolumeGroup struct {
	SystemID   string `json:"system_id,omitempty"`
	ID         int    `json:"id,omitempty"`
	Incomplete bool   `json:"__incomplete__,omitempty"`
}

// MachineBCache represents a Machine's "bcaches" list item.
// This type should not be used directly.
type MachineBCache struct {
	SystemID   string `json:"system_id,omitempty"`
	ID         int    `json:"id,omitempty"`
	Incomplete bool   `json:"__incomplete__,omitempty"`
}

// MachineRAID represents a Machine's "raids" list item.
// This type should not be used directly.
type MachineRAID struct {
	SystemID   string `json:"system_id,omitempty"`
	ID         int    `json:"id,omitempty"`
	Incomplete bool   `json:"__incomplete__,omitempty"`
}

// MachineSpecialFilesystem represents a Machine's "special_filesystems" list item.
// This type should not be used directly.
type MachineSpecialFilesystem struct {
	FSType       string `json:"fstype,omitempty"`
	Label        string `json:"label,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	MountPoint   string `json:"mount_point,omitempty"`
	MountOptions string `json:"mount_options,omitempty"`
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
	CPUCount      int    `url:"cpu_count,omitempty"`
	Memory        int64  `url:"memory,omitempty"`
	SwapSize      int64  `url:"swap_size,omitempty"`
	PXEMacAddress string `url:"mac_addresses,omitempty"`
	Architecture  string `url:"architecture,omitempty"`
	MinHWEKernel  string `url:"min_hwe_kernel,omitempty"`
	PowerType     string `url:"power_type,omitempty"`
	Hostname      string `url:"hostname,omitempty"`
	Description   string `url:"description,omitempty"`
	Domain        string `url:"domain,omitempty"`
	Pool          string `url:"pool,omitempty"`
	Zone          string `url:"zone,omitempty"`
	Commission    bool   `url:"commission,omitempty"`
}

// MachineCommissionParams enumerates the parameters for the commission operation
type MachineCommissionParams struct {
	EnableSSH            int    `url:"enable_ssh,omitempty"`
	SkipBMCConfig        int    `url:"skip_bmc_config,omitempty"`
	SkipNetworking       int    `url:"skip_networking,omitempty"`
	SkipStorage          int    `url:"skip_storage,omitempty"`
	CommissioningScripts string `url:"commissioning_scripts,omitempty"`
	TestingScripts       string `url:"testing_scripts,omitempty"`
}

// MachineAllocateParams enumerates the options for the allocate operation.
type MachineAllocateParams struct {
	Tags             []string `url:"tags,omitempty"`
	NotTags          []string `url:"not_tags,omitempty"`
	NotInZone        []string `url:"not_in_zone,omitempty"`
	NotInPool        []string `url:"not_in_pool,omitempty"`
	Subnets          []string `url:"subnets,omitempty"`
	NotSubnets       []string `url:"not_subnets,omitempty"`
	Storage          []string `url:"storage,omitempty"`
	Fabrics          []string `url:"fabrics,omitempty"`
	NotFabrics       []string `url:"not_fabrics,omitempty"`
	FabricClasses    []string `url:"fabric_classes,omitempty"`
	NotFabricClasses []string `url:"not_fabric_classes,omitempty"`
	Name             string   `url:"name,omitempty"`
	SystemID         string   `url:"system_id,omitempty"`
	Arch             string   `url:"arch,omitempty"`
	Zone             string   `url:"zone,omitempty"`
	Pool             string   `url:"pool,omitempty"`
	VMHost           string   `url:"pod,omitempty"`
	NotVMHost        string   `url:"not_pod,omitempty"`
	VMHostType       string   `url:"pod_type,omitempty"`
	NotVMHostType    string   `url:"not_pod_type,omitempty"`
	Interfaces       string   `url:"interfaces,omitempty"`
	AgentName        string   `url:"agent_name,omitempty"`
	Comment          string   `url:"comment,omitempty"`
	CPUCount         int      `url:"cpu_count,omitempty"`
	Mem              int64    `url:"mem,omitempty"`
	BridgeFD         int      `url:"bridge_fd,omitempty"`
	BridgeAll        bool     `url:"bridge_all,omitempty"`
	BridgeSTP        bool     `url:"bridge_stp,omitempty"`
	DryRun           bool     `url:"dry_run,omitempty"`
	Verbose          bool     `url:"verbose,omitempty"`
}

// MachineDeployParams enumerates the parameters for the deploy operation
type MachineDeployParams struct {
	UserData       string `url:"user_data,omitempty"`
	DistroSeries   string `url:"distro_series,omitempty"`
	HWEKernel      string `url:"hwe_kernel,omitempty"`
	AgentName      string `url:"agent_name,omitempty"`
	Comment        string `url:"comment,omitempty"`
	BridgeFD       int    `url:"bridge_fd,omitempty"`
	BridgeAll      bool   `url:"bridge_all,omitempty"`
	BridgeSTP      bool   `url:"bridge_stp,omitempty"`
	InstallRackD   bool   `url:"install_rackd,omitempty"`
	InstallKVM     bool   `url:"install_kvm,omitempty"`
	RegisterVMHost bool   `url:"register_vmhost,omitempty"`
	EnableHwSync   bool   `url:"enable_hw_sync,omitempty"`
}

// MachineReleaseParams enumerates the parameters for the release operation
type MachineReleaseParams struct {
	Comment     string `url:"comment,omitempty"`
	Erase       bool   `url:"erase,omitempty"`
	Force       bool   `url:"force,omitempty"`
	QuickErase  bool   `url:"quick_erase,omitempty"`
	SecureErase bool   `url:"secure_erase,omitempty"`
}
