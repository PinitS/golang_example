package main

import (
	"fmt"
	"go_example/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type IntegerType struct {
// 	Int    int
// 	Int8   int8
// 	Int16  int16
// 	Int32  int32
// 	Int64  int64
// 	Uint   uint
// 	Uint8  uint8
// 	Uint16 uint16
// 	Uint32 uint32
// 	Uint64 uint64
// }

// type StringType struct {
// 	Text    string `gorm:"type:text"`
// 	Varchar string `gorm:"type:varchar(100)"`
// 	Char    string `gorm:"type:char"`
// }

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Activated int    `gorm:"default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"unique"`
	CategoryID int
	Category   Category
	Activated  int `gorm:"default:1"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// type User struct {
// 	gorm.Model
// 	Name      string
// 	CompanyID int
// 	Company   Company
// }

// type Company struct {
// 	ID   int
// 	Name string
// }

func main() {
	dsn := config.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
	fmt.Println(err)
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Product{})
	// db.AutoMigrate(&IntegerType{})
	// db.AutoMigrate(&StringType{})
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Company{})
}
