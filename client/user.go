package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
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
func (u *User) Delete(userName string) error {
	return u.client(userName).Delete()
}
