package entity

// BlockDevice represents the MAAS BlockDevice endpoint.
type BlockDevice struct {
	Filesystem         PartitionFileSystem    `json:"filesystem,omitempty"`
	Path               string                 `json:"path,omitempty"`
	Model              string                 `json:"model,omitempty"`
	ResourceURI        string                 `json:"resource_uri,omitempty"`
	PartitionTableType string                 `json:"partition_table_type,omitempty"`
	UUID               string                 `json:"uuid,omitempty"`
	Serial             string                 `json:"serial,omitempty"`
	IDPath             string                 `json:"id_path,omitempty"`
	UsedFor            string                 `json:"used_for,omitempty"`
	FirmwareVersion    string                 `json:"firmware_version,omitempty"`
	SystemID           string                 `json:"system_id,omitempty"`
	StoragePool        string                 `json:"storage_pool,omitempty"`
	Type               string                 `json:"type,omitempty"`
	Name               string                 `json:"name,omitempty"`
	Partitions         []BlockDevicePartition `json:"partitions,omitempty"`
	Tags               []string               `json:"tags,omitempty"`
	Size               int64                  `json:"size,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	BlockSize          int                    `json:"block_size,omitempty"`
	UsedSize           int64                  `json:"used_size,omitempty"`
	AvailableSize      int64                  `json:"available_size,omitempty"`
	NUMANode           int                    `json:"numa_node,omitempty"`
}

type BlockDeviceParams struct {
	Name      string `url:"name,omitempty"`
	UUID      string `url:"uuid,omitempty"`
	Model     string `url:"model,omitempty"`
	Serial    string `url:"serial,omitempty"`
	IDPath    string `url:"id_path,omitempty"`
	Size      int64  `url:"size,omitempty"`
	BlockSize int    `url:"block_size,omitempty"`
}
