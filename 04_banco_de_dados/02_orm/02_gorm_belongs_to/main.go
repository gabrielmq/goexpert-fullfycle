// 02 - Belongs To
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
	ID         int `gorm:"primaryKey"` // tags utilizadas pelo gorm
	Name       string
	Price      float64
	CategoryID int      // faz a relação com a categoria indiretamente (belongsTo)
	Category   Category // alem do id é necessário adicionar a struct que faz o relacionamento
	gorm.Model          // aqui o gorm.Model vai criar os campos created_at, updated_at, deleted_at para gerenciamento do orm
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	// abrindo uma conexão com mysql usando GORM (ORM do go)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// criando um auto migrate para sempre criar automaticamente a tabela product e category
	db.AutoMigrate(&Product{}, &Category{})

	// cria uma categoria
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// cria um produto
	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: category.ID, // criando o relacionamento do produto com a categoria
	})

	var products []Product
	// o Preload indica para o gorm preencher os relacionamentos que product tem durante a consulta
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
}
