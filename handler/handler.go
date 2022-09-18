package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/natashameiliana/learnapi/handler/todolist"
)

type Pool struct {
	TodolistHandler todolist.Handler
}

func NewPoolHandler(engine *gin.Engine) Pool {
	return Pool{
		// /kertas/
		TodolistHandler: todolist.NewHandler(engine.Group("paper")),
	}
}
