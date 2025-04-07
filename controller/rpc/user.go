package rpc

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/m03ed/gozargah-node/common"
)

func (s *Service) SyncUser(stream grpc.ClientStreamingServer[common.User, common.Empty]) error {
	for {
		user, err := stream.Recv()
		if err != nil {
			return stream.SendAndClose(&common.Empty{})
		}

		if user.GetEmail() == "" {
			return errors.New("email is required")
		}

		if err = s.GetBackend().SyncUser(stream.Context(), user); err != nil {
			return status.Errorf(codes.Internal, "failed to update user: %v", err)
		}
	}
}

func (s *Service) SyncUsers(ctx context.Context, users *common.Users) (*common.Empty, error) {
	if err := s.GetBackend().SyncUsers(ctx, users.GetUsers()); err != nil {
		return nil, err
	}

	return nil, nil
}
