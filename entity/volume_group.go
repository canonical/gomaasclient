package entity

// VolumeGroup represents the MAAS VolumeGroup endpoint.
type VolumeGroup struct {
	LogicalVolumes     []LogicalVolume `json:"logical_volumes,omitempty"`
	Devices            interface{}     `json:"devices,omitempty"`
	HumanSize          string          `json:"human_size,omitempty"`
	UUID               string          `json:"uuid,omitempty"`
	HumanAvailableSize string          `json:"human_available_size,omitempty"`
	SystemID           string          `json:"system_id,omitempty"`
	ResourceURI        string          `json:"resource_uri,omitempty"`
	HumanUsedSize      string          `json:"human_used_size,omitempty"`
	Name               string          `json:"name,omitempty"`
	Size               int64           `json:"size,omitempty"`
	UsedSize           int64           `json:"used_size,omitempty"`
	AvailableSize      int64           `json:"available_size,omitempty"`
	ID                 int             `json:"id,omitempty"`
}

// VolumeGroupParams enumerates the parameters for the volume group operation
type VolumeGroupParams struct {
	Name         string   `url:"name,omitempty"`
	BlockDevices []string `url:"block_devices,omitempty"`
	Partitions   []string `url:"partitions,omitempty"`
	UUID         string   `url:"uuid,omitempty"`
}

// LogicalVolume represents the LogicalVolume object.
type LogicalVolume struct {
	SystemID      string `json:"system_id,omitempty"`
	ID            int    `json:"id,omitempty"`
	Name          string `url:"name,omitempty"`
	Size          int64  `json:"size,omitempty"`
	UsedSize      int64  `json:"used_size,omitempty"`
	AvailableSize int64  `json:"available_size,omitempty"`
	UUID          string `url:"uuid,omitempty"`
	ResourceURI   string `json:"resource_uri,omitempty"`
}

// LogicalVolumeParams enumerates the parameters for the logical volume operation
type LogicalVolumeParams struct {
	Name string `url:"name,omitempty"`
	Size int64  `json:"size,omitempty"`
	UUID string `url:"uuid,omitempty"`
}
