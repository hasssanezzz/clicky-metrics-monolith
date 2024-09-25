package repository

import (
	"time"

	"github.com/hasssanezzz/clicky-metrics-monolith/internal/domain"
	"github.com/jmoiron/sqlx"
)

type URLRepository struct {
	db *sqlx.DB
}

func NewURLRepository(db *sqlx.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(url *domain.URL) error {
	query := `INSERT INTO "url" (user_username, short, long, active, created_at, updated_at) 
              VALUES (:user_username, :short, :long, :active, :created_at, :updated_at)`

	url.CreatedAt = time.Now()
	url.UpdatedAt = time.Now()

	_, err := r.db.NamedExec(query, url)
	if err != nil {
		return err
	}
	return nil
}

func (r *URLRepository) GetByShortCode(shortCode string) (*domain.URL, error) {
	var url domain.URL
	query := `SELECT * FROM "url" WHERE short = ?`

	err := r.db.Get(&url, query, shortCode)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *URLRepository) Update(url *domain.URL) error {
	query := `UPDATE "url" 
              SET long = :long, active = :active, updated_at = :updated_at
              WHERE id = :id`

	url.UpdatedAt = time.Now()

	_, err := r.db.NamedExec(query, url)
	if err != nil {
		return err
	}
	return nil
}

func (r *URLRepository) Deactivate(shortCode string) error {
	query := `UPDATE "url" 
              SET active = 0, updated_at = ? 
              WHERE short = ?`

	_, err := r.db.Exec(query, time.Now(), shortCode)
	if err != nil {
		return err
	}
	return nil
}

func (r *URLRepository) GetUserURLs(username string) ([]domain.URL, error) {
	var urls []domain.URL
	query := `SELECT * FROM "url" WHERE user_username = ?`

	err := r.db.Select(&urls, query, username)
	if err != nil {
		return nil, err
	}
	return urls, nil
}
