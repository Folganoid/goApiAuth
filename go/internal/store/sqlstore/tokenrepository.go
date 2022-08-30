package sqlstore

import "goApiAuth/go/internal/models"

type TokenRepository struct {
	store *Store
}

func (r *TokenRepository) Create(token string) (*models.Token, error) {
	return nil, nil
}

func (r *TokenRepository) Check(token string) (*models.Token, error) {
	return nil, nil
}



