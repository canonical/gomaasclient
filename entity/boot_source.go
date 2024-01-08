package entity

type BootSource struct {
	Created         string `json:"created,omitempty"`
	Updated         string `json:"updated,omitempty"`
	URL             string `json:"url,omitempty"`
	KeyringFilename string `json:"keyring_filename,omitempty"`
	KeyringData     string `json:"keyring_data,omitempty"`
	ResourceURI     string `json:"resource_uri,omitempty"`
	ID              int    `json:"id,omitempty"`
}

type BootSourceParams struct {
	KeyringData     string `url:"keyring_data,omitempty"`
	KeyringFilename string `url:"keyring_filename,omitempty"`
	URL             string `url:"url,omitempty"`
}
