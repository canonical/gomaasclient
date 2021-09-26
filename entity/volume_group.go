package entity

// VolumeGroup represents the MaaS volume-group endpoint.
type VolumeGroup struct {
    SystemID           string                 `json:"system_id,omitempty"`
    UUID               string                 `json:"uuid,omitempty"`
    UsedSize           int                    `json:"used_size,omitempty"`
    Name               string                 `json:"name,omitempty"`
    Size               int                    `json:"size,omitempty"`
    AvailableSize      int                    `json:"available_size,omitempty"`
    ID                 int                    `json:"id,omitempty"`
	ResourceURI        string                 `json:"resource_uri,omitempty"`
}

