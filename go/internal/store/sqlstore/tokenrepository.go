package sqlstore

import (
	"errors"
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
	token.IsValid = true

	req := `INSERT INTO tokens(user_id, token, created_at, expired_at) 
			VALUES ($1, $2, $3, $4) RETURNING id
		`
	r.store.LogSql(req, token.User.ID, token.Token, token.CreatedAt, token.ExpiredAt)

	return r.store.db.QueryRow(
		req,
		token.User.ID,
		token.Token,
		token.CreatedAt,
		token.ExpiredAt,
	).Scan(&token.ID)
}

func (r *TokenRepository) Check(tokenStr string) (*models.Token, error) {

	token := &models.Token{
		Token: tokenStr,
	}

	req := `
		SELECT t.id, t.user_id, t.created_at, t.expired_at,
		       u.id, u.username, u.email, u.register_at, u.notice,
		       r.id, r.name, r.level, r.notice 
		FROM tokens as t 
    		LEFT JOIN users as u on u.id = t.user_id 
		    LEFT JOIN roles as r on u.role_id = r.id 
		WHERE t.token=$1`
	r.store.LogSql(req, token.User.ID, token.Token, token.CreatedAt, token.ExpiredAt)

	r.store.db.QueryRow(req, tokenStr).
		Scan(
			&token.ID, &token.User.ID, &token.CreatedAt, &token.ExpiredAt,
			&token.User.ID, &token.User.Username, &token.User.Email, &token.User.RegisterAt, &token.User.Notice,
			&token.User.Role.ID, &token.User.Role.Name, &token.User.Role.Level, &token.User.Role.Notice,
			)

	if token.ID == 0 || token.User.ID == 0 || token.User.Role.ID == 0 {
		return nil, errors.New("invalid token")
	}

	now := time.Now()
	diff := now.Sub(token.ExpiredAt)

	if diff < 0 {
		return nil, errors.New("token expired")
	}

	token.IsValid = true
	return token, nil
}



