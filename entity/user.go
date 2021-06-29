package entity

type User struct {
	IsSuperUser bool   `json:"is_superuser"`
	IsLocal     bool   `json:"is_local"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	ResourceURI string `json:"resource_uri"`
}
