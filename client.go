package pandascore

import (
	"github.com/flusaka/pandascore-go/clients"
)

type Client struct {
	Dota2 *clients.Dota2Client
}

func NewClient(accessToken string) *Client {
	return &Client{
		Dota2: clients.NewDota2Client(accessToken),
	}
}
