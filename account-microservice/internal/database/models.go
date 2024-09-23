package database

import "github.com/lib/pq"

type User struct {
	ID           uint
	LastName     string         `json:"lastName"`
	FirstName    string         `json:"firstName"`
	UserName     string         `json:"username"`
	Password     string         `json:"password"`
	Roles        pq.StringArray `gorm:"type:text[]" json:"roles"` // admin, user, doctor, manager
	RefreshToken string
}
