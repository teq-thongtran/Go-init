package user

import (
	"github.com/labstack/echo/v4"
	"myapp/appError"
	"myapp/customError"
	"myapp/payload"
	"myapp/presenter"
	"myapp/teq"
	"myapp/usecase"
	"strconv"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}
	group.GET("/:id", r.GetByID)
	group.GET("/myself", r.GetMyself)
	group.PUT("/:id", r.Update)
	group.DELETE("/:id", r.Delete)
}

func (r *Route) Delete(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	err = r.UseCase.User.Delete(ctx, &payload.DeleteRequest{ID: id})
	if err != nil {
		return teq.Response.Error(c, err.(appError.TeqError))
	}

	return teq.Response.Success(c, nil)
}

func (r *Route) GetList(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.GetListRequest{}
		resp *presenter.ListUserResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.User.GetList(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(appError.TeqError))
	}

	return teq.Response.Success(c, resp)
}

func (r *Route) Update(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		resp  *presenter.UserResponseWrapper
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	req := payload.UpdateUserRequest{
		ID: id,
	}

	if err = c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.User.Update(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(appError.TeqError))
	}

	return teq.Response.Success(c, resp)
}

func (r *Route) GetByID(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		resp  *presenter.UserResponseWrapper
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, customError.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.User.GetByID(ctx, &payload.GetByIDRequest{ID: id})
	if err != nil {
		return teq.Response.Error(c, err.(appError.TeqError))
	}

	return teq.Response.Success(c, resp)
}

func (r *Route) GetMyself(c echo.Context) error {
	var (
		ctx = &teq.CustomEchoContext{Context: c}
	)
	userId := ctx.Context.Get("user_id").(int64)
	resp, err := r.UseCase.User.GetByID(ctx, &payload.GetByIDRequest{ID: userId})
	if err != nil {
		return teq.Response.Error(c, err.(appError.TeqError))
	}

	return teq.Response.Success(c, resp)
}
