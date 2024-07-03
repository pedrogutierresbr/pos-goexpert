// Aulas
// Has many usando serial number --> quando o registro pode ter varias opções de outros

// para essa aula, limpe as tableas do banco antes de começar
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{Name: "Cozinha"}
	result := db.Create(&category)
	if result.Error != nil {
		panic(result.Error)
	}

	// create product
	product := Product{
		Name:       "Geladeira",
		Price:      2500.00,
		CategoryID: category.ID,
	}
	result = db.Create(&product)
	if result.Error != nil {
		panic(result.Error)
	}

	// create serial number
	serialNumber := SerialNumber{
		Number:    "123456",
		ProductID: product.ID,
	}
	result = db.Create(&serialNumber)
	if result.Error != nil {
		panic(result.Error)
	}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, "SerialNumber:", product.SerialNumber.Number)
		}
	}
}
