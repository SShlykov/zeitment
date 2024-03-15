package entity

import (
	"github.com/jackc/pgx/v5"
	"time"
)

type UserToken struct {
	ID         string
	UserID     string
	Token      string
	Context    string
	SentTo     string
	InsertedAt time.Time
}

func (ut UserToken) TableName() string {
	return "user_tokens"
}

func (ut UserToken) AllFields() []string {
	return []string{"id", "user_id", "token", "context", "sent_to", "inserted_at"}
}

func (ut UserToken) InsertOrUpdateFields() []string {
	return []string{"user_id", "token", "context", "sent_to", "inserted_at"}
}

func (ut UserToken) EntityToInsertValues(impl *UserToken) []interface{} {
	return []interface{}{
		impl.UserID, impl.Token, impl.Context, impl.SentTo, impl.InsertedAt,
	}
}

func (ut UserToken) ReadItem(row pgx.Row) (UserToken, error) {
	var userToken UserToken
	err := row.Scan(&userToken.ID, &userToken.UserID, &userToken.Token, &userToken.Context, &userToken.SentTo, &userToken.InsertedAt)
	if err != nil {
		return UserToken{}, err
	}
	return userToken, nil
}

func (ut UserToken) ReadList(rows pgx.Rows) ([]UserToken, error) {
	var userTokens []UserToken
	for rows.Next() {
		userToken, err := ut.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		userTokens = append(userTokens, userToken)
	}
	return userTokens, nil
}
