package rest

import (
	"context"

	"github.com/Levap123/task-manager-api-gateway/proto"
	"github.com/gin-gonic/gin"
)

func (r *Rest) signIn(c *gin.Context) {

}

type signUpBody struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *Rest) signUp(c *gin.Context) {
	var input signUpBody
	if err := r.readJSON(c, &input); err != nil {
		return
	}
	user := &proto.User{
		Name:     input.Username,
		Email:    input.Email,
		Password: input.Password,
	}
	r.rpcClient.Auth.SignUp(context.Background(), user)
}
