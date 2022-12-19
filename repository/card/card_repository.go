package card

import (
	"context"

	"gorm.io/gorm"
	"myapp/model"
)

type Repository interface {
	Create(ctx context.Context, data *model.Card) error
	Update(ctx context.Context, data *model.Card) error
	GetByID(ctx context.Context, id int64) (*model.Card, error)
	Delete(ctx context.Context, data *model.Card, unscoped bool) error
	GetList(
		ctx context.Context,
		search string,
		page int,
		limit int,
		conditions interface{},
		order []string,
	) ([]model.Card, int64, error)
}

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p *pgRepository) Create(ctx context.Context, data *model.Card) error {
	return p.getDB(ctx).Create(data).Error
}

func (p *pgRepository) Update(ctx context.Context, data *model.Card) error {
	return p.getDB(ctx).Save(data).Error
}

func (p *pgRepository) GetByID(ctx context.Context, id int64) (*model.Card, error) {
	var user model.Card

	err := p.getDB(ctx).
		Where("id = ?", id).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *pgRepository) Delete(ctx context.Context, data *model.Card, unscoped bool) error {
	db := p.getDB(ctx)

	if unscoped {
		db = db.Unscoped()
	}

	return db.Delete(data).Error
}

func (p *pgRepository) GetList(
	ctx context.Context,
	search string,
	page int,
	limit int,
	conditions interface{},
	order []string,
) ([]model.Card, int64, error) {
	var (
		db     = p.getDB(ctx).Model(&model.Card{}).Preload("User")
		data   = make([]model.Card, 0)
		total  int64
		offset int
	)

	if conditions != nil {
		db = db.Where(conditions)
	}

	for i := range order {
		db = db.Order(order[i])
	}

	if page != 1 {
		offset = limit * (page - 1)
	}

	if limit != -1 {
		err := db.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	}

	err := db.Limit(limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	if limit == -1 {
		total = int64(len(data))
	}

	return data, total, nil
}
