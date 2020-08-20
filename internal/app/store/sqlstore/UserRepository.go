package sqlstore

import (
	"database/sql"
	"github.com/MaximKudryavtsev/go-test/internal/app/model"
	"github.com/MaximKudryavtsev/go-test/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id",
		user.Email,
		user.Password,
	).Scan(&user.ID); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, password FROM users WHERE email = $1", email,
	).Scan(&u.ID, &u.Email, &u.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil
}
