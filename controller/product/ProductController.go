package ProductController

import (
	"fmt"
	"go_example/config"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int
	Name       string
	CategoryID int
	Category   Category
	Activated  int
}
type Category struct {
	ID        int
	Name      string
	Activated int
}

func Create(c *gin.Context) {
	dsn := config.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
	fmt.Println(err)
	var product Product
	c.BindJSON(&product)
	result := db.Create(&product)
	fmt.Println("*****************************Create Product*****************************")
	if result.Error != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": result.Error,
			"request": product,
			"c bind":  &product,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Add Product Success fully",
		})
	}
}
func GetAll(c *gin.Context) {
	dsn := config.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
	fmt.Println(err)
	var product []Product
	result := db.Preload("Category").Find(&product)
	fmt.Println("*****************************GetAll Product*****************************")
	if result.Error != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": result.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  true,
			"message": &product,
		})
	}
}
