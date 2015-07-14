package v1

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func createHTTPRequest(c *Client, stem string) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.URL+stem, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.AuthKey)
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

// GetUser fetches a user with a given id from the API
func (c *Client) GetUser(id string) (*User, error) {
	req, err := createHTTPRequest(c, "/v1/user/"+id)
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
// list of IDs that the current user is a part of in the form of
// {
//      "user_id": "38459979-e5f9-4c50-8bf6-8721bae33b4e",
//      "convo_ids": [
//          "e22200f1-3fc0-451a-a733-40681bcbc91f",
//          "0f41b840-56ae-417b-bb44-b099273ce50c"
//      ]
// }
func (c *Client) GetCurrentConvoIDs() (JSONObject, error) {
	req, err := createHTTPRequest(c, "/convos/as/ids")
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
