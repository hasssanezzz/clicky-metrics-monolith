package repository

import (
	"errors"
	"time"

	"github.com/hasssanezzz/clicky-metrics-monolith/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *domain.User) error {
	query := `INSERT INTO "user" (username, email, password) 
	VALUES (:username, :email, :password)`

	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByID(id int) (*domain.User, error) {
	var user domain.User
	query := `SELECT * FROM "user" WHERE id = $1`

	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	query := `SELECT * FROM "user" WHERE username = $1`

	err := r.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *domain.User) error {
	query := `UPDATE "user" 
              SET email = :email, password = :password, updated_at = :updated_at 
              WHERE id = :id`

	user.UpdatedAt = time.Now()
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM "user" WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
