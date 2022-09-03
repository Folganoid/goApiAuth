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

	query := `INSERT INTO users(username, email, hash_password, register_at, role_id, notice) 
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
		`

	r.store.LogSql(query, user.Username, user.Email, user.HashPassword, user.RegisterAt, user.Role.ID, user.Notice)

	return r.store.db.QueryRow(query,
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

	req := `SELECT username, email, register_at, role_id, notice FROM users WHERE id=$1`
	r.store.LogSql(req, id)
	r.store.db.QueryRow(req, id).Scan(&user.Username, &user.Email, &user.RegisterAt, &user.Role.ID, &user.Notice)

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

func (r *UserRepository) Delete(id int) error {

	req := `DELETE FROM users WHERE id=$1`
	r.store.LogSql(req, id)
	res, err := r.store.db.Exec(req, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("User was not deleted")
	}

	return nil
}

func (r *UserRepository) Update(user *models.User) error {

	if err := user.BeforeUpdate(); err != nil {
		return err
	}

	role, err := r.store.Role().GetById(user.Role.ID)
		if err != nil {
			return err
		}
	user.Role = role

	req := `
		UPDATE users 
		SET username = $2,
		    email = $3,
		    hash_password = $4,
		    role_id = $5, 
		    notice = $6
		WHERE id = $1;
		`
	r.store.LogSql(req, user.ID, user.Username, user.Email, user.HashPassword, user.Role.ID, user.Notice)

	res, err := r.store.db.Exec(req, user.ID, user.Username, user.Email, user.HashPassword, user.Role.ID, user.Notice)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("User was not updated")
	}

	return nil
}

func (r *UserRepository) GetByLoginPass(username, pass string) (models.User, error) {

	user := models.User{Username: username, Password: pass}
	user.HashPassword = models.GetMD5Hash(pass)
	user.Password = ""

	req := `SELECT id, username, email, register_at, role_id, notice FROM users WHERE username=$1 and hash_password=$2`
	r.store.LogSql(req, user.Username, user.HashPassword)

	r.store.db.QueryRow(`SELECT id, username, email, register_at, role_id, notice FROM users WHERE username=$1 and hash_password=$2`, user.Username, user.HashPassword).
		Scan(&user.ID, &user.Username, &user.Email, &user.RegisterAt, &user.Role.ID, &user.Notice)

	if user.Email == "" {
		return user, errors.New("Can not define user")
	}

	return user, nil
}



