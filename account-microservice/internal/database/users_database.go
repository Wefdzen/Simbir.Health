package database

import (
	"fmt"
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

func (r *GormUserRepository) AddNewUser(user *User, role []string) {
	rolesToNewUser := role
	//rolesToNewUser = append(rolesToNewUser, "user")
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

func (r *GormUserRepository) UpdateDataAccountByAdmin(idUser string, user User) {
	tmp, _ := service.HashPassword(user.Password)
	user.Password = tmp
	updates := map[string]interface{}{"last_name": user.LastName, "first_name": user.FirstName, "user_name": user.UserName, "password": user.Password, "roles": user.Roles}
	r.db.Model(&User{}).Where("id = ?", idUser).Updates(updates)
}

func (r *GormUserRepository) SoftDeleteAccountByAdmin(idUser string) {
	r.db.Where("id = ?", idUser).Delete(&User{})
}

func (r *GormUserRepository) GetFullNameHowIsDoctors(from, count int, nameFilter string) []User {
	var user []User
	nameFilterC := fmt.Sprintf("%v%v%v", "%", nameFilter, "%")
	// r.db.Where("last_name LIKE ? OR first_name LIKE ?", "%"+nameFilter+"%", "%"+nameFilter+"%").Where("roles ARRAY ?", "doctor").Limit(count).Offset(from).Find(&user)
	r.db.Where("last_name LIKE ? OR first_name LIKE ?", nameFilterC, nameFilterC).
		Where("roles @> ARRAY[?]", "doctor").
		Limit(count).
		Offset(from).
		Find(&user)
	return user
}

func (r *GormUserRepository) GetInfoByIDDoctor(idUser string) User {
	var user User
	r.db.Where("id = ?", idUser).
		Where("roles @> ARRAY[?]", "doctor").
		First(&user)
	return user
}

func (r *GormUserRepository) CheckExistDoctorByID(idDoctor string) bool {
	var user User
	r.db.Where("id = ?", idDoctor).
		Where("roles @> ARRAY[?]", "doctor").
		First(&user)
	return len(user.Roles) != 0 // если равен нулю значит запись с такими критериями не было найдено
}
