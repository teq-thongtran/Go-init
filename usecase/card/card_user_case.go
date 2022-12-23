package card

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"myapp/customError"
	"myapp/payload"
	"myapp/presenter"
	"myapp/repository"
	"myapp/repository/card"
	"myapp/repository/user"
	"strings"

	"myapp/model"
)

type CardUseCase interface {
	Create(ctx context.Context, req *payload.CreateCardRequest) (*presenter.CardResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdateCardRequest) (*presenter.CardResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.CardResponseWrapper, error)
	GetList(ctx context.Context, req *payload.GetListRequest) (*presenter.ListCardResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteRequest) error
}

type UseCase struct {
	CardRepo card.Repository
	UserRepo user.Repository
}

func New(repo *repository.Repository) CardUseCase {
	return &UseCase{
		CardRepo: repo.Card,
		UserRepo: repo.User,
	}
}

func (u *UseCase) validateCreate(req *payload.CreateCardRequest) error {
	if req.NameCard == "" {
		return customError.ErrRequestInvalidParam("name_card")
	}

	if req.CardType == "" {
		return customError.ErrRequestInvalidParam("card_type")
	}

	req.NameCard = strings.TrimSpace(req.NameCard)
	if len(req.NameCard) == 0 {
		req.NameCard = ""
		return customError.ErrRequestInvalidParam("name_card")
	}

	if len(req.CardType) == 0 {
		return customError.ErrRequestInvalidParam("card_type")
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

	myUser, err := u.UserRepo.GetByID(ctx, req.UserId)

	if err != nil {
		return nil, customError.ErrModelGet(err, "User")
	}

	myCard := &model.Card{
		NameCard: req.NameCard,
		CardType: req.CardType,
		UserId:   myUser.ID,
	}

	err = u.CardRepo.Create(ctx, myCard)
	if err != nil {
		return nil, customError.ErrModelCreate(err)
	}

	return &presenter.CardResponseWrapper{Card: myCard}, nil
}

func (u *UseCase) validateUpdate(ctx context.Context, req *payload.UpdateCardRequest) (*model.Card, error) {
	myCard, err := u.CardRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, customError.ErrModelGet(err, "Card")
	}

	if req.NameCard != nil {
		*req.NameCard = strings.TrimSpace(*req.NameCard)
		if len(*req.NameCard) == 0 {
			return nil, customError.ErrRequestInvalidParam("name")
		}

		myCard.NameCard = *req.NameCard
	}

	if req.CardType != nil {
		*req.CardType = strings.TrimSpace(*req.CardType)
		if len(*req.CardType) == 0 {
			return nil, customError.ErrRequestInvalidParam("Cardname")
		}

		myCard.CardType = *req.CardType
	}

	myUser, err := u.UserRepo.GetByID(ctx, req.UserId)

	if err != nil {
		return nil, customError.ErrModelGet(err, "User")
	}

	if myCard.UserId != myUser.ID {
		return nil, customError.ErrRequestInvalidParam("userId")
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
		return nil, customError.ErrModelUpdate(err)
	}

	return &presenter.CardResponseWrapper{Card: myCard}, nil
}

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteRequest) error {
	myCard, err := u.CardRepo.GetByID(ctx, req.ID)
	if err != nil {
		return customError.ErrModelGet(err, "Card")
	}

	err = u.CardRepo.Delete(ctx, myCard, false)
	if err != nil {
		return customError.ErrModelDelete(err)
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
	conditions["user_id"] = ctx.Value("user_id")
	myCards, total, err := u.CardRepo.GetList(ctx, req.Search, req.Page, req.Limit, conditions, order)
	if err != nil {
		return nil, customError.ErrModelGet(err, "Card")
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
			return nil, customError.ErrModelNotFound()
		}

		return nil, customError.ErrModelGet(err, "Card")
	}

	return &presenter.CardResponseWrapper{Card: myCard}, nil
}
