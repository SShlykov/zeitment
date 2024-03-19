package entity

import "github.com/jackc/pgx/v5"

type Role struct {
	ID   string
	Name string
}

func (r Role) TableName() string {
	return "roles"
}

func (r Role) AllFields() []string {
	return []string{"id", "name"}
}

func (r Role) InsertOrUpdateFields() []string {
	return []string{"name"}
}

func (r Role) EntityToInsertValues(impl *Role) []interface{} {
	return []interface{}{
		impl.Name,
	}
}

func (r Role) ReadItem(row pgx.Row) (Role, error) {
	var role Role
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return Role{}, err
	}
	return role, nil
}

func (r Role) ReadList(rows pgx.Rows) ([]Role, error) {
	var roles []Role
	for rows.Next() {
		role, err := r.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
