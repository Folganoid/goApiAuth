package sqlstore

import "goApiAuth/go/internal/models"

type TokenRepository struct {
	store *Store
}

func (r *TokenRepository) Create(token *models.Token) (*models.Token, error) {
	return nil, nil
}

func (r *TokenRepository) CheckToken(token string) (bool, error) {
	return true,nil
}
func (r *TokenRepository) DeleteByToken(token string) error {
	return nil
}
func (r *TokenRepository) DeleteByUserId(userId int) error {
	return nil
}



