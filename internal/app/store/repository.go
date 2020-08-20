package store

import "github.com/MaximKudryavtsev/go-test/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}
