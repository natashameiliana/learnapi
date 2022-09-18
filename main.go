package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natashameiliana/learnapi/handler"
)

func main() {
	// todolist

	// postman req -> handler -> usecase -> repository
	// postman res <-handler <-usecase <- repository

	fmt.Println("Hola")

	engine := gin.New()
	handler.NewPoolHandler(engine)

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
