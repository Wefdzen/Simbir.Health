package database

import (
	"log"
	"time"

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

func (r *GormUserRepository) CreateNewTimetableForDoctor(newTimetable *Timetable) {
	r.db.Create(&Timetable{HospitalId: newTimetable.HospitalId, DoctorId: newTimetable.DoctorId, From: newTimetable.From, To: newTimetable.To, Room: newTimetable.Room})
}

// TODO Нельзя изменить если есть записавшиеся на прием
func (r *GormUserRepository) UpdateDataInTimetable(idTimetable string, newTimetable Timetable) {
	updates := map[string]interface{}{"hospital_id": newTimetable.HospitalId, "doctor_id": newTimetable.DoctorId, "from": newTimetable.From, "to": newTimetable.To, "room": newTimetable.Room}
	r.db.Model(&Timetable{}).Where("id = ?", idTimetable).Updates(updates)
}

func (r *GormUserRepository) DeleteDataTimetable(idTimetable string) {
	r.db.Where("id = ?", idTimetable).Delete(&Timetable{})
}

func (r *GormUserRepository) DeleteDataTimetableForDoctor(idDoctor string) {
	r.db.Where("doctor_id = ?", idDoctor).Delete(&Timetable{})
}

func (r *GormUserRepository) DeleteDataTimetableForHospital(idHospital string) {
	r.db.Where("hospital_id = ?", idHospital).Delete(&Timetable{})
}

func (r *GormUserRepository) GetTimetableByIdHospital(idHospital string, from, to time.Time) []Timetable {
	var timeTables []Timetable
	fromUTC := from.UTC()
	toUTC := to.UTC()
	_ = toUTC
	r.db.Where("hospital_id = ?", idHospital).
		// Where("'from' >= ? AND 'to' <= ?", from, to).
		Where("from >= ?", fromUTC).
		Find(&timeTables)
	return timeTables
}
