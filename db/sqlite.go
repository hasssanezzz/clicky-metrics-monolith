package sqlite

import (
	"log"

	"github.com/hasssanezzz/clicky-metrics-monolith/bootstrap"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func MustSetup(env *bootstrap.Env) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", env.DBHost)
	if err != nil {
		log.Fatalln("Failed to open SQLite database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln("Failed to connect to SQLite database:", err)
	}

	return db
}
