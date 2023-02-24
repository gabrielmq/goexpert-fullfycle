package database_test

import (
	"testing"

	"github.com/gabrielmq/apis/internal/entity"
	"github.com/gabrielmq/apis/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGivenAValidUser_WhenCallsCreate_ThenShouldPersistIt(t *testing.T) {
	// given
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	userDB := database.NewUser(db)

	expectedName := "John Doe"
	expectedEmail := "j@j.com"
	expectedPassword := "123"
	user, _ := entity.NewUser(expectedName, expectedEmail, expectedPassword)

	// when
	actualError := userDB.Create(user)

	// then
	assert.Nil(t, actualError)

	var persistedUser entity.User
	err = db.First(&persistedUser, "id = ?", user.ID).Error
	assert.Nil(t, err)

	assert.Equal(t, user.ID, persistedUser.ID)
	assert.Equal(t, expectedName, persistedUser.Name)
	assert.Equal(t, expectedEmail, persistedUser.Email)
	assert.NotNil(t, persistedUser.Password)
}

func TestGivenAValidEmail_WhenCallsFindByEmail_ThenShouldReturnUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	userDB := database.NewUser(db)

	expectedName := "John Doe"
	expectedEmail := "j@j.com"
	expectedPassword := "123"
	user, _ := entity.NewUser(expectedName, expectedEmail, expectedPassword)

	err = userDB.Create(user)
	assert.Nil(t, err)

	// when
	actualUser, actualError := userDB.FindByEmail(expectedEmail)

	// then
	assert.Nil(t, actualError)
	assert.Equal(t, user.ID, actualUser.ID)
	assert.Equal(t, expectedName, actualUser.Name)
	assert.Equal(t, expectedEmail, actualUser.Email)
	assert.NotNil(t, actualUser.Password)
}

func TestGivenAnInvalidEmail_WhenCallsFindByEmail_ThenShouldReturnError(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	userDB := database.NewUser(db)

	expectedName := "John Doe"
	expectedEmail := "j@j.com"
	expectedPassword := "123"
	user, _ := entity.NewUser(expectedName, expectedEmail, expectedPassword)

	err = userDB.Create(user)
	assert.Nil(t, err)

	// when
	_, actualError := userDB.FindByEmail("teste")

	// then
	assert.NotNil(t, actualError)
}
