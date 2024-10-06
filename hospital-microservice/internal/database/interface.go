package database

import "fmt"

type UserRepository interface {
	CreateHospitalAdmin(hospital *Hospital)
	SoftDeleteHospitalByAdmin(idHospital string)
	UpdateDataHospitalByAdmin(idHospital string, hospital *Hospital)
	GetListHospitalsByUser(from, count int) []Hospital
	GetInfoAboutHospitalByID(idHospital string) Hospital
	CheckExistRoomHospitalID(room, idHospital string) bool
}

func CreateHospitalByAdmin(repo UserRepository, hospital *Hospital) {
	repo.CreateHospitalAdmin(hospital)
}

func SoftDeleteByAdmin(repo UserRepository, idHospital string) {
	repo.SoftDeleteHospitalByAdmin(idHospital)
}

func UpdateHospitalData(repo UserRepository, idHospital string, hospital *Hospital) {
	repo.UpdateDataHospitalByAdmin(idHospital, hospital)
}

func GetListHospitals(repo UserRepository, from, count int) []Hospital {
	return repo.GetListHospitalsByUser(from, count)
}

func GetInfoHospitalByID(repo UserRepository, idHospital string) Hospital {
	return repo.GetInfoAboutHospitalByID(idHospital)
}

func CheckExistRoomInHospitalID(repo UserRepository, room, idHospital string) bool {
	fmt.Println("ROOOOOOM: ", room)
	return repo.CheckExistRoomHospitalID(room, idHospital)
}
