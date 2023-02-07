// Aula 01 - SQL puro
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // pacote vai ser usado implicitamente, o _ indica para o Go ignorar se nao estivemos utilizando o pacote
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	// estabelecendo uma conexão com o mysql
	// problemas na conexao só serao vistos quando o banco for realmente utilizado
	// durante a conexao nao apareceram erros de conexao
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Notebook", 2000.0)
	if err := insertProduct(db, product); err != nil {
		panic(err)
	}

	product.Price = 1999.99
	if err := updateProduct(db, product); err != nil {
		panic(err)
	}

	product, err = findProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)

	products, err := findProducts(db)
	if err != nil {
		panic(err)
	}

	for _, p := range products {
		if err := deleteProduct(db, p.ID); err != nil {
			panic(err)
		}
	}
}

// sempre deve ser passada uma conexao com o db para serem feitas as consultas
// usar Exec para executar uma acao
// usar Query para buscar

func insertProduct(db *sql.DB, product *Product) error {
	// preprarando a inserção evitando sql injection sanitizando os dados a serem inseridos
	stmt, err := db.Prepare("insert into products(id,name,price) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// executando o insert na base
	// os parametros devem ser passados na ordem das ? do stmt
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	// preparando a atualização evitando sql injection sanitizando os dados a serem inseridos
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// executando o update na base
	// os parametros devem ser passados na ordem das ? do stmt
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func findProduct(db *sql.DB, id string) (*Product, error) {
	// preparando o select evitando sql injection sanitizando o id a ser buscado
	stmt, err := db.Prepare("select * from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product
	// QueryRow é usado para retornar apenas 1 linha da base
	// o Scan atribui o valor de cada coluna retornada para os atributos da struct
	// necessário passar a referencia da memoria
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func findProducts(db *sql.DB) ([]Product, error) {
	// db.Query é utlizado quando a busca retornar mais de um registro da base
	rows, err := db.Query("select * from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	// rows.Next() usado para percorrer as linhas encontradas até não sobrar mais nenhuma
	for rows.Next() {
		var product Product
		// rows.Scan(...) faz o bind dos valores das colunas com os atributos de uma struct
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		// append é utilizado para adicionar itens em um slice
		products = append(products, product)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
