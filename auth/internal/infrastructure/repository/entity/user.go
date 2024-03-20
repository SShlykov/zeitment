package entity

import (
	"database/sql"
	"github.com/jackc/pgx/v5"
	"time"
)

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.Null[time.Time]

	LoggedAt    sql.Null[time.Time]
	ConfirmedAt sql.Null[time.Time]

	Login string           // index
	Email sql.Null[string] // index

	PasswordHash string

	DeletedBy        sql.Null[string]
	AccessTemplateID sql.Null[int]
	UpdateAfter      sql.Null[int64]
}

func (u User) TableName() string {
	return "users"
}

func (u User) AllFields() []string {
	return []string{"id", "created_at", "updated_at", "deleted_at", "logged_at", "confirmed_at", "login",
		"email", "password_hash", "deleted_by", "access_template_id", "update_after"}
}

func (u User) InsertOrUpdateFields() []string {
	return []string{"logged_at", "confirmed_at", "login", "email", "password_hash", "deleted_by",
		"access_template_id", "update_after"}
}

func (u User) EntityToInsertValues(impl *User) []interface{} {
	return []interface{}{
		impl.LoggedAt, impl.ConfirmedAt, impl.Login, impl.Email, impl.PasswordHash, impl.DeletedBy,
		impl.AccessTemplateID, impl.UpdateAfter,
	}
}

func (u User) ReadItem(row pgx.Row) (User, error) {
	var user User
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.LoggedAt, &user.ConfirmedAt,
		&user.Login, &user.Email, &user.PasswordHash, &user.DeletedBy, &user.AccessTemplateID, &user.UpdateAfter)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u User) ReadList(rows pgx.Rows) ([]User, error) {
	var users []User
	for rows.Next() {
		user, err := u.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
