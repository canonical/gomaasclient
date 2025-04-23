package entity

// VolumeGroup represents the MAAS VolumeGroup endpoint.
type VolumeGroup struct {
	Devices            interface{}          `json:"devices,omitempty"`
	ResourceURI        string               `json:"resource_uri,omitempty"`
	HumanSize          string               `json:"human_size,omitempty"`
	UUID               string               `json:"uuid,omitempty"`
	HumanAvailableSize string               `json:"human_available_size,omitempty"`
	SystemID           string               `json:"system_id,omitempty"`
	HumanUsedSize      string               `json:"human_used_size,omitempty"`
	Name               string               `json:"name,omitempty"`
	LogicalVolumes     []VirtualBlockDevice `json:"logical_volumes,omitempty"`
	Size               int64                `json:"size,omitempty"`
	UsedSize           int64                `json:"used_size,omitempty"`
	AvailableSize      int64                `json:"available_size,omitempty"`
	ID                 int                  `json:"id,omitempty"`
}

// VolumeGroupCreateParams enumerates the parameters for the volume group create operation.
type VolumeGroupCreateParams struct {
	Name         string   `url:"name,omitempty"`
	UUID         string   `url:"uuid,omitempty"`
	BlockDevices []string `url:"block_devices,omitempty"`
	Partitions   []string `url:"partitions,omitempty"`
}

// VolumeGroupUpdateParams enumerates the parameters for the volume group update operation.
type VolumeGroupUpdateParams struct {
	Name               string   `url:"name,omitempty"`
	UUID               string   `url:"uuid,omitempty"`
	AddBlockDevices    []string `url:"add_block_devices,omitempty"`
	RemoveBlockDevices []string `url:"remove_block_devices,omitempty"`
	AddPartitions      []string `url:"add_partitions,omitempty"`
	RemovePartitions   []string `url:"remove_partitions,omitempty"`
}

// LogicalVolumeParams enumerates the parameters for the logical volume operation.
type LogicalVolumeParams struct {
	Name string `url:"name,omitempty"`
	UUID string `url:"uuid,omitempty"`
	Size int64  `url:"size,omitempty"`
}

// VirtualBlockDevice represents a logical volume and extends BlockDevice.
type VirtualBlockDevice struct {
	BlockDevice
}
