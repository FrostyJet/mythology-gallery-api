package setting

import "os"

type Database struct {
	Host     string
	Password string
	Name     string
	User     string
	Port     string
}

var DatabaseSetting = &Database{}

func Setup() {

	// Init databse settings
	DatabaseSetting.Host = os.Getenv("POSTGRES_DB_HOST")
	DatabaseSetting.Password = os.Getenv("POSTGRES_PASSWORD")
	DatabaseSetting.Name = os.Getenv("POSTGRES_DB")
	DatabaseSetting.User = os.Getenv("POSTGRES_USER")
	DatabaseSetting.Port = os.Getenv("POSTGRES_DB_PORT")
}
