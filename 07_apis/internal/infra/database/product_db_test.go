package database_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gabrielmq/apis/internal/entity"
	"github.com/gabrielmq/apis/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGivenAValidProduct_WhenCallsCreate_ThenShouldPersistIt(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	expectedName := "Product 1"
	expectedPrice := 100.0

	product, err := entity.NewProduct(expectedName, expectedPrice)
	assert.NoError(t, err)

	// when
	actualError := productDB.Create(product)

	// then
	assert.NoError(t, actualError)

	var persistedProduct entity.Product
	err = db.First(&persistedProduct, "id = ?", product.ID).Error
	assert.NoError(t, err)

	assert.Equal(t, product.ID, persistedProduct.ID)
	assert.Equal(t, expectedName, persistedProduct.Name)
	assert.Equal(t, expectedPrice, persistedProduct.Price)
	assert.NotNil(t, persistedProduct.CreatedAt)
}

func TestGivenAValidProductID_WhenCallsFindByID_ThenShouldReturnProduct(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	expectedName := "Product 1"
	expectedPrice := 100.0

	product, err := entity.NewProduct(expectedName, expectedPrice)
	assert.NoError(t, err)

	db.Create(product)

	// when
	actualProduct, _ := productDB.FindByID(product.ID.String())

	// then
	assert.NoError(t, err)

	assert.NotNil(t, actualProduct.ID)
	assert.Equal(t, expectedName, actualProduct.Name)
	assert.Equal(t, expectedPrice, actualProduct.Price)
	assert.NotNil(t, actualProduct.CreatedAt)
}

func TestGivenAnInvalidProductID_WhenCallsFindByID_ThenShouldReturnError(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	expectedName := "Product 1"
	expectedPrice := 100.0

	product, err := entity.NewProduct(expectedName, expectedPrice)
	assert.NoError(t, err)

	db.Create(product)

	// when
	_, actualError := productDB.FindByID("invalid")

	// then
	assert.Error(t, actualError)
}

func TestGivenAValidParams_WhenCallsFindAll_ThenShouldReturnProducts(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	for i := 1; i < 24; i++ {
		expectedName := fmt.Sprintf("Product %d", i)
		expectedPrice := rand.Float64() * 100

		product, err := entity.NewProduct(expectedName, expectedPrice)
		assert.NoError(t, err)

		db.Create(product)
	}

	expectedSort := "asc"
	expectedPaginationSize := 10

	// when
	actualProducts, actualError := productDB.FindAll(1, 10, expectedSort)

	// then
	assert.NoError(t, actualError)
	assert.Len(t, actualProducts, expectedPaginationSize)
	assert.Equal(t, "Product 1", actualProducts[0].Name)
	assert.Equal(t, "Product 10", actualProducts[9].Name)

	actualProducts, actualError = productDB.FindAll(2, 10, expectedSort)

	assert.NoError(t, actualError)
	assert.Len(t, actualProducts, expectedPaginationSize)
	assert.Equal(t, "Product 11", actualProducts[0].Name)
	assert.Equal(t, "Product 20", actualProducts[9].Name)

	actualProducts, actualError = productDB.FindAll(3, 10, expectedSort)

	assert.NoError(t, actualError)
	assert.Len(t, actualProducts, 3)
	assert.Equal(t, "Product 21", actualProducts[0].Name)
	assert.Equal(t, "Product 23", actualProducts[2].Name)
}

func TestGivenAnEmptySort_WhenCallsFindAll_ThenShouldReturnProductsAsc(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	for i := 1; i < 24; i++ {
		expectedName := fmt.Sprintf("Product %d", i)
		expectedPrice := rand.Float64() * 100

		product, err := entity.NewProduct(expectedName, expectedPrice)
		assert.NoError(t, err)

		db.Create(product)
	}

	expectedSort := ""
	expectedPaginationSize := 10

	// when
	actualProducts, actualError := productDB.FindAll(1, 10, expectedSort)

	// then
	assert.NoError(t, actualError)
	assert.Len(t, actualProducts, expectedPaginationSize)
	assert.Equal(t, "Product 1", actualProducts[0].Name)
	assert.Equal(t, "Product 10", actualProducts[9].Name)

	actualProducts, actualError = productDB.FindAll(2, 10, expectedSort)

	assert.NoError(t, actualError)
	assert.Len(t, actualProducts, expectedPaginationSize)
	assert.Equal(t, "Product 11", actualProducts[0].Name)
	assert.Equal(t, "Product 20", actualProducts[9].Name)

	actualProducts, actualError = productDB.FindAll(3, 10, expectedSort)

	assert.NoError(t, actualError)
	assert.Len(t, actualProducts, 3)
	assert.Equal(t, "Product 21", actualProducts[0].Name)
	assert.Equal(t, "Product 23", actualProducts[2].Name)
}

