package sqlstore

import (
	"REST_Api/store"
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// func (s *Store) Open() error {
// 	db, err := sql.Open("postgres", s.config.DatabaseURL)
// 	if err != nil {
// 		return err
// 	}
// 	if err := db.Ping(); err != nil {
// 		return err
// 	}

// 	s.db = db
// 	return nil
// }

// func (s *Store) Close() error {
// 	s.db.Close()

// 	return nil
// }

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
