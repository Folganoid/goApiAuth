package sqlstore

import "goApiAuth/go/internal/models"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	return nil,nil
}

func (r *UserRepository) GetById(id int) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	return nil, nil
}



