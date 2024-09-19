package database

import (
	"log"
	"wefdzen/internal/service"

	"gorm.io/gorm"
)

type GormUserRepositroy struct {
	db *gorm.DB
}

func NewGormUserRepositroy() *GormUserRepositroy {
	db, err := Connect()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return &GormUserRepositroy{db: db}
}

func (r *GormUserRepositroy) AddNewUser(user *User) {
	r.db.Create(&User{LastName: user.LastName, FirstName: user.FirstName, UserName: user.UserName, Password: user.Password, Role: "user", RefreshToken: ""})
}

func (r *GormUserRepositroy) CheckPasswordUser(user *User) bool {
	var tmp User
	r.db.Where("user_name = ?", user.UserName).First(&tmp)
	return service.CheckPassword(user.Password, tmp.Password)
}

func (r *GormUserRepositroy) GetIDByUserName(user *User) uint {
	var tmp User
	r.db.Where("user_name = ?", user.UserName).First(&tmp)
	return tmp.ID
}

func (r *GormUserRepositroy) GetRefreshTokenUser(idUser string) string {
	var user User
	r.db.Where("id = ?", idUser).First(&user)
	return user.RefreshToken
}

func (r *GormUserRepositroy) GetRoleUser(idUser string) string {
	var user User
	r.db.Where("id = ?", idUser).First(&user)
	return user.Role
}

func (r *GormUserRepositroy) SetRefreshToken(idUser, refreshToken string) {
	r.db.Model(&User{}).Where("id = ?", idUser).UpdateColumn("refresh_token", refreshToken)
}
