package sqlstore

import (
	"errors"
	"goApiAuth/go/internal/models"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *models.User) error {

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	role, err := r.store.Role().GetById(user.Role.ID)
	if err != nil {
		return err
	}

	user.Role = role

	return r.store.db.QueryRow(
		`INSERT INTO users(username, email, hash_password, register_at, role_id, notice) 
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
		`,
		user.Username,
		user.Email,
		user.HashPassword,
		user.RegisterAt,
		user.Role.ID,
		user.Notice,
	).Scan(&user.ID)

}

func (r *UserRepository) GetById(id int) (*models.User, error) {

	user := &models.User{ID: id}
	r.store.db.QueryRow(`SELECT username, email, register_at, role_id, notice FROM users WHERE id=$1`, id).Scan(&user.Username, &user.Email, &user.RegisterAt, &user.Role.ID, &user.Notice)

	if user.Email == "" {
		return nil, errors.New("Can not define user")
	}

	role, err := r.store.Role().GetById(user.Role.ID)
	if err != nil {
		return nil, err
	}

	user.Role = role
	return user, nil

}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {

	user := &models.User{Email: email}
	r.store.db.QueryRow(`SELECT id, username, register_at, role_id, notice FROM users WHERE email=$1`, email).Scan(&user.ID, &user.Username, &user.RegisterAt, &user.Role.ID, &user.Notice)

	if user.Email == "" {
		return nil, errors.New("Can not define user")
	}

	role, err := r.store.Role().GetById(user.Role.ID)
	if err != nil {
		return nil, err
	}

	user.Role = role
	return user, nil
}



