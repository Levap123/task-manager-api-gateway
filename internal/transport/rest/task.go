package rest

import (
	"context"
	"fmt"
	"github.com/Levap123/task-manager-api-gateway/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskBody struct {
	Title string `json:"title,omitempty" binding:"required"`
	Body  string `json:"body,omitempty" binding:"required"`
}

func (r *Rest) Create(c *gin.Context) {
	userID := c.Value("userId")
	var input TaskBody
	if err := r.readJSON(c, &input); err != nil {
		return
	}
	in := &proto.Task{
		Title:  input.Title,
		Body:   input.Body,
		UserId: int64(userID.(uint64)),
	}
	resp, err := r.rpcClient.Tasks.Create(context.TODO(), in)
	fmt.Print(resp, 123)
	if err != nil {
		r.sendErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	r.sendJSON(c, resp)
}

func (r *Rest) GetUserById(c *gin.Context) {
	userID := c.Value("userId")
	in := &proto.UserRequest{
		Id: int64(userID.(uint64)),
	}
	resp, err := r.rpcClient.Tasks.GetAll(context.TODO(), in)
	if err != nil {
		r.sendErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	r.sendJSON(c, resp)

}
