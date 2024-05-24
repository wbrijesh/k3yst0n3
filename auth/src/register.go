package src

import (
	"context"
	pb "keystone_auth/rpc"
)

func (s *Server) Register(ctx context.Context, user *pb.User) (authToken *pb.AuthToken, err error) {
	return &pb.AuthToken{}, nil
}
