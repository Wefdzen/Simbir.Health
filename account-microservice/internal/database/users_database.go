package database

import (
	"log"
	"wefdzen/internal/service"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository() *GormUserRepository {
	db, err := Connect()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) AddNewUser(user *User) {
	var rolesToNewUser []string
	rolesToNewUser = append(rolesToNewUser, "user")
	r.db.Create(&User{LastName: user.LastName, FirstName: user.FirstName, UserName: user.UserName, Password: user.Password, Roles: rolesToNewUser, RefreshToken: ""})
}

func (r *GormUserRepository) CheckPasswordUser(user *User) bool {
	var tmp User
	r.db.Where("user_name = ?", user.UserName).First(&tmp)
	return service.CheckPassword(user.Password, tmp.Password)
}

func (r *GormUserRepository) GetIDByUserName(user *User) uint {
	var tmp User
	r.db.Where("user_name = ?", user.UserName).First(&tmp)
	return tmp.ID
}

func (r *GormUserRepository) GetRefreshTokenUser(idUser string) string {
	var user User
	r.db.Where("id = ?", idUser).First(&user)
	return user.RefreshToken
}

func (r *GormUserRepository) GetRolesUser(idUser string) []string {
	var user User
	r.db.Where("id = ?", idUser).First(&user)
	return user.Roles
}

func (r *GormUserRepository) SetRefreshToken(idUser, refreshToken string) {
	r.db.Model(&User{}).Where("id = ?", idUser).UpdateColumn("refresh_token", refreshToken)
}

func (r *GormUserRepository) GetAllInfoByIDUser(idUser string) User {
	var user User
	r.db.Where("id = ?", idUser).First(&user)
	return user
}

func (r *GormUserRepository) UpdateDataAccountUser(idUser string, user User) {
	tmp, _ := service.HashPassword(user.Password)
	user.Password = tmp
	updates := map[string]interface{}{"last_name": user.LastName, "first_name": user.FirstName, "password": user.Password}
	r.db.Model(&User{}).Where("id = ?", idUser).Updates(updates)
}

func (r *GormUserRepository) GetAllInfoAllAccountsAdmin(from, count int) []User {
	var user []User
	r.db.Limit(count).Offset(from).Find(&user)
	return user
}

func (r *GormUserRepository) CreateAccountByAdmin(user *User) {
	r.db.Create(&User{LastName: user.LastName, FirstName: user.FirstName, UserName: user.UserName, Password: user.Password, Roles: user.Roles, RefreshToken: ""})
}

// func (r *GormUserRepository) UpdateDataAccountByAdmin(idUser string, user *User) {
// 	r.db.Create(&User{LastName: user.LastName, FirstName: user.FirstName, UserName: user.UserName, Password: user.Password, Roles: user.Roles, RefreshToken: ""})
// }

func (r *GormUserRepository) UpdateDataAccountByAdmin(idUser string, user User) {
	tmp, _ := service.HashPassword(user.Password)
	user.Password = tmp
	updates := map[string]interface{}{"last_name": user.LastName, "first_name": user.FirstName, "user_name": user.UserName, "password": user.Password, "roles": user.Roles}
	r.db.Model(&User{}).Where("id = ?", idUser).Updates(updates)
}

func (r *GormUserRepository) SoftDeleteAccountByAdmin(idUser string) {
	r.db.Where("id = ?", idUser).Delete(&User{})
}
