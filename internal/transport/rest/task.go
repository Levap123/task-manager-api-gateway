package rest

import (
	"context"
	"fmt"
	"github.com/Levap123/task-manager-api-gateway/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type taskBody struct {
	Title string `json:"title,omitempty" binding:"required"`
	Body  string `json:"body,omitempty" binding:"required"`
}

// @Summary create task
// @Tags task
// @Description create task and assign it to user
// @ID create_task
// @Accept  json
// @Produce  json
// @Param input body taskBody true "task title and body"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} proto.TaskHelperBody
// @Failure default {object} errorResponse
// @Router /api/v1/tasks [post]
func (r *Rest) Create(c *gin.Context) {
	userID := c.Value("userId")
	var input taskBody
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

// @Summary get tasks by user id
// @Tags task
// @Description get all tasks that assign to user by him id
// @ID getall_task
// @Accept  json
// @Produce  json
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {array} proto.Task
// @Failure default {object} errorResponse
// @Router /api/v1/tasks [get]
func (r *Rest) GetTasksByUserId(c *gin.Context) {
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

// @Summary get task by id
// @Tags task
// @Description get one task by user id and task id
// @ID get_task
// @Accept  json
// @Produce  json
// @Param id path string true " ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} proto.Task
// @Failure default {object} errorResponse
// @Router /api/v1/tasks/{id} [get]
func (r *Rest) GetTaskByTaskId(c *gin.Context) {
	taskID := c.Param("taskId")
	userID := c.Value("userId")
	fmt.Println(taskID)
	in := &proto.UserAndTask{
		UserId: int64(userID.(uint64)),
		TaskId: taskID,
	}
	resp, err := r.rpcClient.Tasks.Get(context.TODO(), in)
	if err != nil {
		r.sendErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	r.sendJSON(c, resp)
}

func (r *Rest) UpdateTask(c *gin.Context) {
	userID := c.Value("userId")
	var input taskBody
	if err := r.readJSON(c, &input); err != nil {
		return
	}
	taskID := c.Param("taskId")
	in := &proto.Task{
		Id:     taskID,
		Title:  input.Title,
		Body:   input.Body,
		UserId: int64(userID.(uint64)),
	}
	resp, err := r.rpcClient.Tasks.Update(context.TODO(), in)
	if err != nil {
		r.sendErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	r.sendJSON(c, resp)
}

// @Summary delete task by id
// @Tags task
// @Description delete one task by user id and task id
// @ID delete_task
// @Accept  json
// @Produce  json
// @Param id path string true " ID"
// @Param Authorization header string true "With the bearer started" default(Bearer <Add access token here>)
// @Success 200 {object} proto.TaskHelperBody
// @Failure default {object} errorResponse
// @Router /api/v1/tasks/{id} [delete]
func (r *Rest) DeleteTask(c *gin.Context) {
	taskID := c.Param("taskId")
	userID := c.Value("userId")
	in := &proto.UserAndTask{
		UserId: int64(userID.(uint64)),
		TaskId: taskID,
	}

	resp, err := r.rpcClient.Tasks.Delete(context.TODO(), in)
	if err != nil {
		r.sendErrorJSON(c, http.StatusBadRequest, err)
		return
	}
	r.sendJSON(c, resp)
}
