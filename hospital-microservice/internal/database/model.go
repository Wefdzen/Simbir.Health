package database

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Hospital struct {
	ID           uint
	Name         string         `json:"name"`
	Address      string         `json:"address"`
	ContactPhone string         `json:"contactPhone"`
	Rooms        pq.StringArray `gorm:"type:text[]" json:"rooms"`
	Deleted      gorm.DeletedAt
}
