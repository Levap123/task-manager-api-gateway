package rpc

import "github.com/Levap123/task-manager-api-gateway/proto"

type Client struct {
	Auth  proto.AuthClient
	Tasks proto.TaskManagerClient
}

func NewClient(auth proto.AuthClient, task proto.TaskManagerClient) *Client {
	return &Client{
		Auth:  auth,
		Tasks: task,
	}
}
