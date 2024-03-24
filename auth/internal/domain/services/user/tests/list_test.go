package tests

import (
	"github.com/SShlykov/zeitment/auth/internal/domain/services/user"
	"github.com/SShlykov/zeitment/auth/internal/models/adapters"
	"github.com/SShlykov/zeitment/auth/internal/test/mocks/repository/mocks"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	"github.com/SShlykov/zeitment/postgres/dbutils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/net/context"
	"testing"
)

type Args struct {
	ctx context.Context
	req *user_v1.ListUsersRequest
}

type repoMockFunc func(*mocks.MockUsersRepo) user.Repository

type TestStruct struct {
	name     string
	args     Args
	mock     repoMockFunc
	expected *user_v1.ListUsersResponse
	err      error
}

func TestFindUsers(t *testing.T) {
	t.Parallel()

	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)

		page     uint64 = 1
		pageSize uint64 = 10

		req = user_v1.ListUsersRequest{
			Options: &user_v1.GetOptions{Pagination: &user_v1.Pagination{Page: page, PageSize: pageSize}},
		}

		resGrpc = []*user_v1.User{{Id: "1", Login: "login1"}, {Id: "2", Login: "login2"}, {Id: "3", Login: "login3"}}
		resEnt  = adapters.ProtoToUsers(resGrpc)
		status  = &user_v1.Status{Status: "ok", Message: ""}

		metaTemplate = &user_v1.PaginationMetadata{Page: 1, PageSize: 10, Total: 3, TotalPages: 1}
	)
	defer ctrl.Finish()

	tests := []TestStruct{
		{
			name: "success",
			args: Args{ctx: ctx, req: &req},
			mock: func(m *mocks.MockUsersRepo) user.Repository {
				m.EXPECT().List(ctx, dbutils.NewPaginationWithLimitOffset(metaTemplate.Page, metaTemplate.PageSize)).Return(resEnt, nil)
				m.EXPECT().Count(ctx).Return(uint64(len(resEnt)), nil)
				return m
			},
			expected: &user_v1.ListUsersResponse{Users: resGrpc, Status: status, PaginationMetadata: metaTemplate},
			err:      nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := mocks.NewMockUsersRepo(ctrl)
			svc := user.NewUserServiceServer(mockRepo)
			tt.mock(mockRepo)
			got, err := svc.Find(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.expected.Status.Status, got.Status.Status)
			assert.Equal(t, tt.expected.Status.Message, got.Status.Message)
			assert.Equal(t, tt.expected.PaginationMetadata, got.PaginationMetadata)
		})
	}
}
