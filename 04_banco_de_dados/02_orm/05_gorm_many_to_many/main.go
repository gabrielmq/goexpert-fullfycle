// 05 - Many To Many
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"` // indica para o gorm que é um relacionamento many to many, e que será necessario a criação de uma tabela intermediaria com os ids das categorias e produtos
}

type Product struct {
	ID         int `gorm:"primaryKey"` // tags utilizadas pelo gorm
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	// abrindo uma conexão com mysql usando GORM (ORM do go)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// criando um auto migrate para sempre criar automaticamente a tabela product e category
	db.AutoMigrate(&Product{}, &Product{})

	// cria uma categoria
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletronicos"}
	db.Create(&category2)

	// cria um produto
	db.Create(&Product{
		Name:       "Fogão",
		Price:      1000.00,
		Categories: []Category{category, category2},
	})

	var categories []Category
	// Carregando os products dentro de categories
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	// percorrendo as categorias com seus produtos
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("-", product.Name, category.Name)
		}
	}
}
