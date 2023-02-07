// 01 - GORM Básico
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"` // tags utilizadas pelo gorm
	Name       string
	Price      float64
	gorm.Model // aqui o gorm.Model vai criar os campos created_at, updated_at, deleted_at para gerenciamento do orm
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	// abrindo uma conexão com mysql usando GORM (ORM do go)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// criando um auto migrate para sempre criar automaticamente a tabela product
	db.AutoMigrate(&Product{})

	// inserindo um product
	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.0,
	})

	// criando os produtos em batch
	products := []Product{
		{Name: "Notebook", Price: 1000.0},
		{Name: "Mouse", Price: 10.0},
		{Name: "Keyboard", Price: 200.0},
	}
	db.Create(&products)

	// buscando um produto pelo id
	var product Product
	// db.First(&product, 1)
	// fmt.Println(product)

	// buscando um produto pelo nome
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// buscando todos os produtos
	// var products []Product
	// db.Find(&products)

	// limitando o retorno de registros pela query
	// db.Limit(2).Find(&products)
	// fmt.Println(products)

	// usando offset para ter paginacao, retornando 2 registros por pagina
	// db.Limit(2).Offset(2).Find(&products)
	// fmt.Println(products)

	// buscando todos os produtos que sejam notebook
	// db.Find(&products, "name = ?", "Mouse")
	// fmt.Println(products)

	// buscando usando WHERE
	// db.Where("price > ?", 100.0).Find(&products)
	// fmt.Println(products)

	// buscando usando o LIKE
	// db.Where("name LIKE ?", "%book%").Find(&products)
	// fmt.Println(products)

	// // atualizando o produto com Save
	db.First(&product, 1)
	product.Name = "New Mouse"
	db.Save(&product)
	fmt.Println("atualizado", product)

	// removendo um produto
	db.Delete(&product)
}
