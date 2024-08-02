package entity

import (
	"encoding/json"
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/canonical/gomaasclient/entity/node"
)

// Machine represents the MAAS Machine endpoint.
type Machine struct {
	HwNextSync              MAASTime                   `json:"next_sync,omitempty"`
	HwLastSync              MAASTime                   `json:"last_sync,omitempty"`
	OwnerData               interface{}                `json:"owner_data,omitempty"`
	HardwareInfo            map[string]string          `json:"hardware_info,omitempty"`
	WorkloadAnnotations     map[string]string          `json:"workload_annotations,omitempty"`
	OSystem                 string                     `json:"osystem,omitempty"`
	FQDN                    string                     `json:"fqdn,omitempty"`
	StatusName              string                     `json:"status_name,omitempty"`
	StatusMessage           string                     `json:"status_message,omitempty"`
	StatusAction            string                     `json:"status_action,omitempty"`
	Description             string                     `json:"description,omitempty"`
	CPUTestStatusName       string                     `json:"cpu_test_status_name,omitempty"`
	OtherTestStatusName     string                     `json:"other_test_status_name,omitempty"`
	Hostname                string                     `json:"hostname,omitempty"`
	ResourceURI             string                     `json:"resource_uri,omitempty"`
	Architecture            string                     `json:"architecture,omitempty"`
	PowerType               string                     `json:"power_type,omitempty"`
	MemoryTestStatusName    string                     `json:"memory_test_status_name,omitempty"`
	PowerState              string                     `json:"power_state,omitempty"`
	DistroSeries            string                     `json:"distro_series,omitempty"`
	MinHWEKernel            string                     `json:"min_hwe_kernel,omitempty"`
	CommissioningStatusName string                     `json:"commissioning_status_name,omitempty"`
	SystemID                string                     `json:"system_id,omitempty"`
	NodeTypeName            string                     `json:"node_type_name,omitempty"`
	StorageTestStatusName   string                     `json:"storage_test_status_name,omitempty"`
	Owner                   string                     `json:"owner,omitempty"`
	HWEKernel               string                     `json:"hwe_kernel,omitempty"`
	TestingStatusName       string                     `json:"testing_status_name,omitempty"`
	Version                 string                     `json:"version,omitempty"`
	BiosBootMethod          string                     `json:"bios_boot_method,omitempty"`
	HardwareUUID            string                     `json:"hardware_uuid,omitempty"`
	InterfaceTestStatusName string                     `json:"interface_test_status_name,omitempty"`
	NetworkTestStatusName   string                     `json:"network_test_status_name,omitempty"`
	Pool                    ResourcePool               `json:"pool,omitempty"`
	Zone                    Zone                       `json:"zone,omitempty"`
	SpecialFilesystems      []MachineSpecialFilesystem `json:"special_filesystems,omitempty"`
	VirtualBlockDeviceSet   []BlockDevice              `json:"virtualblockdevice_set,omitempty"`
	ISCSIBlockDeviceSet     []BlockDevice              `json:"iscsiblockdevice_set,omitempty"`
	ServiceSet              []MachineServiceSet        `json:"service_set,omitempty"`
	PhysicalBlockDeviceSet  []BlockDevice              `json:"physicalblockdevice_set,omitempty"`
	TagNames                []string                   `json:"tag_names,omitempty"`
	RAIDs                   []MachineRAID              `json:"raids,omitempty"`
	IPAddresses             []net.IP                   `json:"ip_addresses,omitempty"`
	BCaches                 []MachineBCache            `json:"bcaches,omitempty"`
	InterfaceSet            []NetworkInterface         `json:"interface_set,omitempty"`
	VolumeGroups            []MachineVolumeGroup       `json:"volume_groups,omitempty"`
	CacheSets               []MachineCacheSet          `json:"cache_sets,omitempty"`
	BlockDeviceSet          []BlockDevice              `json:"blockdevice_set,omitempty"`
	NUMANodeSet             []NUMANode                 `json:"numanode_set,omitempty"`
	DefaultGateways         struct {
		IPv4 struct {
			GatewayIP net.IP `json:"gateway_ip,omitempty"`
			LinkID    int    `json:"link_id,omitempty"`
		} `json:"ipv4,omitempty"`
		IPv6 struct {
			GatewayIP net.IP `json:"gateway_ip,omitempty"`
			LinkID    int    `json:"link_id,omitempty"`
		} `json:"ipv6,omitempty"`
	} `json:"default_gateways,omitempty"`
	Domain                       Domain           `json:"domain,omitempty"`
	BootDisk                     BlockDevice      `json:"boot_disk,omitempty"`
	BootInterface                NetworkInterface `json:"boot_interface,omitempty"`
	VMHost                       VMHost           `json:"pod,omitempty"`
	CurrentTestingResultID       int              `json:"current_testing_result_id,omitempty"`
	Memory                       int64            `json:"memory,omitempty"`
	NodeType                     int              `json:"node_type,omitempty"`
	HwSyncInterval               int              `json:"sync_interval,omitempty"`
	CPUTestStatus                int              `json:"cpu_test_status,omitempty"`
	AddressTTL                   int              `json:"address_ttl,omitempty"`
	Storage                      float64          `json:"storage,omitempty"`
	CPUSpeed                     int              `json:"cpu_speed,omitempty"`
	CPUCount                     int              `json:"cpu_count,omitempty"`
	Status                       node.Status      `json:"status,omitempty"`
	CurrentInstallationResultID  int              `json:"current_installation_result_id,omitempty"`
	CurrentCommissioningResultID int              `json:"current_commissioning_result_id,omitempty"`
	CommissioningStatus          int              `json:"commissioning_status,omitempty"`
	OtherTestStatus              int              `json:"other_test_status,omitempty"`
	TestingStatus                int              `json:"testing_status,omitempty"`
	InterfaceTestStatus          int              `json:"interface_test_status,omitempty"`
	NetworkTestStatus            int              `json:"network_test_status,omitempty"`
	StorageTestStatus            int              `json:"storage_test_status,omitempty"`
	SwapSize                     int64            `json:"swap_size,omitempty"`
	MemoryTestStatus             int              `json:"memory_test_status,omitempty"`
	EnableHwSync                 bool             `json:"enable_hw_sync,omitempty"`
	DisableIPv4                  bool             `json:"disable_ipv4,omitempty"`
	Netboot                      bool             `json:"netboot,omitempty"`
	Locked                       bool             `json:"locked,omitempty"`
	EphemeralDeploy              bool             `json:"ephemeral_deploy,omitempty"`
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
	Incomplete bool   `json:"__incomplete__,omitempty"` //nolint:tagliatelle // MAAS returns this field intentionally
}

