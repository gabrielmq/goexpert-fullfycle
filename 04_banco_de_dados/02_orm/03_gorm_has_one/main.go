// 03 - Has One
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primaryKey"` // tags utilizadas pelo gorm
	Name         string
	Price        float64
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
	db.AutoMigrate(&Product{}, &SerialNumber{})

	// cria um produto
	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.00,
	})

	// cria um serial number
	db.Create(&SerialNumber{
		Number:    "123445",
		ProductID: 1,
	})

	var products []Product
	// o Preload indica para o gorm preencher os relacionamentos que product tem durante a consulta
	db.Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.SerialNumber.Number)
	}
}
