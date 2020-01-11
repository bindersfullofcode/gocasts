package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	BaseURL string
}

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
}

func New() *Client {
	return &Client{
		BaseURL: "https://jsonplaceholder.typicode.com",
	}
}

func (c *Client) Get() ([]User, error) {
	res, err := http.Get(fmt.Sprintf("%v/users", c.BaseURL))
	if err != nil {
		return nil, errors.Wrap(err, "get users request failed")
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, errors.Errorf("get users request returned %v status code", res.Status)
	}

	var users []User
	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal get users response")
	}

	return users, nil
}