func TestGivenAnInvalidSort_WhenCallsFindAll_ThenShouldReturnProductsAsc(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	for i := 1; i < 24; i++ {
		expectedName := fmt.Sprintf("Product %d", i)
		expectedPrice := rand.Float64() * 100

		product, err := entity.NewProduct(expectedName, expectedPrice)
		assert.NoError(t, err)

		db.Create(product)
	}

	expectedSort := "test"
	expectedPaginationSize := 10

	// when
	actualProducts, actualError := productDB.FindAll(1, 10, expectedSort)

	// then
	assert.NoError(t, actualError)
	assert.Len(t, actualProducts, expectedPaginationSize)
	assert.Equal(t, "Product 1", actualProducts[0].Name)
	assert.Equal(t, "Product 10", actualProducts[9].Name)

	actualProducts, actualError = productDB.FindAll(2, 10, expectedSort)

	assert.NoError(t, actualError)
	assert.Len(t, actualProducts, expectedPaginationSize)
	assert.Equal(t, "Product 11", actualProducts[0].Name)
	assert.Equal(t, "Product 20", actualProducts[9].Name)

	actualProducts, actualError = productDB.FindAll(3, 10, expectedSort)

	assert.NoError(t, actualError)
	assert.Len(t, actualProducts, 3)
	assert.Equal(t, "Product 21", actualProducts[0].Name)
	assert.Equal(t, "Product 23", actualProducts[2].Name)
}

func TestGivenAPageAndLimitZero_WhenCallsFindAll_ThenShouldReturnProductsWithoutPaging(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	for i := 1; i <= 10; i++ {
		expectedName := fmt.Sprintf("Product %d", i)
		expectedPrice := rand.Float64() * 100

		product, err := entity.NewProduct(expectedName, expectedPrice)
		assert.NoError(t, err)

		db.Create(product)
	}

	expectedSort := "asc"
	expectedPaginationSize := 10

	// when
	actualProducts, actualError := productDB.FindAll(0, 0, expectedSort)

	// then
	assert.NoError(t, actualError)
	assert.Len(t, actualProducts, expectedPaginationSize)
	assert.Equal(t, "Product 1", actualProducts[0].Name)
	assert.Equal(t, "Product 10", actualProducts[9].Name)
}

func TestGivenAValidProduct_WhenCallsUpdate_ThenShouldRefreshIt(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	expectedName := "Product 2"
	expectedPrice := 500.0

	product, err := entity.NewProduct("Product 1", 100.0)
	assert.NoError(t, err)

	db.Create(product)

	product.Name = expectedName
	product.Price = expectedPrice

	// when
	actualError := productDB.Update(product)

	// then
	assert.NoError(t, actualError)

	persistedProduct, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, persistedProduct.Name)
	assert.Equal(t, product.Price, persistedProduct.Price)
}

func TestGivenAnInvalidProduct_WhenCallsUpdate_ThenShouldReturnError(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	expectedName := "Product 1"
	expectedPrice := 100.0

	product, err := entity.NewProduct(expectedName, expectedPrice)
	assert.NoError(t, err)

	// when
	actualError := productDB.Update(product)

	// then
	assert.Error(t, actualError)
}

func TestGivenAValidProductID_WhenCallsDelete_ThenShouldDeleteIt(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	expectedName := "Product 2"
	expectedPrice := 500.0

	product, err := entity.NewProduct(expectedName, expectedPrice)
	assert.NoError(t, err)

	db.Create(product)

	// when
	actualError := productDB.Delete(product.ID.String())

	// then
	assert.NoError(t, actualError)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}

func TestGivenAnInvalidProductID_WhenCallsDelete_ThenShouldReturnError(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)

	// when
	actualError := productDB.Delete("invalidID")

	// then
	assert.Error(t, actualError)
}
