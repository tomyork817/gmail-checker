package gmail

import (
	"google.golang.org/api/gmail/v1"
)

type Client struct {
	service *gmail.Service
}

func NewClient(service *gmail.Service) *Client {
	return &Client{service: service}
}
