package handlers

import (
	"Golang/internal/userService"
	"Golang/internal/web/users"
	"context"
)

type UserHandler struct {
	service *userService.UserService
}

func NewUserHandler(serv *userService.UserService) *UserHandler {
	return &UserHandler{service: serv}
}

func (uh *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	AllUsers, err := uh.service.GetUsers()

	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range AllUsers {
		user := users.User{
			Password: &usr.Password,
			Email:    &usr.Email,
			Id:       &usr.ID,
		}
		response = append(response, user)
	}

	return response, nil
}

func (uh *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	req := request.Body
	LocalUser := userService.User{
		Email:    *req.Email,
		Password: *req.Password,
	}

	Local, err := uh.service.CreateUser(LocalUser)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &Local.ID,
		Email:    &Local.Email,
		Password: &Local.Password,
	}
	return response, nil
}

func (uh *UserHandler) DeleteUsersUserId(_ context.Context, request users.DeleteUsersUserIdRequestObject) (users.DeleteUsersUserIdResponseObject, error) {
	id := request.UserId
	if err := uh.service.DeleteTaskByID(uint(id)); err != nil {
		return nil, err
	}
	response := users.DeleteUsersUserId204Response{}
	return response, nil
}

func (uh *UserHandler) PatchUsersUserId(_ context.Context, request users.PatchUsersUserIdRequestObject) (users.PatchUsersUserIdResponseObject, error) {
	req := request.Body
	id := request.UserId
	var ForUpdate userService.User
	if req.Id != nil {
		ForUpdate.ID = *req.Id
	}
	if req.Email != nil {
		ForUpdate.Email = *req.Email
	}
	if req.Password != nil {
		ForUpdate.Password = *req.Password
	}
	UpdatedUser, err := uh.service.UpdateUserByID(id, ForUpdate)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsersUserId200JSONResponse{
		Id:       &UpdatedUser.ID,
		Email:    &UpdatedUser.Email,
		Password: &UpdatedUser.Password,
	}
	return response, nil
}
