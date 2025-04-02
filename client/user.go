package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// User implements api.User
type User struct {
	APIClient APIClient
}

func (u *User) client(userName string) *APIClient {
	return u.APIClient.SubClient(fmt.Sprintf("users/%s", userName))
}

// Get fetches a User by username
func (u *User) Get(ctx context.Context, userName string) (*entity.User, error) {
	user := new(entity.User)
	err := u.client(userName).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, user)
	})

	return user, err
}

// Delete deletes a given User
func (u *User) Delete(ctx context.Context, userName string) error {
	return u.client(userName).Delete(ctx)
}
