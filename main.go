package main

import (
	"fmt"
	CategoryController "go_example/controller/category"
	ProductController "go_example/controller/product"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Tent-Official")
	r := gin.Default()
	r.POST("/categoryCreate", CategoryController.Create)
	r.GET("/categoryGet", CategoryController.GetAll)
	r.POST("/categoryUpdate", CategoryController.Update)
	r.POST("/categoryDelete", CategoryController.Delete)

	r.POST("/productCreate", ProductController.Create)
	r.GET("/productGet", ProductController.GetAll)
	r.Run()
}
