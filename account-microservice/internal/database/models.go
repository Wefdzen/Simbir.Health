package database

type User struct {
	ID           uint
	LastName     string `json:"lastName"`
	FirstName    string `json:"firstName"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	Role         string // admin, user, doctor, manager
	RefreshToken string
}
