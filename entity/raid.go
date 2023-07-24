package entity

// RAID represents the MAAS RAID endpoint.
type RAID struct {
	ID            int          `json:"id,omitempty"`
	UUID          string       `json:"uuid,omitempty"`
	Name          string       `json:"name,omitempty"`
	Level         string       `json:"level,omitempty"`
	Devices       []RAIDDevice `json:"devices,omitempty"`
	SpareDevices  []RAIDDevice `json:"spare_devices,omitempty"`
	Size          int64        `json:"size,omitempty"`
	HumanSize     string       `json:"human_size,omitempty"`
	VirtualDevice BlockDevice  `json:"virtual_device,omitempty"`
	ResourceURI   string       `json:"resource_uri"`
	SystemID      string       `json:"system_id"`
}

// RAIDDevice is a combination of a BlockDevice and BlockDevicePartition since a RAID can contain either at devices field
type RAIDDevice struct {
	// BlockDevice fields
	BlockSize          int                    `json:"block_size,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	IDPath             string                 `json:"id_path,omitempty"`
	Model              string                 `json:"model,omitempty"`
	Name               string                 `json:"name,omitempty"`
	Path               string                 `json:"path,omitempty"`
	Serial             string                 `json:"serial,omitempty"`
	Size               int64                  `json:"size,omitempty"`
	Tags               []string               `json:"tags,omitempty"`
	FirmwareVersion    string                 `json:"firmware_version,omitempty"`
	SystemID           string                 `json:"system_id,omitempty"`
	AvailableSize      int64                  `json:"available_size,omitempty"`
	UsedSize           int64                  `json:"used_size,omitempty"`
	PartitionTableType string                 `json:"partition_table_type,omitempty"`
	Partitions         []BlockDevicePartition `json:"partitions,omitempty"`
	Filesystem         PartitionFileSystem    `json:"filesystem,omitempty"`
	StoragePool        string                 `json:"storage_pool,omitempty"`
	UsedFor            string                 `json:"used_for,omitempty"`
	Type               string                 `json:"type,omitempty"`
	UUID               string                 `json:"uuid,omitempty"`
	ResourceURI        string                 `json:"resource_uri,omitempty"`
	NUMANode           int                    `json:"numa_node,omitempty"`

	// BlockDevicePartition unique fields
	Bootable bool `json:"bootable,omitempty"`
	DeviceID int  `json:"device_id,omitempty"`
}

// RAIDUpdateParams enumerates the parameters for the RAID update operation
type RAIDUpdateParams struct {
	Name                  string   `url:"name,omitempty"`
	UUID                  string   `url:"uuid,omitempty"`
	AddBlockDevices       []string `url:"add_block_devices,omitempty"`
	AddPartitions         []string `url:"add_partitions,omitempty"`
	AddSpareDevices       []string `url:"add_spare_devices,omitempty"`
	AddSparePartitions    []string `url:"add_spare_partitions,omitempty"`
	RemoveBlockDevices    []string `url:"remove_block_devices,omitempty"`
	RemovePartitions      []string `url:"remove_partitions,omitempty"`
	RemoveSpareDevices    []string `url:"remove_spare_devices,omitempty"`
	RemoveSparePartitions []string `url:"remove_spare_partitions,omitempty"`
}

// RAIDCreateParams enumerates the parameters for the RAID create operation
type RAIDCreateParams struct {
	Name            string   `url:"name,omitempty"`
	UUID            string   `url:"uuid,omitempty"`
	Level           string   `url:"level,omitempty"`
	BlockDevices    []string `url:"block_devices,omitempty"`
	Partitions      []string `url:"partitions,omitempty"`
	SpareDevices    []string `url:"spare_devices,omitempty"`
	SparePartitions []string `url:"spare_partitions,omitempty"`
}
