package rpc

import "github.com/Levap123/task-manager-api-gateway/proto"

type Client struct {
	Auth proto.AuthClient
}

func NewClient(auth proto.AuthClient) *Client {
	return &Client{
		Auth: auth,
	}
}
