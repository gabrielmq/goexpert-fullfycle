// 04 - Has Many
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product // gorm vai perceber automaticamente que a categoria pode ter mais de 1 produto criando relacionamento 1:n
}

type Product struct {
	ID           int `gorm:"primaryKey"` // tags utilizadas pelo gorm
	Name         string
	Price        float64
	CategoryID   int          // faz a relação com a categoria indiretamente (belongsTo)
	Category     Category     // alem do id é necessário adicionar a struct que faz o relacionamento
	SerialNumber SerialNumber // fazendo uma relação 1:1
	gorm.Model                // aqui o gorm.Model vai criar os campos created_at, updated_at, deleted_at para gerenciamento do orm
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int // vai ser obrigatório um productId
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	// abrindo uma conexão com mysql usando GORM (ORM do go)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// criando um auto migrate para sempre criar automaticamente a tabela product e category
	db.AutoMigrate(&Product{}, &Product{}, &SerialNumber{})

	// cria uma categoria
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	// cria um produto
	db.Create(&Product{
		Name:       "Fogão",
		Price:      1000.00,
		CategoryID: category.ID, // criando o relacionamento do produto com a categoria
	})

	// cria um serial number
	db.Create(&SerialNumber{
		Number:    "12345",
		ProductID: 1,
	})

	var categories []Category
	// Carregando os products e seus serial numbers dentro de categories
	// Preload("Products.SerialNumber") o gorm entende que tem que buscar todos os produtos e os serial numbers desses produtos
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	// percorrendo as categorias com seus produtos
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("-", product.Name, category.Name, product.SerialNumber.Number)
		}
	}
}
