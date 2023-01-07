package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message" binding:"required"`
}

func (r *Rest) sendJSON(c *gin.Context, input interface{}) {
	c.JSON(http.StatusOK, input)
}

func (r *Rest) sendErrorJSON(c *gin.Context, status int, err error) {
	r.logger.Log.Errorln(err)
	c.JSON(status, errorResponse{Message: err.Error()})
}

func (r *Rest) readJSON(c *gin.Context, dest interface{}) error {
	if err := c.ShouldBindJSON(dest); err != nil {
		r.sendErrorJSON(c, http.StatusBadRequest, fmt.Errorf("incorrect input body"))
		return err
	}
	return nil
}
