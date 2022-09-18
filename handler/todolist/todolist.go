package todolist

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	CreateTask()
	UpdateTask()
	DeleteTask()
	GetTask()
}

type handler struct {
	group *gin.RouterGroup
}

func (h *handler) CreateTask() {
	//TODO implement me
	const endpoint = `/task/create`
	h.group.POST(endpoint, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Create Endpoint ok"})
	})
}

func (h *handler) UpdateTask() {
	//TODO implement me
	const endpoint = `/task/update`
	h.group.PUT(endpoint, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "update Endpoint ok"})
	})
}

func (h *handler) DeleteTask() {
	//TODO implement me
	const endpoint = `/task/delete`
	h.group.DELETE(endpoint, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "delete Endpoint ok"})
	})
}

func (h *handler) GetTask() {
	//TODO implement me
	const endpoint = `/task/get`
	h.group.GET(endpoint, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "get Endpoint ok"})
	})
}

func NewHandler(group *gin.RouterGroup) Handler {
	h := &handler{
		group: group,
	}
	h.DeleteTask()
	h.CreateTask()
	h.GetTask()
	h.UpdateTask()
	return h
}
