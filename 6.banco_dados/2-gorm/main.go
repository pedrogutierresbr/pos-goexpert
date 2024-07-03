// Aulas
// Iniciando com GORM
// Configurando e criando operações
// Realizandp primeira consultas
// Realizando consultas com where
// alterando e removendo registros
// trabalhando com soft delete
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	gorm.Model // traz recursos que são gerenciados pelo gorm (created_at, updated_at, deleted_at.....)
}

func main() {
	// configurado conexão, e adicionado configurações de parse e charset
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// AutoMigrate ele vai rodar e criar o bando de dados respeitando a esturtura que vc passa como referência
	db.AutoMigrate(&Product{})

	// create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	// create batch
	// products := []Product{
	// 	{Name: "TV", Price: 5039.00},
	// 	{Name: "Geladeira", Price: 4509.00},
	// 	{Name: "Fogão", Price: 3039.00},
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Fogão")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// limitando busca a x valores
	// var products []Product
	// db.Limit(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// limitando busca a x valores, como pagination
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 3600).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where com Like
	// var products []Product
	// db.Where("name LIKE ?", "%o%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// alterando dados
	// var p Product
	// db.First(&p, 1) //busquei registro 1
	// p.Name = "Novo Notebook"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)

	// removendo dados --> nessa maneira ele esta trabalhando com osft delete
	// soft delete --> nenhum registro vai ser deletado do db nunca, mas da pra gente pegar os dados do que foi deletado (tipo uma lixeira)
	// você mantem consistência (você remover as infos, mas elas continuam ali)
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)
}
