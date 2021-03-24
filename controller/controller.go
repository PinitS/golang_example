package controller

import (
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  true,
		"message": "create",
	})
}
func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  true,
		"message": "get",
	})
}
func Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  true,
		"message": "update",
	})
}
func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  true,
		"message": "delete",
	})
}
