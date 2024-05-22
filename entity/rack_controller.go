package entity

import (
	"net"
)

// RackController represents the MAAS Rack Controller endpoint.
type RackController struct {
	HardwareInfo                 map[string]string   `json:"hardware_info,omitempty"`
	TestingStatusName            string              `json:"testing_status_name,omitempty"`
	Architecture                 string              `json:"architecture,omitempty"`
	Version                      string              `json:"version,omitempty"`
	FQDN                         string              `json:"fqdn,omitempty"`
	DistroSeries                 string              `json:"distro_series,omitempty"`
	OSystem                      string              `json:"osystem,omitempty"`
	StatusAction                 string              `json:"status_action,omitempty"`
	SystemID                     string              `json:"system_id,omitempty"`
	Description                  string              `json:"description,omitempty"`
	PowerType                    string              `json:"power_type,omitempty"`
	StorageTestStatusName        string              `json:"storage_test_status_name,omitempty"`
	ResourceURI                  string              `json:"resource_uri,omitempty"`
	PowerState                   string              `json:"power_state,omitempty"`
	HardwareUUID                 string              `json:"hardware_uuid,omitempty"`
	CommissioningStatusName      string              `json:"commissioning_status_name,omitempty"`
	CPUTestStatusName            string              `json:"cpu_test_status_name,omitempty"`
	InterfaceTestStatusName      string              `json:"interface_test_status_name,omitempty"`
	MemoryTestStatusName         string              `json:"memory_test_status_name,omitempty"`
	NetworkTestStatusName        string              `json:"network_test_status_name,omitempty"`
	OtherTestStatusName          string              `json:"other_test_status_name,omitempty"`
	NodeTypeName                 string              `json:"node_type_name,omitempty"`
	Hostname                     string              `json:"hostname,omitempty"`
	Zone                         Zone                `json:"zone,omitempty"`
	InterfaceSet                 []NetworkInterface  `json:"interface_set,omitempty"`
	IPAddresses                  []net.IP            `json:"ip_addresses,omitempty"`
	ServiceSet                   []MachineServiceSet `json:"service_set,omitempty"`
	TagNames                     []string            `json:"tag_names,omitempty"`
	Domain                       Domain              `json:"domain,omitempty"`
	TestingStatus                int                 `json:"testing_status,omitempty"`
	CurrentTestingResultID       int                 `json:"current_testing_result_id,omitempty"`
	CurrentCommissioningResultID int                 `json:"current_commissioning_result_id,omitempty"`
	Memory                       int64               `json:"memory,omitempty"`
	CPUCount                     int                 `json:"cpu_count,omitempty"`
	NodeType                     int                 `json:"node_type,omitempty"`
	CurrentInstallationResultID  int                 `json:"current_installation_result_id,omitempty"`
	SwapSize                     int64               `json:"swap_size,omitempty"`
	CommissioningStatus          int                 `json:"commissioning_status,omitempty"`
	CPUTestStatus                int                 `json:"cpu_test_status,omitempty"`
	InterfaceTestStatus          int                 `json:"interface_test_status,omitempty"`
	MemoryTestStatus             int                 `json:"memory_test_status,omitempty"`
	NetworkTestStatus            int                 `json:"network_test_status,omitempty"`
	OtherTestStatus              int                 `json:"other_test_status,omitempty"`
	StorageTestStatus            int                 `json:"storage_test_status,omitempty"`
	CPUSpeed                     int                 `json:"cpu_speed,omitempty"`
}

// PowerType represents a Rack Controller's power type.
type PowerType struct {
	DriverType      string            `json:"driver_type,omitempty"`
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	Fields          []PowerTypeField  `json:"fields,omitempty"`
	MissingPackages []string          `json:"missing_packages,omitempty"`
	Chassis         bool              `json:"chassis,omitempty"`
	CanProbe        bool              `json:"can_probe,omitempty"`
	Queryable       bool              `json:"queryable,omitempty"`
	Defaults        PowerTypeDefaults `json:"defaults,omitempty"`
}

// PowerTypeField represents a Rack Controller Power Type "field".
// This type should not be used directly.
type PowerTypeField struct {
	Name      string     `json:"name,omitempty"`
	Label     string     `json:"label,omitempty"`
	FieldType string     `json:"field_type,omitempty"`
	Default   string     `json:"default,omitempty"`
	Scope     string     `json:"scope,omitempty"`
	Choices   [][]string `json:"choices,omitempty"`
	Required  bool       `json:"required,omitempty"`
	Secret    bool       `json:"secret,omitempty"`
}

// PowerTypeDefaults represents a Rack Controller Power Type "defaults".
// This type should not be used directly.
type PowerTypeDefaults struct {
	Cores   int   `json:"cores,omitempty"`
	Memory  int64 `json:"memory,omitempty"`
	Storage int64 `json:"storage,omitempty"`
}

// RackControllerDetails represent the MAAS rack controller details
type RackControllerDetails NodeDetails

// RackControllerPowerState represent current rack controller's power state
type RackControllerPowerState NodePowerState

// RackControllerSetZoneParams enumerates the parameters for the rack controller set_zone operation
type RackControllerSetZoneParams struct {
	Zone  string   `url:"zone,omitempty"`
	Nodes []string `url:"nodes,omitempty"`
}

// RackControllersGetParams enumerates the parameters for the get rack controllers operation
type RackControllersGetParams NodeGetParams

// RackControllerParams enumerates the parameters for the update rack controller operation
type RackControllerParams struct {
	Zone      string `url:"zone,omitempty"`
	Domain    string `url:"domain,omitempty"`
	PowerType string `url:"power_type,omitempty"`
}

// RackControllerPowerOnParams enumerates the parameters for the rack controller power on operation
// UserData should be Base64-encoded data
type RackControllerPowerOnParams NodePowerOnParams

// RackControllerPowerOffParams enumerates the parameters for the rack controller power off operation
type RackControllerPowerOffParams NodePowerOffParams
