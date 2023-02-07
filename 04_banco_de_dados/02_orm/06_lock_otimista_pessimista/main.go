// 05 - Many To Many
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Product{})

	// cria uma categoria
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	// inicia uma transacao
	tx := db.Begin()
	var c Category
	// usando o lock pessimista para lockar a base para ninguem mais atualizar o dado
	// isso garante que nenhum outro processo concorrente atualize o dado
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}

	// realizando a atualizacao e da um Save para deslockar a linha. Liberando para atualizacao;
	c.Name = "Eletronicos"
	tx.Debug().Save(&c)
	// finaliza transacao comitando
	tx.Commit()
}
