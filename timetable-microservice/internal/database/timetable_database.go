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

	r.db.Where("hospital_id = ?", idHospital).
		Where("\"from\" >= ? AND \"to\" <= ?", from, to).
		Find(&timeTables)

	return timeTables
}

func (r *GormUserRepository) GetTimetableByIdDoctor(idDoctor string, from, to time.Time) []Timetable {
	var timeTables []Timetable

	r.db.Where("doctor_id = ?", idDoctor).
		Where("\"from\" >= ? AND \"to\" <= ?", from, to).
		Find(&timeTables)

	return timeTables
}

func (r *GormUserRepository) GetTimetableByIdHospitalAndRoom(idHospital, room string, from, to time.Time) []Timetable {
	var timeTables []Timetable

	r.db.Where("hospital_id = ?", idHospital).
		Where("room = ? ", room).
		Where("\"from\" >= ? AND \"to\" <= ?", from, to).
		Find(&timeTables)

	return timeTables
}

func (r *GormUserRepository) CreateNewAppointmentInTimetable(newAppointment *Appointment) {
	r.db.Create(&Appointment{TimetableId: newAppointment.TimetableId, ClientId: newAppointment.ClientId, Time: newAppointment.Time})
}

func (r *GormUserRepository) CheckAvailibleOfTimeInTimetable(time time.Time) bool {
	var appoint Appointment
	r.db.Where("time = ?", time).Find(&appoint)
	return appoint.ClientId == 0 // если запись не найдена то заполнится по умолч следовательно можно делать запись на это время
}

func (r *GormUserRepository) CheckThisClientCreateThisAppo(idClient, idAppointment string) bool {
	var appoint Appointment
	r.db.Where("client_id = ?", idClient).
		Where("id = ?", idAppointment).
		Find(&appoint)
	return appoint.ClientId != 0 // тру значит что это того чела
}

func (r *GormUserRepository) DeleteDataAppointment(idAppointment string) {
	r.db.Where("id = ?", idAppointment).Delete(&Appointment{})
}

// NOIDEI
func (r *GormUserRepository) GetFreeAppointments(idTimetable string) Timetable {
	var timeTable Timetable

	r.db.Where("id = ?", idTimetable).
		Find(&timeTable)

	return timeTable
}
