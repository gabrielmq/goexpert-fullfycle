package entity

import (
	"github.com/gabrielmq/apis/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"` // id como VO (Value Object)
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	// gerando um hash da senha informada
	// bcrypt lib go pra criar hashs para senhas
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash), // senha sempre deve ser um hash
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	// validando as senhas para saber se o hash gerado é igual
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
