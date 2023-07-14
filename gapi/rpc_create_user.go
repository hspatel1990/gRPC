package gapi

import (
	"context"

	db "github.com/hspatel1990/gRPC/db/sqlc"
	"github.com/hspatel1990/gRPC/pb"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	arg := db.CreateUserParams{
		Username: req.GetUsername(),
		ProvID: pgtype.Int4{
			Int32: int32(req.GetProvid()),
			Valid: req.Provid != nil,
		},
	}

	txResult, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(txResult),
	}
	return rsp, nil
}
