package main

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
	userv1 "github.com/Broderick-Westrope/example-go-connect-jet-atlas/gen/user/v1"
	"github.com/Broderick-Westrope/example-go-connect-jet-atlas/internal/data"
	"github.com/Broderick-Westrope/example-go-connect-jet-atlas/internal/transform"
	"github.com/Broderick-Westrope/example-go-connect-jet-atlas/internal/validation"
)

func (app *application) CreateUser(
	ctx context.Context,
	req *connect.Request[userv1.CreateUserRequest],
) (*connect.Response[userv1.CreateUserResponse], error) {
	err := validation.ValidateCreateUser(req.Msg)
	if err != nil {
		app.log.ErrorContext(ctx, "invalid request", slog.Any("err", err))
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
	}

	user, err := app.repo.CreateUser(ctx, req.Msg.Email, req.Msg.Name)
	if err != nil {
		app.log.ErrorContext(ctx, "failed to create user", slog.Any("err", err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	return connect.NewResponse(&userv1.CreateUserResponse{
		User: transform.User_InternalToV1(user),
	}), nil
}

func (app *application) GetUser(
	ctx context.Context,
	req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.GetUserResponse], error) {
	err := validation.ValidateGetUser(req.Msg)
	if err != nil {
		app.log.ErrorContext(ctx, "invalid request", slog.Any("err", err))
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
	}

	user, err := app.repo.GetUserByID(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, nil)
		}
		app.log.ErrorContext(ctx, "failed to get user", slog.Any("err", err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	return connect.NewResponse(&userv1.GetUserResponse{
		User: transform.User_InternalToV1(user),
	}), nil
}

func (app *application) UpdateUser(
	ctx context.Context,
	req *connect.Request[userv1.UpdateUserRequest],
) (*connect.Response[userv1.UpdateUserResponse], error) {
	err := validation.ValidateUpdateUser(req.Msg)
	if err != nil {
		app.log.ErrorContext(ctx, "invalid request", slog.Any("err", err))
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
	}

	user, err := app.repo.UpdateUser(ctx, req.Msg.Id, req.Msg.Email, req.Msg.Name)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, nil)
		}
		app.log.ErrorContext(ctx, "failed to update user", slog.Any("err", err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	return connect.NewResponse(&userv1.UpdateUserResponse{
		User: transform.User_InternalToV1(user),
	}), nil
}

func (app *application) DeleteUser(
	ctx context.Context,
	req *connect.Request[userv1.DeleteUserRequest],
) (*connect.Response[userv1.DeleteUserResponse], error) {
	err := validation.ValidateDeleteUser(req.Msg)
	if err != nil {
		app.log.ErrorContext(ctx, "invalid request", slog.Any("err", err))
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
	}

	err = app.repo.DeleteUser(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, nil)
		}
		app.log.ErrorContext(ctx, "failed to delete user", slog.Any("err", err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	return connect.NewResponse(&userv1.DeleteUserResponse{}), nil
}
