package store

import (
	"database/sql"

	_ "github.com/lib/pq" // Анонимный импорт, чтобы не импортировались методы.
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Будет инициализироваться при подключении к БД и отлавливать ошибки
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// При завершении сервера для отключения от БД.
func (s *Store) Close() {
	s.db.Close()
}
