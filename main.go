package main

import (
	"go_example/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/create", controller.Create)
	r.GET("/get", controller.GetAll)
	r.POST("/update", controller.Update)
	r.POST("/delete", controller.Delete)
	r.Run()
}
