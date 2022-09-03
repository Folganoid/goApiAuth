package sqlstore

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"goApiAuth/go/internal/store"
	log "github.com/sirupsen/logrus"
)

type Store struct {
	db *sql.DB
	logger *log.Logger
	logSql func()
	userRepository *UserRepository
	tokenRepository *TokenRepository
	roleRepository *RoleRepository
}

func New(db *sql.DB, logger *log.Logger) *Store {
	return &Store{
		db: db,
		logger: logger,
	}
}

func (s *Store) LogSql (sql string, args ...interface{} ) {

	argsStr := ""
	for _, v := range args {
		argsStr += fmt.Sprintf(" %v; ", v)
	}
	s.logger.Debug("SQL request: " + sql + "with args: " + argsStr)
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