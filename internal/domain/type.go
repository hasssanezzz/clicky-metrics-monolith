package domain

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type URL struct {
	ID           int       `json:"id" db:"id"`
	UserUsername string    `json:"user_username" db:"user_username"`
	Short        string    `json:"short" db:"short"`
	Long         string    `json:"long" db:"long"`
	Active       bool      `json:"active" db:"active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type URLAnalytics struct {
	ID         int       `json:"id" db:"id"`
	URLID      int       `json:"url_id" db:"url_id"`
	AccessedAt time.Time `json:"accessed_at" db:"accessed_at"`
	IPAddress  string    `json:"ip_address" db:"ip_address"`
	UserAgent  string    `json:"user_agent" db:"user_agent"`
	Browser    string    `json:"browser" db:"browser"`
	Device     string    `json:"device" db:"device"`
	Location   string    `json:"location" db:"location"`
}
