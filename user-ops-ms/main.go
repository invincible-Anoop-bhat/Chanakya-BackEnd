package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to user operations microservice")
	router := gin.Default()

	router.GET("/index", Index)

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to User Operations Backend!")
}
