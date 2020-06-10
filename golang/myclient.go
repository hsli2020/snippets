package myclient

import "bytes"
import "encoding/json"
import "net/http"

type Client struct {
	Client *http.Client
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Client) CreateUser(name string) (*User, error) {
	data, err := json.Marshal(map[string]string{
		"name": name,
	})
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Post("https://api.exmaple.com/users", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

package main

import "net/http"
import "log"
import "example.com/myclient"

func main() {
  client := myclient.Client{
    Client: &http.Client{},
  }
  
  _, err := client.CreateUser("Boaty McBoatface")
  if err != nil {
    log.Fatalln(err)
  }
}

//////////// Better ////////////////////////////////////////

package myclient

import "net/http"
import "time"

type Client struct {
	APIKey     string
	BaseURL    string
	httpClient *http.Client
}

func NewClient(apiKey string, opts ...func(*Client) error) (*Client, error) {
	client := &Client{
		APIKey:     apiKey,
		BaseURL:    "api.example.com",
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}

	for _, opt := range opts {
		err := opt(client)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

// WithHTTPClient allows users of our API client to override the default HTTP Client!
func WithHTTPClient(client *http.Client) func(*Client) error {
  return func(c *Client) error {
    c.httpClient = client
    return nil
  }
}

// client.go
package myclient

type Client struct {
  httpClient *http.Client
  
  Users      *UserService
  Accounts   *AccountService
}

func NewClient() *Client {
  c := &Client{
    httpClient: &http.Client{},
  }
  
  c.Users = &UserService{client: c}
  c.Accounts = &AccountService{client: c}
  
  return c
}

// users.go
package myclient

type UserService struct {
  client *Client
}

type User struct {
  ID   int `json:"id"`
  Name string `json"name"`
}

func (u *UserService) Get(name string) *User {
  resp, _ := u.client.httpClient.Get("api.example.com/users/" + name)
  
  defer resp.Body.Close()

	var user User
	_ = json.NewDecoder(resp.Body).Decode(&user)
	
  return &user
}