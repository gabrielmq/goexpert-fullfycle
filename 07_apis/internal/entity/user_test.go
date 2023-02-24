package entity_test

import (
	"testing"

	"github.com/gabrielmq/apis/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAValidParams_WhenCallsNewUser_ThenShouldInstantiateUser(t *testing.T) {
	// given
	expectedName := "John Doe"
	expectedEmail := "j@j.com"
	expectedPassword := "12345"

	// when
	actualUser, err := entity.NewUser(expectedName, expectedEmail, expectedPassword)

	// then
	assert.Nil(t, err)
	assert.NotNil(t, actualUser)
	assert.NotNil(t, actualUser.ID)
	assert.Equal(t, expectedName, actualUser.Name)
	assert.Equal(t, expectedEmail, actualUser.Email)
	assert.NotEmpty(t, actualUser.Password)
}

func TestGivenAValidPassword_WhenCallsValidatePassword_thenShouldBeOk(t *testing.T) {
	// given
	expectedPassword := "12345"

	// when
	actualUser, err := entity.NewUser("John Doe", "j@j.com", expectedPassword)

	// then
	assert.Nil(t, err)
	assert.True(t, actualUser.ValidatePassword(expectedPassword))
}

func TestGivenAnInvalidPassword_WhenCallsValidatePassword_thenShouldBeOk(t *testing.T) {
	// when
	actualUser, err := entity.NewUser("John Doe", "j@j.com", "12345")

	// then
	assert.Nil(t, err)
	assert.False(t, actualUser.ValidatePassword("123"))
}
