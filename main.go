package main

import (
	"fmt"
	"go_example/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi golang Hell")
	r := gin.Default()
	r.POST("/create", controller.Create)
	r.GET("/get", controller.GetAll)
	r.POST("/update", controller.Update)
	r.POST("/delete", controller.Delete)
	r.Run()
}
