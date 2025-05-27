package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// User implements api.User
type User struct {
	APIClient APIClient
}

func (u *User) client(userName string) APIClient {
	return u.APIClient.GetSubObject(fmt.Sprintf("users/%s", userName))
}

// Get fetches a User by username
func (u *User) Get(userName string) (*entity.User, error) {
	user := new(entity.User)
	err := u.client(userName).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, user)
	})

	return user, err
}

// Delete deletes a given User
func (u *User) Delete(params *entity.UserDeleteParams) error {
	qsp, err := query.Values(params)
	if err != nil {
		return err
	}
	return u.client(params.UserName).DeleteWithParams(qsp)
}
