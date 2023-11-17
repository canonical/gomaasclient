package entity

// VolumeGroup represents the MaaS VolumeGroup endpoint.
type VolumeGroup struct {
	HumanUsedSize      string      `json:"human_used_size,omitempty"`
	UUID               string      `json:"uuid,omitempty"`
	HumanAvailableSize string      `json:"human_available_size,omitempty"`
	SystemID           string      `json:"system_id,omitempty"`
	ID                 int         `json:"id,omitempty"`
	LogicalVolumes     interface{} `json:"logical_volumes,omitempty"`
	Size               int64       `json:"size,omitempty"`
	Devices            interface{} `json:"devices,omitempty"`
	AvailableSize      int64       `json:"available_size,omitempty"`
	UsedSize           int64       `json:"used_size,omitempty"`
	HumanSize          string      `json:"human_size,omitempty"`
	Name               string      `json:"name,omitempty"`
	ResourceURI        string      `json:"resource_uri,omitempty"`
}
