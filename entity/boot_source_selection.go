package entity

type BootSourceSelection struct {
	OS           string   `json:"os,omitempty"`
	Release      string   `json:"release,omitempty"`
	ResourceURI  string   `json:"resource_uri,omitempty"`
	Arches       []string `json:"arches,omitempty"`
	Subarches    []string `json:"subarches,omitempty"`
	Labels       []string `json:"labels,omitempty"`
	ID           int      `json:"id,omitempty"`
	BootSourceID int      `json:"boot_source_id,omitempty"`
}

type BootSourceSelectionParams struct {
	OS        string   `url:"os,omitempty"`
	Release   string   `url:"release,omitempty"`
	Arches    []string `url:"arches,omitempty"`
	Subarches []string `url:"subarches,omitempty"`
	Labels    []string `url:"labels,omitempty"`
}
