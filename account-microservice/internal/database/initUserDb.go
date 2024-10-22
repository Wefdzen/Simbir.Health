package database

import "wefdzen/internal/service"

// Предустановленные пользователи
func InitDbTask() {
	var defaultUsers = []User{
		{
			LastName:  "adminLast",
			FirstName: "adminFirst",
			UserName:  "admin",
			Password:  "admin", // Пароль будет хеширован
			Roles:     []string{"admin"},
		},
		{
			LastName:  "managerLast",
			FirstName: "managerFirst",
			UserName:  "manager",
			Password:  "manager",
			Roles:     []string{"manager"},
		},
		{
			LastName:  "doctorLast",
			FirstName: "doctorFirst",
			UserName:  "doctor",
			Password:  "doctor",
			Roles:     []string{"doctor"},
		},
		{
			LastName:  "userLast",
			FirstName: "userFirst",
			UserName:  "user",
			Password:  "user",
			Roles:     []string{"user"},
		},
	}
	//hashing password
	for i := range defaultUsers {
		hashedPassword, _ := service.HashPassword(defaultUsers[i].Password)
		defaultUsers[i].Password = hashedPassword
	}

	//connect to db
	userRepo := NewGormUserRepository()

	//add to db a default users
	for _, user := range defaultUsers {
		RegisterUser(userRepo, &user, user.Roles)
	}
}
