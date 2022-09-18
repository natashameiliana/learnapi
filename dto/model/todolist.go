package model

import (
	"github.com/natashameiliana/learnapi/dto/entity"
	"github.com/natashameiliana/learnapi/dto/standard"
)

// creation task request response object
type (
	CreateTaskRequest struct {
		UserID   string `json:"user_id"`
		Title    string `json:"title"`
		Desc     string `json:"desc"`
		Done     bool   `json:"done"`
		DueDate  string `json:"dueDate"`
		Category string `json:"category"`
		Priority int    `json:"prioroty"`
	}
	CreateTaskResponse struct {
		standard.Response
	}
)

// modify task request response object
type (
	UpdateTaskRequest struct {
		Title    string `json:"title"`
		Desc     string `json:"desc"`
		Done     bool   `json:"done"`
		DueDate  string `json:"dueDate"`
		Category string `json:"category"`
		Priority int    `json:"prioroty"`
	}
	UpdateTaskResponse struct {
		standard.Response
	}
)

// Deletion task request response object
type (
	DeleteTaskRequest struct {
		ID string `json:"delete"`
	}
	DeleteTaskResponse struct {
		standard.Response
	}
)

// retrival task request response object
type (
	GetTaskRequest struct {
		UserID string `json:"user_id"`
	}
	GetTaskResponse struct {
		standard.Response
		List []entity.Todolist `json:"list"`
	}
)
