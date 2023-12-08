package entity

// BlockDevicePartition represents the MAAS block device partition endpoint.
type BlockDevicePartition struct {
	FileSystem  PartitionFileSystem `json:"filesystem"`
	UUID        string              `json:"uuid"`
	SystemID    string              `json:"system_id"`
	UsedFor     string              `json:"used_for,omitempty"`
	Type        string              `json:"type"`
	Path        string              `json:"path"`
	ResourceURI string              `json:"resource_uri"`
	Tags        []string            `json:"tags,omitempty"`
	Size        int64               `json:"size"`
	ID          int                 `json:"id"`
	DeviceID    int                 `json:"device_id"`
	Bootable    bool                `json:"bootable"`
}

type PartitionFileSystem struct {
	FSType       string `json:"fstype"`
	UUID         string `json:"uuid"`
	MountPoint   string `json:"mount_point"`
	Label        string `json:"label,omitempty"`
	MountOptions string `json:"mount_options,omitempty"`
}

type BlockDevicePartitionParams struct {
	UUID     string `url:"uuid,omitempty"`
	Size     int64  `url:"size,omitempty"`
	Bootable bool   `url:"bootable"`
}
