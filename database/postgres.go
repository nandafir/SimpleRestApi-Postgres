package database

import (
	"anaconda/config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	// "xapiens.id/geosurvey/config"
)

// Postgres ...
type Postgres struct {
	master *sqlx.DB
	auth   *sqlx.DB
}

// New ...
func New() *Postgres {
	// connect geosurvey DB
	config := config.Get()
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			config.GeoSurveyDBUsername, config.GeoSurveyDBPassword,
			config.GeoSurveyDBHost, config.GeoSurveyDB,
		))
	db.SetConnMaxLifetime(time.Minute)

	auth, err := sqlx.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			config.AuthDBUsername, config.AuthDBPassword,
			config.AuthDBHost, config.AuthDB,
		))
	auth.SetConnMaxLifetime(time.Minute)
	if err != nil {
		log.Fatal("connection error")
	}

	return &Postgres{
		master: db,
		auth:   auth,
	}
}

// GetActiveDB ...
func (p *Postgres) GetActiveDB() *sqlx.DB {
	return p.master
}

// GetAuthDB ...
func (p *Postgres) GetAuthDB() *sqlx.DB {
	return p.auth
}
