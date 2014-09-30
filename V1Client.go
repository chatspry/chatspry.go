package chatspry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type V1Client struct {
	BaseURL   string
	SessionID string
	User      *User
}

type User struct {
	ID        *string    `json:"id"`
	Name      *string    `json:"name"`
	Handle    *string    `json:"handle"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewV1Client(baseUrl string) *V1Client {
	return &V1Client{
		BaseURL: baseUrl,
	}
}

func (client *V1Client) Login(username, passphrase string) error {
	url := fmt.Sprintf("%v/v1/session", client.BaseURL)
	body, err := json.Marshal(&map[string]interface{}{
		"session": &map[string]string{
			"identifier": username,
			"passphrase": passphrase,
		},
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var session struct {
		Session struct {
			ID string `json:"id"`
		} `json:"session"`
		User *User `json:"user"`
	}
	err = json.Unmarshal(body, &session)
	if err != nil {
		return err
	}

	client.SessionID = session.Session.ID
	client.User = session.User

	return nil
}
