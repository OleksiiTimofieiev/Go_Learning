package teststore

import (
	"REST_Api/model"
	"REST_Api/store"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
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
		fmt.Println("here")
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
