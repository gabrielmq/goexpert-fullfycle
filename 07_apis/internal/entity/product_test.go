package entity_test

import (
	"testing"

	pkg "github.com/gabrielmq/apis/pkg/entity"

	"github.com/gabrielmq/apis/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAValidParams_WhenCallsNewProduct_ThenShouldInstantiateProduct(t *testing.T) {
	// given
	expectedName := "Product 1"
	expectedPrice := 100.0

	// when
	actualProduct, err := entity.NewProduct(expectedName, expectedPrice)

	// then
	assert.Nil(t, err)
 
	assert.NotNil(t, actualProduct.ID)
	assert.Equal(t, expectedName, actualProduct.Name)
	assert.Equal(t, expectedPrice, actualProduct.Price)
	assert.NotNil(t, actualProduct.CreatedAt)
}

func TestGivenAEmptyName_WhenCallsNewProduct_ThenShouldReturnError(t *testing.T) {
	// given
	expectedName := ""
	expectedPrice := 100.0

	expectedErrorMessage := "name is required"

	// when
	actualProduct, actualError := entity.NewProduct(expectedName, expectedPrice)

	// then
	assert.Nil(t, actualProduct)
	assert.NotNil(t, actualError)
	assert.Equal(t, expectedErrorMessage, actualError.Error())
}

func TestGivenAPriceZero_WhenCallsNewProduct_ThenShouldReturnError(t *testing.T) {
	// given
	expectedName := "Product 1"
	expectedPrice := 0.0

	expectedErrorMessage := "price is required"

	// when
	actualProduct, actualError := entity.NewProduct(expectedName, expectedPrice)
	
	// then
	assert.Nil(t, actualProduct)
	assert.NotNil(t, actualError)
	assert.Equal(t, expectedErrorMessage, actualError.Error())
}

func TestGivenANegativePrice_WhenCallsNewProduct_ThenShouldReturnError(t *testing.T) {
	// given
	expectedName := "Product 1"
	expectedPrice := -10.0

	expectedErrorMessage := "price is invalid"

	// when
	actualProduct, actualError := entity.NewProduct(expectedName, expectedPrice)

	// then
	assert.Nil(t, actualProduct)
	assert.NotNil(t, actualError)
	assert.Equal(t, expectedErrorMessage, actualError.Error())
}

func TestGivenAValidProduct_WhenCallsValidate_ThenShouldNotReturnError(t *testing.T) {
	// given
	expectedID := pkg.NewID()
	expectedName := "Product 1"
	expectedPrice := 100.0

	product := entity.Product{
		ID:    expectedID,
		Name:  expectedName,
		Price: expectedPrice,
	}

	// when
	actualError := product.Validate()

	// then
	assert.Nil(t, actualError)
}
