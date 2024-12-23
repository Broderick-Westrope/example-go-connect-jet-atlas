package validation

import (
	userv1 "github.com/Broderick-Westrope/example-go-connect-jet-atlas/gen/user/v1"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateCreateUser(req *userv1.CreateUserRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Email,
			validation.Required,
			validation.Length(5, 100),
			is.Email,
		),
		validation.Field(&req.Name,
			validation.Required,
			validation.Length(2, 50),
		),
	)
}

func ValidateGetUser(req *userv1.GetUserRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Id,
			validation.Required,
		),
	)
}

func ValidateUpdateUser(req *userv1.UpdateUserRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Id,
			validation.Required,
		),
		validation.Field(&req.Name,
			validation.Required,
			validation.Length(2, 50),
		),
		validation.Field(&req.Email,
			validation.Required,
			validation.Length(5, 100),
			is.Email,
		),
	)
}

func ValidateDeleteUser(req *userv1.DeleteUserRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Id,
			validation.Required,
		),
	)
}
