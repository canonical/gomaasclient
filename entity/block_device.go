package entity

// BlockDevice represents the MaaS BlockDevice endpoint.
type BlockDevice struct {
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
}

type BlockDeviceParams struct {
	Name      string `url:"name,omitempty"`
	Size      int64  `url:"size,omitempty"`
	BlockSize int    `url:"block_size,omitempty"`
	UUID      string `url:"uuid,omitempty"`
	Model     string `url:"model,omitempty"`
	Serial    string `url:"serial,omitempty"`
	IDPath    string `url:"id_path,omitempty"`
}
