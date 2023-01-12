package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Levap123/task-manager-api-gateway/proto"
	"github.com/gin-gonic/gin"
)

type signUpBody struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary sign up
// @Tags auth
// @Description register
// @ID register
// @Accept  json
// @Produce  json
// @Param input body signUpBody true "credentials"
// @Success 200 {object} proto.User
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
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
	resp, err := r.rpcClient.Auth.SignUp(context.Background(), user)
	if err != nil {
		r.sendErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	r.sendJSON(c, resp)
}

type signInBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary sign in
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInBody true "credentials"
// @Success 200 {object} proto.SignInResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (r *Rest) signIn(c *gin.Context) {
	var input signInBody
	if err := r.readJSON(c, &input); err != nil {
		return
	}
	credentials := &proto.SignInBody{
		Name:     input.Username,
		Password: input.Password,
	}
	resp, err := r.rpcClient.Auth.SignIn(context.TODO(), credentials)
	if err != nil {
		r.sendErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	r.sendJSON(c, resp)
}

type refreshBody struct {
	RefreshToken string `json:"refresh_token,omitempty" binding:"required"`
}

// @Summary refresh
// @Tags auth
// @Description refresh access token using refresh token
// @ID refresh
// @Accept  json
// @Produce  json
// @Param input body refreshBody true "credentials"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} refreshBody
// @Failure default {object} errorResponse
// @Router /auth/refresh [post]
func (r *Rest) refresh(c *gin.Context) {
	errTokenExpired := errors.New("token expired")
	var input refreshBody
	fmt.Println(input.RefreshToken)
	if err := r.readJSON(c, &input); err != nil {
		return
	}
	request := &proto.Tokens{
		Acces:   c.Value("access").(string),
		Refresh: input.RefreshToken,
	}
	tokenPair, err := r.rpcClient.Auth.Refresh(context.TODO(), request)
	if err != nil {
		if err == errTokenExpired {
			r.sendErrorJSON(c, http.StatusUnauthorized, fmt.Errorf("refresh token expired lol"))
			return
		}
		r.sendErrorJSON(c, http.StatusInternalServerError, fmt.Errorf("invalid refresh token"))
		return
	}
	r.sendJSON(c, tokenPair)
}
