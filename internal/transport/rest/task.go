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
	r.readJSON(c, &input)
	in := &proto.Task{
		Title:  input.Title,
		Body:   input.Body,
		UserId: int64(userID.(uint64)),
	}
	resp, err := r.rpcClient.Tasks.Create(context.TODO(), in)
	fmt.Print(resp, 123)
	if err != nil {
		r.sendErrorJSON(c, http.StatusBadRequest, err)
	}
	r.sendJSON(c, resp)
}
