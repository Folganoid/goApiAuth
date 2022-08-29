package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"goApiAuth/go/internal/store"
)

type Store struct {
	db *sql.DB
	userRepository *UserRepository
	tokenRepository *TokenRepository
	roleRepository *RoleRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Token() store.TokenRepository {
	if s.tokenRepository != nil {
		return s.tokenRepository
	}
	s.tokenRepository = &TokenRepository{
		store: s,
	}

	return s.tokenRepository
}

func (s *Store) Role() store.RoleRepository {
	if s.roleRepository != nil {
		return s.roleRepository
	}
	s.roleRepository = &RoleRepository{
		store: s,
	}

	return s.roleRepository
}