package gmail

import (
	"google.golang.org/api/gmail/v1"
)

type Client struct {
	service *gmail.Service
	user    string
}

const (
	user = "me"
)

func NewClient(service *gmail.Service) *Client {
	return &Client{service: service, user: user}
}
