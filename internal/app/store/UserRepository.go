package store

import (
	"github.com/MaximKudryavtsev/go-test/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.sb.QueryRow(
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id",
		user.Email,
		user.Password,
	).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.sb.QueryRow(
		"SELECT id, email, password FROM users WHERE email = $1", email,
	).Scan(&u.ID, &u.Email, &u.Password); err != nil {
		return nil, err
	}

	return u, nil
}
