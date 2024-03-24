package dto

import (
	"github.com/SShlykov/zeitment/auth/internal/models/types"
	"time"
)

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt types.Null[time.Time]

	LoggedAt    types.Null[time.Time]
	ConfirmedAt types.Null[time.Time]

	Login string             // index
	Email types.Null[string] // index

	PasswordHash string

	DeletedBy        types.Null[string]
	AccessTemplateID types.Null[int]
	UpdateAfter      types.Null[int64]
}
