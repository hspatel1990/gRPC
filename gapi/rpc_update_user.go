package gapi

import (
	"context"
	"errors"

	db "github.com/hspatel1990/gRPC/db/sqlc"
	"github.com/hspatel1990/gRPC/pb"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {

	arg := db.UpdateUserParams{

		ID: req.Id,
		Username: pgtype.Text{
			String: req.GetUsername(),
			Valid:  req.Username != nil,
		},

		ProvID: pgtype.Int4{
			Int32: int32(req.GetProvid()),
			Valid: req.Provid != nil,
		},
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	rsp := &pb.UpdateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}
