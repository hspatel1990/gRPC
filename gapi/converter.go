package gapi

import (
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Id:        user.ID,
		Username:  user.Username,
		ProvId:    int64(user.ProvID.Int32),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
