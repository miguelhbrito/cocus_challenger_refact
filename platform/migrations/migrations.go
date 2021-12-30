package migrations

import (
	"bufio"
	"database/sql"
	"os"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"fmt"
)

var (
	locationMigrations = ""
)

func InitMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		locationMigrations,
		"postgres", driver)
	if err != nil {
		panic(err)
	}

	_ = m.Down()
	_ = m.Up()

	fmt.Println("Successfully migrations applied!")
}

func InitLocationMigrations(dir string) {
	if runtime.GOOS == "windows" {
		fmt.Print("Windows OS detected, please enter project path(example: C:/Users/username/Documents/dev/github/):")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		locationMigrations = fmt.Sprintf("file://" + text + "cocus_challenger_refact/platform/migrations/")
	} else {

		if dir != "" {
			err := os.Chdir(dir)
			if err != nil {
				panic(err)
			}
		}

		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		locationMigrations = fmt.Sprintf("file://%s/platform/migrations/", pwd)
	}
}
