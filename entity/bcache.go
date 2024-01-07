package entity

type BCache struct {
	HumanSize     string         `json:"human_size,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	Name          string         `json:"name,omitempty"`
	SystemID      string         `json:"system_id"`
	CacheMode     string         `json:"cache_mode,omitempty"`
	ResourceURI   string         `json:"resource_uri,omitempty"`
	VirtualDevice BlockDevice    `json:"virtual_device,omitempty"`
	BackingDevice BlockDevice    `json:"backing_device,omitempty"`
	CacheSet      BCacheCacheSet `json:"cache_set"`
	ID            int            `json:"id,omitempty"`
	Size          int64          `json:"size,omitempty"`
}

type BCacheParams struct {
	Name             string `url:"name,omitempty"`
	UUID             string `url:"uuid,omitempty"`
	CacheSet         string `url:"cache_set,omitempty"`
	BackingDevice    string `url:"backing_device,omitempty"`
	BackingPartition string `url:"backing_partition,omitempty"`
	CacheMode        string `url:"cache_mode,omitempty"`
}