// MachineVolumeGroup represents a Machine's "volume_groups" list item.
// This type should not be used directly.
type MachineVolumeGroup struct {
	SystemID   string `json:"system_id,omitempty"`
	ID         int    `json:"id,omitempty"`
	Incomplete bool   `json:"__incomplete__,omitempty"` //nolint:tagliatelle // MAAS returns this field intentionally
}

// MachineBCache represents a Machine's "bcaches" list item.
// This type should not be used directly.
type MachineBCache struct {
	SystemID   string `json:"system_id,omitempty"`
	ID         int    `json:"id,omitempty"`
	Incomplete bool   `json:"__incomplete__,omitempty"` //nolint:tagliatelle // MAAS returns this field intentionally
}

// MachineRAID represents a Machine's "raids" list item.
// This type should not be used directly.
type MachineRAID struct {
	SystemID   string `json:"system_id,omitempty"`
	ID         int    `json:"id,omitempty"`
	Incomplete bool   `json:"__incomplete__,omitempty"` //nolint:tagliatelle // MAAS returns this field intentionally
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

// MachinePowerState represent current machines power state
type MachinePowerState NodePowerState

// MachineToken represents Machine token
type MachineToken struct {
	TokenKey    string `json:"token_key,omitempty"`
	TokenSecret string `json:"token_secret,omitempty"`
	ConsumerKey string `json:"consumer_key,omitempty"`
}

// MachineParams enumerates the parameters for the machine create/update operation
type MachineParams struct {
	Domain       string   `url:"domain,omitempty"`
	Architecture string   `url:"architecture,omitempty"`
	MinHWEKernel string   `url:"min_hwe_kernel,omitempty"`
	Hostname     string   `url:"hostname,omitempty"`
	Description  string   `url:"description,omitempty"`
	PowerType    string   `url:"power_type,omitempty"`
	Pool         string   `url:"pool,omitempty"`
	Zone         string   `url:"zone,omitempty"`
	MACAddresses []string `url:"mac_addresses,omitempty"`
	Memory       int64    `url:"memory,omitempty"`
	SwapSize     int64    `url:"swap_size,omitempty"`
	CPUCount     int      `url:"cpu_count,omitempty"`
	Commission   bool     `url:"commission,omitempty"`
}

// MachineCommissionParams enumerates the parameters for the commission operation
type MachineCommissionParams struct {
	CommissioningScripts string `url:"commissioning_scripts,omitempty"`
	TestingScripts       string `url:"testing_scripts,omitempty"`
	EnableSSH            int    `url:"enable_ssh,omitempty"`
	SkipBMCConfig        int    `url:"skip_bmc_config,omitempty"`
	SkipNetworking       int    `url:"skip_networking,omitempty"`
	SkipStorage          int    `url:"skip_storage,omitempty"`
}

// MachineAllocateParams enumerates the options for the allocate operation.
type MachineAllocateParams struct {
	AgentName        string   `url:"agent_name,omitempty"`
	SystemID         string   `url:"system_id,omitempty"`
	Pool             string   `url:"pool,omitempty"`
	Zone             string   `url:"zone,omitempty"`
	Comment          string   `url:"comment,omitempty"`
	Interfaces       string   `url:"interfaces,omitempty"`
	NotVMHostType    string   `url:"not_pod_type,omitempty"`
	VMHostType       string   `url:"pod_type,omitempty"`
	NotVMHost        string   `url:"not_pod,omitempty"`
	VMHost           string   `url:"pod,omitempty"`
	Arch             string   `url:"arch,omitempty"`
	Name             string   `url:"name,omitempty"`
	NotFabrics       []string `url:"not_fabrics,omitempty"`
	Tags             []string `url:"tags,omitempty"`
	NotInPool        []string `url:"not_in_pool,omitempty"`
	NotInZone        []string `url:"not_in_zone,omitempty"`
	FabricClasses    []string `url:"fabric_classes,omitempty"`
	NotTags          []string `url:"not_tags,omitempty"`
	Fabrics          []string `url:"fabrics,omitempty"`
	Storage          []string `url:"storage,omitempty"`
	NotSubnets       []string `url:"not_subnets,omitempty"`
	NotFabricClasses []string `url:"not_fabric_classes,omitempty"`
	Subnets          []string `url:"subnets,omitempty"`
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
	UserData        string `url:"user_data,omitempty"`
	DistroSeries    string `url:"distro_series,omitempty"`
	HWEKernel       string `url:"hwe_kernel,omitempty"`
	AgentName       string `url:"agent_name,omitempty"`
	Comment         string `url:"comment,omitempty"`
	BridgeFD        int    `url:"bridge_fd,omitempty"`
	BridgeAll       bool   `url:"bridge_all,omitempty"`
	BridgeSTP       bool   `url:"bridge_stp,omitempty"`
	InstallRackD    bool   `url:"install_rackd,omitempty"`
	InstallKVM      bool   `url:"install_kvm,omitempty"`
	RegisterVMHost  bool   `url:"register_vmhost,omitempty"`
	EnableHwSync    bool   `url:"enable_hw_sync,omitempty"`
	EphemeralDeploy bool   `url:"ephemeral_deploy,omitempty"`
}

// MachineReleaseParams enumerates the parameters for the release operation
type MachineReleaseParams struct {
	Comment     string `url:"comment,omitempty"`
	Erase       bool   `url:"erase,omitempty"`
	Force       bool   `url:"force,omitempty"`
	QuickErase  bool   `url:"quick_erase,omitempty"`
	SecureErase bool   `url:"secure_erase,omitempty"`
}

// MachinePowerOnParams enumerates the parameters for the machine power on operation
// UserData should be Base64-encoded data
type MachinePowerOnParams NodePowerOnParams

// MachinePowerOffParams enumerates the parameters for the machine power off operation
type MachinePowerOffParams NodePowerOffParams

// MachinesParams enumerates the parameters for the get machines operation
type MachinesParams NodeGetParams

// MachineDetails represent the MAAS machine details
type MachineDetails NodeDetails
