package entity

import (
	"github.com/jackc/pgx/v5"
	"time"
)

type UserRole struct {
	ID        string
	UserID    string
	RoleID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ur UserRole) TableName() string {
	return "user_roles"
}

func (ur UserRole) AllFields() []string {
	return []string{"id", "user_id", "role_id", "created_at", "updated_at"}
}

func (ur UserRole) InsertOrUpdateFields() []string {
	return []string{"user_id", "role_id", "created_at", "updated_at"}
}

func (ur UserRole) EntityToInsertValues(impl *UserRole) []interface{} {
	return []interface{}{
		impl.UserID, impl.RoleID, impl.CreatedAt, impl.UpdatedAt,
	}
}

func (ur UserRole) ReadItem(row pgx.Row) (UserRole, error) {
	var userRole UserRole
	err := row.Scan(&userRole.ID, &userRole.UserID, &userRole.RoleID, &userRole.CreatedAt, &userRole.UpdatedAt)
	if err != nil {
		return UserRole{}, err
	}
	return userRole, nil
}

func (ur UserRole) ReadList(rows pgx.Rows) ([]UserRole, error) {
	var userRoles []UserRole
	for rows.Next() {
		userRole, err := ur.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		userRoles = append(userRoles, userRole)
	}
	return userRoles, nil
}
