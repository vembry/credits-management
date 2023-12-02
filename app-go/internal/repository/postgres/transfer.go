package postgres

import (
	"app/internal/model"
	"context"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type transfer struct {
	db *gorm.DB
}

func NewTransfer(db *gorm.DB) *transfer {
	return &transfer{
		db: db,
	}
}

func (r *transfer) Create(ctx context.Context, entry *model.Transfer) (*model.Transfer, error) {
	entry.Id = ksuid.New()

	if err := r.db.Table("transfers").Create(entry).WithContext(ctx).Error; err != nil {
		return nil, err
	}
	return entry, nil
}