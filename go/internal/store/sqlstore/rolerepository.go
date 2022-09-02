package sqlstore

import (
	"errors"
	"goApiAuth/go/internal/models"
)

type RoleRepository struct {
	store *Store
}

func (r *RoleRepository) GetById(id int) (models.Role, error) {

	role := models.Role{ID: id}
	sqlStr := "SELECT name, level, notice FROM roles WHERE id=$1"

	req := r.store.db.QueryRow(sqlStr, id)
	req.Scan(&role.Name, &role.Level, &role.Notice)


	if role.Level == 0 || role.Name == "" {
		return role, errors.New("Can not define user role")
	}

	return role, nil

}
