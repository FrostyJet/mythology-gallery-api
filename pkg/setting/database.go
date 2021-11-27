package setting

import "os"

type Database struct {
	Host     string
	Password string
	Name     string
	User     string
	Port     string
}

func (d *Database) Setup() *Database {
	// Init databse settings
	d.Host = os.Getenv("POSTGRES_DB_HOST")
	d.Password = os.Getenv("POSTGRES_PASSWORD")
	d.Name = os.Getenv("POSTGRES_DB")
	d.User = os.Getenv("POSTGRES_USER")
	d.Port = os.Getenv("POSTGRES_DB_PORT")

	return d
}
