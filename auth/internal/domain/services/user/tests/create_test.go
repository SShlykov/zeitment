package tests

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/domain/services/user"
	"github.com/SShlykov/zeitment/auth/internal/models/dto"
	"github.com/SShlykov/zeitment/auth/internal/test/mocks/repository/mocks"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

type CreateArgs struct {
	ctx      context.Context
	user     *dto.User
	password string
}

type TestCreateStruct struct {
	name     string
	args     CreateArgs
	mock     repoMockFunc
	expected *dto.User
	err      error
}

func TestCreate(t *testing.T) {
	t.Parallel()

	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)

		userDto       = &dto.User{ID: gofakeit.UUID(), Login: "login"}
		validPassword = "Password123"
		//entityUser    = adapters.UserToEntity(userDto)

		//errUserExist = errors.New("пользователь с таким логином уже существует")
		//errLogin     = errors.New("логин не соответствует требованиям")
		//errPassword  = errors.New("пароль не соответствует требованиям")
	)
	defer ctrl.Finish()

	tests := []TestCreateStruct{
		{
			name: "success",
			args: CreateArgs{ctx: ctx, user: userDto, password: validPassword},
			mock: func(m *mocks.MockUsersRepo) user.Repository {
				m.EXPECT().Create(ctx, gomock.Any()).Return(userDto.ID, nil)
				m.EXPECT().FindByLogin(ctx, userDto.Login).Return(nil, nil)
				return m
			},
			expected: userDto,
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mocks.NewMockUsersRepo(ctrl)
			repo := tt.mock(mock)
			svc := user.NewService(repo)

			res, err := svc.Create(tt.args.ctx, tt.args.user, tt.args.password)
			assert.Equal(t, tt.expected.ID, res.ID)
			assert.Equal(t, tt.expected.Login, res.Login)
			assert.Equal(t, tt.err, err)
		})
	}
}
