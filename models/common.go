package models

import (
	"time"
)

// enabled consumers per device
type Consumers struct {
	GormModel
	MQTT       bool `json:"mqtt"`
	BrewFather bool `json:"brewfather"`
}

type GormModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
