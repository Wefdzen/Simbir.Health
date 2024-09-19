package database

type UserRepository interface {
	AddNewUser(user *User)
	CheckPasswordUser(user *User) bool
	GetIDByUserName(user *User) uint
	GetRefreshTokenUser(idUser string) string
	GetRoleUser(idUser string) string
	SetRefreshToken(idUser, refreshToken string)
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

func GetRole(repo UserRepository, idUser string) string {
	return repo.GetRoleUser(idUser)
}

func SetRefToken(repo UserRepository, idUser, refreshToken string) {
	repo.SetRefreshToken(idUser, refreshToken)
}
