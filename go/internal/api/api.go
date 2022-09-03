package api

import (
	"database/sql"
	"goApiAuth/go/internal/store/sqlstore"
	"net/http"
	"os"
)

func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	logFile, err := os.OpenFile(config.LogFile, os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer logFile.Close()

	logger, _ := NewLogger(logFile, config.LogLevel)

	store := sqlstore.New(db, logger)
	srv := newServer(store, logger)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}