package sqlstore

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
)

// UserRepository - структура для репозитория для БД
type UserRepository struct {
	store *Store
}

// Create создает пользователя
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	return r.store.connection.QueryRow(
		context.Background(),
		"insert into users (email, encrypted_password) values ($1, $2) returning id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

// FindByEmail ищет пользователя по email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	row := r.store.connection.QueryRow(
		context.Background(),
		"select * from users where email=$1",
		email,
	)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Email, &user.EncryptedPassword)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, store.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
