package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

// if no value was found on the first two(here one) steps, the field will be filled with the default value
// can if have field env-description: use command cleanenv.GetDescription
type ConfigDatabase struct {
	PG_host     string `yaml:"PG_host" env-default:"localhost"`
	PG_port     string `yaml:"PG_port" env-default:"5432"`
	PG_user     string `yaml:"PG_user" env-default:"postgres"`
	PG_db_name  string `yaml:"PG_db_name" env-default:"postgres"`
	PG_password string `yaml:"PG_password" env-default:"1"`
}

func LaunchConfigFile() ConfigDatabase {
	var cfg ConfigDatabase

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal("error: ", err)
	}
	return cfg
}
