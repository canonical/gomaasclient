package entity

type LicenseKey struct {
	DistroSeries string `json:"distro_series"`
	LicenseKey   string `json:"license_key"`
	OSystem      string `json:"osystem"`
	ResourceURI  string `json:"resource_uri"`
}

type LicenseKeyParams struct {
	DistroSeries string `url:"distro_series"`
	LicenseKey   string `url:"license_key"`
	OSystem      string `url:"osystem"`
}
