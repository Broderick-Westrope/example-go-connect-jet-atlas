package transform

import (
	"github.com/Broderick-Westrope/example-go-connect-jet-atlas/gen/dev/public/model"
	userv1 "github.com/Broderick-Westrope/example-go-connect-jet-atlas/gen/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func User_InternalToV1(internal *model.Users) *userv1.User {
	return &userv1.User{
		Id:        internal.ID,
		Email:     internal.Email,
		Name:      internal.Name,
		CreatedAt: timestamppb.New(internal.CreatedAt),
		UpdatedAt: timestamppb.New(internal.UpdatedAt),
	}
}

func User_InternalFromV1(v1 *userv1.User) *model.Users {
	return &model.Users{
		ID:        v1.Id,
		Email:     v1.Email,
		Name:      v1.Name,
		CreatedAt: v1.CreatedAt.AsTime(),
		UpdatedAt: v1.UpdatedAt.AsTime(),
	}
}
