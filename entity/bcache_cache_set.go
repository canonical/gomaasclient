package entity

type BCacheCacheSet struct {
	Name        string      `json:"name:omitempty"`
	SystemID    string      `json:"system_id,omitempty"`
	ResourceURI string      `json:"resource_uri,omitempty"`
	CacheDevice BlockDevice `json:"cache_device,omitempty"`
	ID          int         `json:"id,omitempty"`
}

type BCacheCacheSetParams struct {
	CacheDevice    string `url:"cache_device,omitempty"`
	CachePartition string `url:"cache_partition,omitempty"`
}
