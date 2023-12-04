package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Users implements api.Users
type Users struct {
	ApiClient ApiClient
}

func (u *Users) client() ApiClient {
	return u.ApiClient.GetSubObject("users")
}

// Get fetches a list of User objects
func (u *Users) Get() (users []entity.User, err error) {
	err = u.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &users)
	})
	return
}

// Create creates a new User
func (u *Users) Create(params *entity.UserParams) (user *entity.User, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	user = new(entity.User)
	err = u.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, user)
	})
	return
}
