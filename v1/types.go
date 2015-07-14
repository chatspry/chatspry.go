package v1

import (
	"net/http"
	"time"
)

// JSONObject is a semantic type to represent JSON objects
type JSONObject map[string]interface{}

// Client is an API Client
type Client struct {
	http.Client
	URL     string
	AuthKey string
}

// User represents a user on the service
type User struct {
	ID        *string    `json:"id"`
	Name      *string    `json:"name"`
	Handle    *string    `json:"handle"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
