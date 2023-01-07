package rest

import (
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

}
