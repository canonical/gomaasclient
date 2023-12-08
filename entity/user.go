package entity

type User struct {
	UserName    string `json:"username"`
	Email       string `json:"email"`
	ResourceURI string `json:"resource_uri"`
	IsSuperUser bool   `json:"is_superuser"`
	IsLocal     bool   `json:"is_local"`
}

type UserParams struct {
	UserName    string `url:"username"`
	Password    string `url:"password"`
	Email       string `url:"email"`
	IsSuperUser bool   `url:"is_superuser,int"`
}
