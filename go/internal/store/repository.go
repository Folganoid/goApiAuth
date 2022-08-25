package store

import "goApiAuth/go/internal/models"

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetById(id int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type TokenRepository interface {
	Create(token *models.Token) (*models.Token, error)
	CheckToken(token string) (bool, error)
	DeleteByToken(token string) error
	DeleteByUserId(userId int) error
}
