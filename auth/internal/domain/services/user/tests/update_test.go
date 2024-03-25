package tests

import (
	"context"
	"github.com/SShlykov/zeitment/auth/internal/domain/services/user"
	"github.com/SShlykov/zeitment/auth/internal/models/adapters"
	"github.com/SShlykov/zeitment/auth/internal/test/mocks/repository/mocks"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

type UpdateArgs struct {
	ctx context.Context
	req *user_v1.UpdateUserRequest
}

type TestUpdateStruct struct {
	name     string
	args     UpdateArgs
	mock     repoMockFunc
	expected *user_v1.UpdateUserResponse
	err      error
}

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)

		id         = gofakeit.UUID()
		updateUser = &user_v1.User{Id: id, Login: "login"}
		entityUser = adapters.UserProtoToEntity(updateUser)
	)
	defer ctrl.Finish()

	tests := []TestUpdateStruct{
		{
			name: "success",
			args: UpdateArgs{ctx: ctx, req: &user_v1.UpdateUserRequest{Id: id, User: updateUser}},
			mock: func(m *mocks.MockUsersRepo) user.Repository {
				m.EXPECT().Update(ctx, id, gomock.Any()).Return(entityUser, nil)
				return m
			},
			expected: &user_v1.UpdateUserResponse{
				Status: &user_v1.Status{Status: "ok", Message: ""},
				User:   updateUser,
				Id:     id,
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uss := user.NewService(tt.mock(mocks.NewMockUsersRepo(ctrl)))
			resp, err := uss.Update(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.expected.Id, resp.Id)
			assert.Equal(t, tt.expected.Status, resp.Status)
			assert.Equal(t, tt.expected.User.Login, resp.User.Login)
		})
	}
}
