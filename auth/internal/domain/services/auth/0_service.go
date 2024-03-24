package auth

import (
	"github.com/SShlykov/zeitment/auth/internal/domain/services/user"
	"github.com/SShlykov/zeitment/auth/pkg/grpc/auth_v1"
)

type Service struct {
	auth_v1.UnimplementedAuthServiceServer
	userService *user.Service
}

func NewAuthService(userService *user.Service) auth_v1.AuthServiceServer {
	return &Service{userService: userService}
}
