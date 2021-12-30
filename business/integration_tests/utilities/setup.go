package utilities

import (
	"log"

	"github.com/cocus_challenger_refact/platform/config"
	db "github.com/cocus_challenger_refact/platform/db_connect"
	"github.com/cocus_challenger_refact/platform/migrations"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	BaseURL             = ""
	Db_driver           = ""
	Db_source           = ""
	User                = "cocus"
	Password            = "cocus"
	AuhtorizationHeader = "authorization"
	Token               = ""
)

func Setup() {

	config, err := config.LoadConfig("../../platform")
	if err != nil {
		log.Printf("cannot load config:", err)
	}

	log.Println(config.DBSource)
	BaseURL = config.ServerAddress
	Db_driver = config.DBDriver
	Db_source = config.DBSource

	// Start DB connection
	dbconnection := db.InitDB(Db_driver, Db_source)
	//Starting migrations and restarting tables
	migrations.InitLocationMigrations("../..")
	migrations.InitMigrations(dbconnection)

	defer dbconnection.Close()
}
