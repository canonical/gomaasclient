package entity

// BlockDevicePartition represents the MaaS block device partition endpoint.
type BlockDevicePartition struct {
	UUID        string              `json:"uuid"`
	Size        int64               `json:"size"`
	Bootable    bool                `json:"bootable"`
	Tags        []string            `json:"tags,omitempty"`
	SystemID    string              `json:"system_id"`
	ID          int                 `json:"id"`
	DeviceID    int                 `json:"device_id"`
	FileSystem  PartitionFileSystem `json:"filesystem"`
	UsedFor     string              `json:"used_for,omitempty"`
	Type        string              `json:"type"`
	Path        string              `json:"path"`
	ResourceURI string              `json:"resource_uri"`
}

type PartitionFileSystem struct {
	FSType       string `json:"fstype"`
	UUID         string `json:"uuid"`
	MountPoint   string `json:"mount_point"`
	Label        string `json:"label,omitempty"`
	MountOptions string `json:"mount_options,omitempty"`
}

type BlockDevicePartitionParams struct {
	Size     int64  `url:"size,omitempty"`
	UUID     string `url:"uuid,omitempty"`
	Bootable bool   `url:"bootable"`
}
