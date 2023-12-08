package entity

// RAID represents the MAAS RAID endpoint.
type RAID struct {
	UUID          string       `json:"uuid,omitempty"`
	Name          string       `json:"name,omitempty"`
	Level         string       `json:"level,omitempty"`
	HumanSize     string       `json:"human_size,omitempty"`
	ResourceURI   string       `json:"resource_uri"`
	SystemID      string       `json:"system_id"`
	Devices       []RAIDDevice `json:"devices,omitempty"`
	SpareDevices  []RAIDDevice `json:"spare_devices,omitempty"`
	VirtualDevice BlockDevice  `json:"virtual_device,omitempty"`
	ID            int          `json:"id,omitempty"`
	Size          int64        `json:"size,omitempty"`
}

// RAIDDevice is a combination of a BlockDevice and BlockDevicePartition since a RAID can contain either at devices field
type RAIDDevice struct {
	Filesystem         PartitionFileSystem    `json:"filesystem,omitempty"`
	Type               string                 `json:"type,omitempty"`
	Model              string                 `json:"model,omitempty"`
	ResourceURI        string                 `json:"resource_uri,omitempty"`
	Name               string                 `json:"name,omitempty"`
	Path               string                 `json:"path,omitempty"`
	Serial             string                 `json:"serial,omitempty"`
	IDPath             string                 `json:"id_path,omitempty"`
	UsedFor            string                 `json:"used_for,omitempty"`
	FirmwareVersion    string                 `json:"firmware_version,omitempty"`
	PartitionTableType string                 `json:"partition_table_type,omitempty"`
	StoragePool        string                 `json:"storage_pool,omitempty"`
	UUID               string                 `json:"uuid,omitempty"`
	SystemID           string                 `json:"system_id,omitempty"`
	Partitions         []BlockDevicePartition `json:"partitions,omitempty"`
	Tags               []string               `json:"tags,omitempty"`
	Size               int64                  `json:"size,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	BlockSize          int                    `json:"block_size,omitempty"`
	DeviceID           int                    `json:"device_id,omitempty"`
	AvailableSize      int64                  `json:"available_size,omitempty"`
	NUMANode           int                    `json:"numa_node,omitempty"`
	UsedSize           int64                  `json:"used_size,omitempty"`
	Bootable           bool                   `json:"bootable,omitempty"`
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
