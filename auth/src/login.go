package src

import (
	"context"
	pb "keystone_auth/rpc"
)

type Server struct{}

func (s *Server) Login(ctx context.Context, user *pb.Credentials) (authToken *pb.AuthToken, err error) {
	return &pb.AuthToken{}, nil
}
