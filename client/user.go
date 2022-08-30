package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/maas/gomaasclient/entity"
)

type User struct {
	ApiClient ApiClient
}

func (u *User) client(userName string) ApiClient {
	return u.ApiClient.GetSubObject(fmt.Sprintf("users/%s", userName))
}

func (u *User) Get(userName string) (user *entity.User, err error) {
	user = new(entity.User)
	err = u.client(userName).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, user)
	})
	return
}

func (u *User) Delete(userName string) error {
	return u.client(userName).Delete()
}
