package sqlstore

import (
	"errors"
	"fmt"
	"goApiAuth/go/internal/models"
	"time"
)

type TokenRepository struct {
	store *Store
}

func (r *TokenRepository) Create(token *models.Token) error {

	u, err := r.store.User().GetByLoginPass(token.User.Username, token.User.Password)
	if err != nil {
		return err
	}

	if u.Email == "" {
		return errors.New("Can not define user")
	}

	token.CreatedAt = time.Now()
	token.ExpiredAt = token.CreatedAt.Add(time.Hour * 24)
	token.Token = models.RandomString(128)
	token.User = u

	fmt.Println(u)
	fmt.Println(token)

	return r.store.db.QueryRow(
		`INSERT INTO tokens(user_id, token, created_at, expired_at) 
			VALUES ($1, $2, $3, $4) RETURNING id
		`,
		token.User.ID,
		token.Token,
		token.CreatedAt,
		token.ExpiredAt,
	).Scan(&token.ID)
}

func (r *TokenRepository) Check(token string) (*models.Token, error) {
	return nil, nil
}



