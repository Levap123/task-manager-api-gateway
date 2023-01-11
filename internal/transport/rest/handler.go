package rest

import (
	"github.com/Levap123/task-manager-api-gateway/internal/client/rpc"
	utils "github.com/Levap123/task-manager-api-gateway/pkg"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	logger    *utils.Logger
	rpcClient *rpc.Client
}

func NewRest(logger *utils.Logger, client *rpc.Client) *Rest {
	return &Rest{
		logger:    logger,
		rpcClient: client,
	}
}
func (r *Rest) InitRoutes() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", r.signIn)
		auth.POST("/sign-up", r.signUp)
		refresh := auth.Group("refresh", r.userIdentity)
		{
			refresh.POST("", r.refresh)
		}

	}
	// api := router.Group("/api/v1", r.userIdentity)
	{

	}
	return router
}
