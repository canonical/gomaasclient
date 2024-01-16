package entity

// AuthorisationToken represents user account authorisation Token
type AuthorisationToken struct {
	Name        string `json:"name,omitempty"`
	TokenKey    string `json:"token_key,omitempty"`
	TokenSecret string `json:"token_secret,omitempty"`
	ConsumerKey string `json:"consumer_key,omitempty"`
}

// AuthorisationTokenListItem represents user account authorisation token in list API
type AuthorisationTokenListItem struct {
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
}
