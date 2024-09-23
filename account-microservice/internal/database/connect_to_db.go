package database

import (
	"fmt"
	"log"

	"wefdzen/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Cfg = config.LaunchConfigFile()

func Connect() (*gorm.DB, error) {
	//connect
	urlToDataBase := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", Cfg.PG_user, Cfg.PG_password, Cfg.PG_host, Cfg.PG_port, Cfg.PG_db_name)
	db, err := gorm.Open(postgres.Open(urlToDataBase), &gorm.Config{})
	if err != nil {
		log.Fatal("can't open database")
		return nil, err
	}
	db.AutoMigrate(&User{}) // если такой структуры небыло migrate will be create a new table
	//analog
	// 	create table users(
	// 	id serial primary key,
	// 	last_name text,
	// 	first_name text,
	// 	user_name text,
	// 	password text,
	// 	roles text[],
	// 	refresh_token text
	// );
	return db, nil
}
