package database

import (
	"log"

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

func (r *GormUserRepository) CreateHospitalAdmin(hospital *Hospital) {
	r.db.Create(&Hospital{Name: hospital.Name, Address: hospital.Address, ContactPhone: hospital.ContactPhone, Rooms: hospital.Rooms})
}

func (r *GormUserRepository) SoftDeleteHospitalByAdmin(idHospital string) {
	r.db.Where("id = ?", idHospital).Delete(&Hospital{})
}

func (r *GormUserRepository) UpdateDataHospitalByAdmin(idHospital string, hospital *Hospital) {
	updates := map[string]interface{}{"name": hospital.Name, "address": hospital.Address, "contact_phone": hospital.ContactPhone, "rooms": hospital.Rooms}
	r.db.Model(&Hospital{}).Where("id = ?", idHospital).Updates(updates)
}

func (r *GormUserRepository) GetListHospitalsByUser(from, count int) []Hospital {
	var hospitals []Hospital
	r.db.Limit(count).Offset(from).Find(&hospitals)
	return hospitals
}

func (r *GormUserRepository) GetInfoAboutHospitalByID(idHospital string) Hospital {
	var hospital Hospital
	r.db.Where("id = ?", idHospital).First(&hospital)
	return hospital
}
