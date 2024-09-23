package database

type UserRepository interface {
	AddNewUser(user *User)
	CheckPasswordUser(user *User) bool
	GetIDByUserName(user *User) uint
	GetRefreshTokenUser(idUser string) string
	GetRolesUser(idUser string) []string
	SetRefreshToken(idUser, refreshToken string)
	GetAllInfoByIDUser(idUser string) User
	UpdateDataAccountUser(idUser string, user User)
	GetAllInfoAllAccountsAdmin(from, count int) []User
	CreateAccountByAdmin(user *User)
}

// registration a new user in db by default role with "user"
func RegisterUser(repo UserRepository, user *User) {
	repo.AddNewUser(user)
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
