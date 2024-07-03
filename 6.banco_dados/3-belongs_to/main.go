// Aulas
// Belongs to --> correlacionar tabelas (ex.: produto pertence a uma categoria, essa catergoria pode ter mais produtos)

// para essa aula, limpe as tableas do banco antes de começar
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// // create category
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// // create product --> primeirp esse, ai comenta
	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      1000.00,
	// 	CategoryID: 1,
	// })

	// // create product --- depois esse
	// db.Create(&Product{
	// 	Name:       "Mouse",
	// 	Price:      9.99,
	// 	CategoryID: 1,
	// })

	var products []Product
	// Preload --> vai preencher o categoryID com o category, faz o relacionamento belongs to
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
}
