package repository

import (
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (o *EventRepository) GetAllActiveEvents() ([]string, error) {
	var addresses []string
	sqlQuery := "SELECT address FROM events WHERE is_active = ?"
	result := o.db.Raw(sqlQuery, true).Scan(&addresses)
	if result.Error != nil {
		return nil, result.Error
	}
	return addresses, nil
}
