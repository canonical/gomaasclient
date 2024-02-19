package entity

type HardwareType int

// HardwareType referring from MAAS server
// https://github.com/maas/maas/blob/deab73792a4fe839a2e84a926a6c728d510fc9ad/src/metadataserver/enum.py#L126
const (
	NODE HardwareType = iota
	CPU
	MEMORY
	STORAGE
	NETWORK
	GPU
)

type NodeDeviceBus int

// NodeDeviceBus referring from MAAS server
// https://github.com/maas/maas/blob/deab73792a4fe839a2e84a926a6c728d510fc9ad/src/maasserver/enum.py#L910
const (
	PCIE NodeDeviceBus = iota + 1
	USB
)

// NodeDevice represents a Node Device of a MAAS Node
type NodeDevice struct {
	SystemID                 string           `json:"system_id,omitempty"`
	VendorID                 string           `json:"vendor_id,omitempty"`
	ProductID                string           `json:"product_id,omitempty"`
	VendorName               string           `json:"vendor_name,omitempty"`
	ProductName              string           `json:"product_name,omitempty"`
	CommissioningDriver      string           `json:"commissioning_driver,omitempty"`
	PCIAddress               string           `json:"pci_address,omitempty"`
	BusName                  string           `json:"bus_name,omitempty"`
	HardwareTypeName         string           `json:"hardware_type_name,omitempty"`
	ResourceURI              string           `json:"resource_uri,omitempty"`
	PhysicalBlockDevice      BlockDevice      `json:"physical_blockdevice,omitempty"`
	PhysicalNetworkInterface NetworkInterface `json:"physical_interface,omitempty"`
	Bus                      NodeDeviceBus    `json:"bus,omitempty"`
	HardwareType             HardwareType     `json:"hardware_type,omitempty"`
	BusNumber                int              `json:"bus_number,omitempty"`
	DeviceNumber             int              `json:"device_number,omitempty"`
	NumaNode                 int              `json:"numa_node,omitempty"`
	ID                       int              `json:"id,omitempty"`
}

type NodeDeviceParams struct {
	Bus                 string `url:"bus,omitempty"`
	HardwareType        string `url:"hardware_type,omitempty"`
	VendorID            string `url:"vendor_id,omitempty"`
	ProductID           string `url:"product_id,omitempty"`
	VendorName          string `url:"vendor_name,omitempty"`
	ProductName         string `url:"product_name,omitempty"`
	CommissioningDriver string `url:"commissioning_driver,omitempty"`
}
