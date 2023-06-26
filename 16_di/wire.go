//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/gabrielmq/di/product"
	"github.com/google/wire"
)

var setRepositoryDependecy = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependecy,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
