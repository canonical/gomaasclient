package entity

type Version struct {
	Subversion   string   `json:"subversion,omitempty"`
	Version      string   `json:"version,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
}
