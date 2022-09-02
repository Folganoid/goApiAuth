package store

import "goApiAuth/go/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetById(id int) (*models.User, error)
	Delete(id int) error
	Update(user *models.User) error
	GetByLoginPass(username, pass string) (models.User, error)
}

type TokenRepository interface {
	Create(token *models.Token) error
	Check(tokenStr string) (*models.Token, error)
}

type RoleRepository interface {
	GetById(id int) (models.Role, error)
}
