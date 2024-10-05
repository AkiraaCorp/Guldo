package models

import (
	"gorm.io/gorm"
)

// Odds represents the odds for an event
type OddsHistory struct {
	gorm.Model
	OddsYes      float64 `gorm:"not null"` // Odds for "yes" outcome
	OddsNo       float64 `gorm:"not null"` // Odds for "no" outcome
	EventAddress string  `gorm:"not null"` // Smart contract address for the event
}
