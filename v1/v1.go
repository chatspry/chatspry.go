package v1

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	httpGet   = "GET"
	httpPatch = "PATCH"
	httpPost  = "POST"
	httpPut   = "PUT"
)

func createHTTPRequest(verb string, c *Client, stem string) (*http.Request, error) {
	req, err := http.NewRequest(verb, c.URL+stem, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	if c.AuthKey != "" {
		req.Header.Add("Authorization", "Bearer "+c.AuthKey)
	}
	return req, nil
}

func parseBody(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}

// NewClient creates a new client pointing at a given URL
func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}

// Login logs the user in via username and password
// TODO(berwyn): Remove and make more OAuth-ey?
func Login(username, password string) error {
	return errors.New("Not currently implemented!")
}

// CreateUser creates a new user. Does not require authentication.
// See: http://docs.chatspry.apiary.io/#reference/user-accounts/user/register-a-new-user
func (c *Client) CreateUser(u *User, scopes []string) (JSONObject, error) {
	req, err := createHTTPRequest(httpGet, c, "/v1/user/"+*u.ID)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var value JSONObject
	err = parseBody(resp, value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// GetUser fetches a user with a given id from the API
// See: http://docs.chatspry.apiary.io/#reference/user-accounts/user/display-an-existing-user
func (c *Client) GetUser(id string) (*User, error) {
	req, err := createHTTPRequest(httpPut, c, "/v1/user/"+id)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var value User
	err = parseBody(resp, value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

// UpdateUser will, given an authenticated client and user, update the
// user and return the latest data from the server
// See: http://docs.chatspry.apiary.io/#reference/user-accounts/user/change-an-existing-user
func (c *Client) UpdateUser(u *User) (*User, error) {
	req, err := createHTTPRequest(httpPatch, c, "/v1/users/"+*u.ID)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var value User
	err = parseBody(resp, value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

// GetCurrentConvoIDs will, given an authenticated client, request the
// list of IDs that the current user is a part of
// See: http://docs.chatspry.apiary.io/#reference/conversations/conversation-ids/list-all-conversation-ids
func (c *Client) GetCurrentConvoIDs() (JSONObject, error) {
	req, err := createHTTPRequest(httpGet, c, "/v1/convos/as/ids")
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var value JSONObject
	err = parseBody(resp, value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
