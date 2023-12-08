package entity

// VolumeGroup represents the MAAS VolumeGroup endpoint.
type VolumeGroup struct {
	LogicalVolumes     interface{} `json:"logical_volumes,omitempty"`
	Devices            interface{} `json:"devices,omitempty"`
	HumanSize          string      `json:"human_size,omitempty"`
	UUID               string      `json:"uuid,omitempty"`
	HumanAvailableSize string      `json:"human_available_size,omitempty"`
	SystemID           string      `json:"system_id,omitempty"`
	ResourceURI        string      `json:"resource_uri,omitempty"`
	HumanUsedSize      string      `json:"human_used_size,omitempty"`
	Name               string      `json:"name,omitempty"`
	Size               int64       `json:"size,omitempty"`
	UsedSize           int64       `json:"used_size,omitempty"`
	AvailableSize      int64       `json:"available_size,omitempty"`
	ID                 int         `json:"id,omitempty"`
}
