package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Activated int    `gorm:"default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
	Category_ID int
	Activated   int `gorm:"default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	dsn := "root:@/golang_example"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
	fmt.Println(err)
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Product{})

}
