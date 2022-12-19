package card

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"myapp/myerror"
	"myapp/payload"
	"myapp/presenter"
	"myapp/repository"
	"myapp/repository/card"
	"strings"

	"myapp/model"
)

type CardCardCase interface {
	Create(ctx context.Context, req *payload.CreateCardRequest) (*presenter.CardResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdateCardRequest) (*presenter.CardResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.CardResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListRequest) (*presenter.ListCardResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteRequest) error
}

type UseCase struct {
	CardRepo card.Repository
}

func New(repo *repository.Repository) CardCardCase {
	return &UseCase{
		CardRepo: repo.Card,
	}
}

func (u *UseCase) validateCreate(req *payload.CreateCardRequest) error {
	if req.NameCard == "" {
		return myerror.ErrRequestInvalidParam("name_card")
	}

	if req.CardType == "" {
		return myerror.ErrRequestInvalidParam("card_type")
	}

	req.NameCard = strings.TrimSpace(req.NameCard)
	if len(req.NameCard) == 0 {
		req.NameCard = ""
		return myerror.ErrRequestInvalidParam("name_card")
	}

	if len(req.CardType) == 0 {
		return myerror.ErrRequestInvalidParam("card_type")
	}

	if req.UserId == 0 {
		return myerror.ErrRequestInvalidParam("user_id")
	}

	return nil
}

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateCardRequest,
) (*presenter.CardResponseWrapper, error) {
	if err := u.validateCreate(req); err != nil {
		return nil, err
	}

	myCard := &model.Card{
		NameCard: req.NameCard,
		CardType: req.CardType,
		UserId:   req.UserId,
	}

	err := u.CardRepo.Create(ctx, myCard)
	if err != nil {
		return nil, myerror.ErrModelCreate(err)
	}

	return &presenter.CardResponseWrapper{Card: myCard}, nil
}

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdateCardRequest) (*model.Card, error) {
	myCard, err := u.CardRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrModelNotFound()
		}

		return nil, myerror.ErrModelGet(err)
	}

	if req.NameCard != nil {
		*req.NameCard = strings.TrimSpace(*req.NameCard)
		if len(*req.NameCard) == 0 {
			return nil, myerror.ErrRequestInvalidParam("name")
		}

		myCard.NameCard = *req.NameCard
	}

	if req.CardType != nil {
		*req.CardType = strings.TrimSpace(*req.CardType)
		if len(*req.CardType) == 0 {
			return nil, myerror.ErrRequestInvalidParam("Cardname")
		}

		myCard.CardType = *req.CardType
	}

	if req.UserId != 0 {
		myCard.UserId = req.UserId
	}

	return myCard, nil
}

func (u *UseCase) Update(
	ctx context.Context,
	req *payload.UpdateCardRequest,
) (*presenter.CardResponseWrapper, error) {
	myCard, err := u.validateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = u.CardRepo.Update(ctx, myCard)
	if err != nil {
		return nil, myerror.ErrModelUpdate(err)
	}

	return &presenter.CardResponseWrapper{Card: myCard}, nil
}

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteRequest) error {
	myCard, err := u.CardRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerror.ErrModelNotFound()
		}

		return myerror.ErrModelGet(err)
	}

	err = u.CardRepo.Delete(ctx, myCard, false)
	if err != nil {
		return myerror.ErrModelDelete(err)
	}

	return nil
}

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListRequest,
) (*presenter.ListCardResponseWrapper, error) {
	req.Format()

	var (
		order      = make([]string, 0)
		conditions = map[string]interface{}{}
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s", req.OrderBy))
	}

	myCards, total, err := u.CardRepo.GetList(ctx, req.Search, req.Page, req.Limit, conditions, order)
	if err != nil {
		return nil, myerror.ErrModelGet(err)
	}

	if req.Page == 0 {
		req.Page = 1
	}
	return &presenter.ListCardResponseWrapper{
		Cards: myCards,
		Meta: map[string]interface{}{
			"page":  req.Page,
			"limit": req.Limit,
			"total": total,
		},
	}, nil
}

func (u *UseCase) GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.CardResponseWrapper, error) {
	myCard, err := u.CardRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrModelNotFound()
		}

		return nil, myerror.ErrModelGet(err)
	}

	return &presenter.CardResponseWrapper{Card: myCard}, nil
}
