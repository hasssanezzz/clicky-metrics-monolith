package domain

type UserRepository interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(id int) error
}

type URLRepository interface {
	Create(url *URL) error
	GetByShortCode(shortCode string) (*URL, error)
	Update(url *URL) error
	Deactivate(shortCode string) error
	GetUserURLs(username string) ([]URL, error)
}

type URLAnalyticsRepository interface {
	Create(analytics *URLAnalytics) error
	GetByURL(urlID int) ([]URLAnalytics, error)
	GetByDateRange(urlID int, startDate, endDate string) ([]URLAnalytics, error)
}
