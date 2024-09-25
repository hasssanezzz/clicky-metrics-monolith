package repository

import (
	"time"

	"github.com/hasssanezzz/clicky-metrics-monolith/internal/domain"
	"github.com/jmoiron/sqlx"
)

type URLAnalyticsRepository struct {
	db *sqlx.DB
}

func NewURLAnalyticsRepository(db *sqlx.DB) *URLAnalyticsRepository {
	return &URLAnalyticsRepository{db: db}
}

func (r *URLAnalyticsRepository) Create(analytics *domain.URLAnalytics) error {
	query := `INSERT INTO "url_analytics" (url_id, accessed_at, ip_address, user_agent, browser, device, location) 
              VALUES (:url_id, :accessed_at, :ip_address, :user_agent, :browser, :device, :location)`

	analytics.AccessedAt = time.Now()

	_, err := r.db.NamedExec(query, analytics)
	if err != nil {
		return err
	}
	return nil
}

func (r *URLAnalyticsRepository) GetByURL(urlID int) ([]domain.URLAnalytics, error) {
	var analytics []domain.URLAnalytics
	query := `SELECT * FROM "url_analytics" WHERE url_id = ?`

	err := r.db.Select(&analytics, query, urlID)
	if err != nil {
		return nil, err
	}
	return analytics, nil
}

func (r *URLAnalyticsRepository) GetByDateRange(urlID int, startDate, endDate string) ([]domain.URLAnalytics, error) {
	var analytics []domain.URLAnalytics
	query := `SELECT * FROM "url_analytics" 
              WHERE url_id = ? 
              AND accessed_at BETWEEN ? AND ?`

	err := r.db.Select(&analytics, query, urlID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return analytics, nil
}
