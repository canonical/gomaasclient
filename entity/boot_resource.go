package entity

type BootResource struct {
	Sets         map[string]BootResourceSet `json:"sets,omitempty"`
	Type         string                     `json:"type,omitempty"`
	Name         string                     `json:"name,omitempty"`
	Architecture string                     `json:"architecture,omitempty"`
	ResourceURI  string                     `json:"resource_uri,omitempty"`
	LastDeployed string                     `json:"last_deployed,omitempty"`
	Subarches    string                     `json:"subarches,omitempty"`
	ID           int                        `json:"id,omitempty"`
}

type BootResourceParams struct {
	Name         string `url:"name,omitempty"`
	Architecture string `url:"architecture,omitempty"`
	SHA256       string `url:"sha256,omitempty"`
	Size         string `url:"size,omitempty"`
	Title        string `url:"title,omitempty"`
	Filetype     string `url:"filetype,omitempty"`
	BaseImage    string `url:"base_image,omitempty"`
	Content      string `url:"content,omitempty"`
}

type BootResourcesReadParams struct {
	Type string `json:"type,omitempty"`
}

// BootResourceSet represents a BootResource's "set".
// This type should not be used directly.
type BootResourceSet struct {
	Files    map[string]BootResourceSetFile `json:"files,omitempty"`
	Version  string                         `json:"version,omitempty"`
	Label    string                         `json:"label,omitempty"`
	Size     int64                          `json:"size,omitempty"`
	Complete bool                           `json:"complete,omitempty"`
}

// BootResourceSetFile represents a BootResource set's "file".
// This type should not be used directly.
type BootResourceSetFile struct {
	Filename string `json:"filename,omitempty"`
	Filetype string `json:"filetype,omitempty"`
	SHA256   string `json:"sha256,omitempty"`
	Size     int64  `json:"size,omitempty"`
	Complete bool   `json:"complete,omitempty"`
}
