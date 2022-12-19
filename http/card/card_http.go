package card

import (
	"github.com/labstack/echo/v4"
	"myapp/apperror"
	"myapp/myerror"
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

	group.POST("", r.Create)
	group.GET("", r.GetList)
	group.GET("/:id", r.GetByID)
	group.PUT("/:id", r.Update)
	group.DELETE("/:id", r.Delete)
}

func (r *Route) Create(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.CreateCardRequest{}
		resp *presenter.CardResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, myerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Card.Create(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(apperror.TeqError))
	}

	return teq.Response.Success(c, resp)
}

func (r *Route) Delete(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, myerror.ErrInvalidParams(err))
	}

	err = r.UseCase.Card.Delete(ctx, &payload.DeleteRequest{ID: id})
	if err != nil {
		return teq.Response.Error(c, err.(apperror.TeqError))
	}

	return teq.Response.Success(c, nil)
}

func (r *Route) GetList(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.GetListRequest{}
		resp *presenter.ListCardResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, myerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Card.GetList(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(apperror.TeqError))
	}

	return teq.Response.Success(c, resp)
}

func (r *Route) Update(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		resp  *presenter.CardResponseWrapper
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, myerror.ErrInvalidParams(err))
	}

	req := payload.UpdateCardRequest{
		ID: id,
	}

	if err = c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, myerror.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.Card.Update(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(apperror.TeqError))
	}

	return teq.Response.Success(c, resp)
}

func (r *Route) GetByID(c echo.Context) error {
	var (
		ctx   = &teq.CustomEchoContext{Context: c}
		idStr = c.Param("id")
		resp  *presenter.CardResponseWrapper
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return teq.Response.Error(ctx, myerror.ErrInvalidParams(err))
	}

	resp, err = r.UseCase.Card.GetByID(ctx, &payload.GetByIDRequest{ID: id})
	if err != nil {
		return teq.Response.Error(c, err.(apperror.TeqError))
	}

	return teq.Response.Success(c, resp)
}