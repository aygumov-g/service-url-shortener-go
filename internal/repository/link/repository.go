package link

import (
	"errors"
	"time"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(link *link_d.Link) error {
	return r.db.Create(link).Error
}

func (r *repository) Update(id int64, now time.Time) error {
	return r.db.Model(&link_d.Link{}).
		Where(&link_d.Link{ID: id}).
		UpdateColumns(map[string]interface{}{
			"click_count":      gorm.Expr("click_count + 1"),
			"last_accessed_at": now,
		}).Error
}

func (r *repository) GetByID(id int64) (*link_d.Link, error) {
	var link link_d.Link

	err := r.db.Where(&link_d.Link{ID: id}).First(&link).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, link_d.ErrLinkNotFound
		default:
			return nil, err
		}
	}

	return &link, err
}

func (r *repository) GetByCustomCode(code string) (*link_d.Link, error) {
	var link link_d.Link

	err := r.db.Where("custom_code = ?", code).First(&link).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, link_d.ErrLinkNotFound
		default:
			return nil, err
		}
	}

	return &link, err
}
