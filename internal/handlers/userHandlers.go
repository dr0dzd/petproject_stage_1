package handlers

import (
	"Golang/internal/userService"
	"Golang/internal/web/users"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	service *userService.UserService
}

func NewUserHandler(serv *userService.UserService) *UserHandler {
	return &UserHandler{service: serv}
}

func (uh *UserHandler) GetUsersUserId(_ context.Context, request users.GetUsersUserIdRequestObject) (users.GetUsersUserIdResponseObject, error) {
	user, err := uh.service.GetUserByID(request.UserId)
	if err != nil {
		return users.GetUsersUserId404Response{}, err
	}
	response := users.GetUsersUserId200JSONResponse{
		Id:       &user.ID,
		Email:    &user.Email,
		Password: &user.Password,
	}

	return response, nil
}

func (uh *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	AllUsers, err := uh.service.GetUsers()

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Problem with Data Base")
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
	if err := uh.service.DeleteUserByID(uint(id)); err != nil {
		return nil, err
	}
	response := users.DeleteUsersUserId204Response{}
	return response, nil
}

func (uh *UserHandler) PatchUsersUserId(_ context.Context, request users.PatchUsersUserIdRequestObject) (users.PatchUsersUserIdResponseObject, error) {
	id := request.UserId
	var ForUpdate userService.User

	if err := updateUserFields(&request, &ForUpdate); err != nil {
		return nil, err
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

func updateUserFields(request *users.PatchUsersUserIdRequestObject, user *userService.User) error {
	req := request.Body
	if req.Id != nil {
		user.ID = *req.Id
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Password != nil {
		user.Password = *req.Password
	}
	return nil
}
