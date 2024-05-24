package src

import (
	"context"
	pb "keystone_auth/rpc"
)

func (s *Server) VerifyToken(ctx context.Context, authToken *pb.AuthToken) (verifyTokenResponse *pb.VerifyTokenResponse, err error) {
	return &pb.VerifyTokenResponse{}, nil
}
