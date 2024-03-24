package server

import (
	"fmt"
	"github.com/SShlykov/zeitment/auth/internal/domain/services"
	"github.com/SShlykov/zeitment/auth/internal/infrastructure/repository/pgrepo"
	"github.com/SShlykov/zeitment/auth/internal/interceptor"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/user_v1"
	logPkg "github.com/SShlykov/zeitment/logger"
	"github.com/SShlykov/zeitment/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
)

func NewServer(logger logPkg.Logger, db postgres.Client, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)

	reflection.Register(s)

	RegisterUserService(s, db)

	logger.Info("gRPC server started", logPkg.Int("port", port))

	if err = s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func RegisterUserService(s *grpc.Server, db postgres.Client) {
	repo := pgrepo.NewUsersRepository(db)
	service := services.NewUserServiceServer(repo)
	user_v1.RegisterUserServiceServer(s, service)
}
