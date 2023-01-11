package rest

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/Levap123/task-manager-api-gateway/proto"
	"github.com/gin-gonic/gin"
)

func (r *Rest) userIdentity(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	authParts := strings.Split(auth, "Bearer ")
	if len(authParts) != 2 {
		r.sendErrorJSON(c, http.StatusUnauthorized, errors.New("invalid auth token"))
		return
	}
	request := &proto.Access{
		Access: authParts[1],
	}
	userId, err := r.rpcClient.Auth.Validate(context.TODO(), request)
	if err != nil {
		r.sendErrorJSON(c, http.StatusUnauthorized, errors.New("invalid auth token"))
		return
	}
	c.Set("userId", userId.Id)
}
