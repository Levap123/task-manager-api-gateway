package rest

import (
	_ "github.com/Levap123/task-manager-api-gateway/docs"
	"github.com/Levap123/task-manager-api-gateway/internal/client/rpc"
	utils "github.com/Levap123/task-manager-api-gateway/pkg"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", r.signIn)
		auth.POST("/sign-up", r.signUp)
		auth.POST("/refresh", r.userIdentity, r.refresh)

	}
	api := router.Group("/api/v1", r.userIdentity)
	{
		tasks := api.Group("/tasks")
		{
			tasks.POST("", r.Create)
			tasks.GET("", r.GetTasksByUserId)
			tasks.GET(":taskId", r.GetTaskByTaskId)
			tasks.PUT(":taskId", r.UpdateTask)
		}
	}

	return router
}
