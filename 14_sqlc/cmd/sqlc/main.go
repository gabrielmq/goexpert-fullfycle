package main

import (
	"context"
	"database/sql"

	"github.com/gabrielmq/sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Teste",
	// 	Description: sql.NullString{String: "Teste", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "efbef1e2-15f8-4f20-bc11-ca0c8c891ff6",
		Name:        "Teste 2",
		Description: sql.NullString{String: "Teste 2", Valid: true},
	})
	if err != nil {
		panic(err)
	}
	categories, err = queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	err = queries.DeleteCategory(ctx, "efbef1e2-15f8-4f20-bc11-ca0c8c891ff6")
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
