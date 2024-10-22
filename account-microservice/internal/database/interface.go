package database

type UserRepository interface {
	AddNewUser(user *User, role []string)
	CheckPasswordUser(user *User) bool
	GetIDByUserName(user *User) uint
	GetRefreshTokenUser(idUser string) string
	GetRolesUser(idUser string) []string
	SetRefreshToken(idUser, refreshToken string)
	GetAllInfoByIDUser(idUser string) User
	UpdateDataAccountUser(idUser string, user User)
	GetAllInfoAllAccountsAdmin(from, count int) []User
	CreateAccountByAdmin(user *User)
	UpdateDataAccountByAdmin(idUser string, user User)
	SoftDeleteAccountByAdmin(idUser string)
	GetFullNameHowIsDoctors(from, count int, nameFilter string) []User
	GetInfoByIDDoctor(idUser string) User
	CheckExistDoctorByID(idDoctor string) bool
}

// registration a new user in db by default role with "user"
func RegisterUser(repo UserRepository, user *User, role []string) {
	repo.AddNewUser(user, role)
}

func CheckPassword(repo UserRepository, user *User) bool {
	return repo.CheckPasswordUser(user) //true if password equal with passw in db
}

func GetID(repo UserRepository, user *User) uint {
	return repo.GetIDByUserName(user)
}

func GetRefToken(repo UserRepository, idUser string) string {
	return repo.GetRefreshTokenUser(idUser)
}

func GetRoles(repo UserRepository, idUser string) []string {
	return repo.GetRolesUser(idUser)
}

func SetRefToken(repo UserRepository, idUser, refreshToken string) {
	repo.SetRefreshToken(idUser, refreshToken)
}

func GetAllInfoByID(repo UserRepository, idUser string) User {
	return repo.GetAllInfoByIDUser(idUser)
}

func UpdateDataAccount(repo UserRepository, idUser string, user User) {
	repo.UpdateDataAccountUser(idUser, user)
}

func GetAllInfoAllAccounts(repo UserRepository, from, count int) []User {
	return repo.GetAllInfoAllAccountsAdmin(from, count)
}

func NewAccountByAdmin(repo UserRepository, user *User) {
	repo.CreateAccountByAdmin(user)
}
func UpdateDataAccountAdmin(repo UserRepository, idUser string, user User) {
	repo.UpdateDataAccountByAdmin(idUser, user)
}
func SoftDeleteAccountAdmin(repo UserRepository, idUser string) {
	repo.SoftDeleteAccountByAdmin(idUser)
}

func GetFullNameHowDoctors(repo UserRepository, from, count int, nameFilter string) []User {
	return repo.GetFullNameHowIsDoctors(from, count, nameFilter)
}
func GetInfoIDDoctor(repo UserRepository, idUser string) User {
	return repo.GetInfoByIDDoctor(idUser)
}

func CheckExistDoctorID(repo UserRepository, idDoctor string) bool {
	return repo.CheckExistDoctorByID(idDoctor)
}
