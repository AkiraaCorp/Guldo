package repository

import (
	"guldo/models"

	"gorm.io/gorm"
)

type OddsRepository struct {
	db *gorm.DB
}

func NewOddsRepository(db *gorm.DB) *OddsRepository {
	return &OddsRepository{
		db: db,
	}
}

func (o *OddsRepository) Create(odds models.OddsHistory) error {
	return o.db.Create(&odds).Error
}
