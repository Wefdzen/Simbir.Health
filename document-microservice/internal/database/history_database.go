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

func (r *GormUserRepository) CreateHistory(history *History) {
	r.db.Create(&History{Date: history.Date, PacientId: history.PacientId, HospitalId: history.PacientId, DoctorId: history.DoctorId, Room: history.Room, Data: history.Data})
}

func (r *GormUserRepository) UpdateHistory(idHistory string, history *History) {
	updates := map[string]interface{}{"date": history.Date, "pacient_id": history.PacientId, "hospital_id": history.HospitalId, "doctor_id": history.DoctorId, "room": history.Room, "data": history.Data}
	r.db.Model(&History{}).Where("id = ?", idHistory).Updates(updates)
}

func (r *GormUserRepository) GetListHistory(pacientId string) []History {
	var history []History
	r.db.Where("pacient_id = ?", pacientId).Find(&history)
	return history
}

func (r *GormUserRepository) GetListHistoryById(historyId string) History {
	var history History
	r.db.Where("id = ?", historyId).First(&history)
	return history
}
