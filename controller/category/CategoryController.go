package CategoryController

import (
	"fmt"
	"go_example/config"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID        int
	Name      string
	Product   []Product
	Activated int
}
type Product struct {
	ID         int
	Name       string
	CategoryID int
	Activated  int
}

func Create(c *gin.Context) {
	fmt.Println("**********************************************************")

	dsn := config.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
	fmt.Println(err)
	var category Category
	c.BindJSON(&category)
	result := db.Create(&category)
	fmt.Println("**********************************************************")
	if result.Error != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": result.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Add Category Success fully",
		})
	}
}
func GetAll(c *gin.Context) {
	dsn := config.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
	fmt.Println(err)
	var category []Category
	result := db.Preload("Product").Find(&category)
	fmt.Println("**********************************************************")
	if result.Error != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": result.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  true,
			"message": &category,
		})
	}
}
func Update(c *gin.Context) {
	dsn := config.Dsn()
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var category Category
	var api Category
	c.BindJSON(&api)
	result := db.First(&category, &api.ID)
	fmt.Println("**********************************************************")
	fmt.Println("**********************************************************")
	fmt.Println(result.Error)
	fmt.Println("**********************************************************")
	fmt.Println("**********************************************************")
	if result.Error != nil {
		fmt.Println("has error")
		c.JSON(200, gin.H{
			"status":  false,
			"message": "record not found",
		})
	} else {
		fmt.Println("not error")
		category.Name = api.Name
		category.Activated = api.Activated
		db.Save(&category)
		c.JSON(200, gin.H{
			"status":  true,
			"message": &category,
		})
	}
}
func Delete(c *gin.Context) {
	dsn := config.Dsn()
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var category Category
	var api Category
	c.BindJSON(&api)
	result := db.Delete(&category, &api.ID)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"status":  false,
			"message": result.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Delete Category Success fully",
		})
	}
}
