package user

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"myapp/myerror"
	"myapp/payload"
	"myapp/presenter"
	"myapp/repository"
	"myapp/repository/user"
	"strings"

	"myapp/model"
)

type UserUserCase interface {
	Create(ctx context.Context, req *payload.CreateUserRequest) (*presenter.UserResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdateUserRequest) (*presenter.UserResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.UserResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListRequest) (*presenter.ListUserResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteRequest) error
}

func (u *UseCase) validateCreate(req *payload.CreateUserRequest) error {
	if req.Name == "" {
		return myerror.ErrRequestInvalidParam("name")
	}

	if req.Email == "" {
		return myerror.ErrRequestInvalidParam("Email")
	}

	if req.Username == "" {
		return myerror.ErrRequestInvalidParam("UserName")
	}

	req.Name = strings.TrimSpace(req.Name)
	if len(req.Name) == 0 {
		req.Name = ""
		return myerror.ErrRequestInvalidParam("name")
	}

	if len(req.Email) == 0 {
		return myerror.ErrRequestInvalidParam("email")
	}

	if len(req.Username) == 0 {
		return myerror.ErrRequestInvalidParam("username")
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateUserRequest,
) (*presenter.UserResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	myUser := &model.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Score:    req.Score,
	}

	err := u.UserRepo.Create(ctx, myUser)
	if err != nil {
		return nil, myerror.ErrModelCreate(err)
	}

	return &presenter.UserResponseWrapper{User: myUser}, nil
}

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdateUserRequest) (*model.User, error) {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrModelNotFound()
		}

		return nil, myerror.ErrModelGet(err)
	}

	if req.Name != nil {
		*req.Name = strings.TrimSpace(*req.Name)
		if len(*req.Name) == 0 {
			return nil, myerror.ErrRequestInvalidParam("name")
		}

		myUser.Name = *req.Name
	}

	if req.Username != nil {
		*req.Username = strings.TrimSpace(*req.Username)
		if len(*req.Username) == 0 {
			return nil, myerror.ErrRequestInvalidParam("Username")
		}

		myUser.Username = *req.Username
	}

	if req.Email != nil {
		*req.Email = strings.TrimSpace(*req.Email)
		if len(*req.Email) == 0 {
			return nil, myerror.ErrRequestInvalidParam("Email")
		}

		myUser.Email = *req.Email
	}

	if req.Score != nil {
		myUser.Score = *req.Score
	}

	return myUser, nil
}

func (u *UseCase) Update(
	ctx context.Context,
	req *payload.UpdateUserRequest,
) (*presenter.UserResponseWrapper, error) {
	myUser, err := u.validateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = u.UserRepo.Update(ctx, myUser)
	if err != nil {
		return nil, myerror.ErrModelUpdate(err)
	}

	return &presenter.UserResponseWrapper{User: myUser}, nil
}

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteRequest) error {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerror.ErrModelNotFound()
		}

		return myerror.ErrModelGet(err)
	}

	err = u.UserRepo.Delete(ctx, myUser, false)
	if err != nil {
		return myerror.ErrModelDelete(err)
	}

	return nil
}

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListRequest,
) (*presenter.ListUserResponseWrapper, error) {
	req.Format()

	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s", req.OrderBy))
	}

	myUsers, total, err := u.UserRepo.GetList(ctx, req.Search, req.Page, req.Limit, conditions, order)
	if err != nil {
		return nil, myerror.ErrModelGet(err)
	}

	if req.Page == 0 {
		req.Page = 1
	}
	return &presenter.ListUserResponseWrapper{
		Users: myUsers,
		Meta: map[string]interface{}{
			"page":  req.Page,
			"limit": req.Limit,
			"total": total,
		},
	}, nil
}

func (u *UseCase) GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.UserResponseWrapper, error) {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrModelNotFound()
		}

		return nil, myerror.ErrModelGet(err)
	}

	return &presenter.UserResponseWrapper{User: myUser}, nil
}

type UseCase struct {
	UserRepo user.Repository
}

func New(repo *repository.Repository) UserUserCase {
	return &UseCase{
		UserRepo: repo.User,
	}
}
