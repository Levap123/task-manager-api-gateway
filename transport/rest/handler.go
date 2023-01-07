package rest

import (
	utils "github.com/Levap123/task-manager-api-gateway/pkg"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	logger *utils.Logger
}

func NewRest(logger *utils.Logger) *Rest {
	return &Rest{
		logger: logger,
	}
}
func (r *Rest) InitRoutes() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", r.signIn)
		auth.POST("/sign-up", r.signUp)
		auth.POST("/refresh")
	}
	return router
}
