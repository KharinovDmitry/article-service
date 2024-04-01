package migrator

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"log"
)

func MustRun(connStr, driverName, migrationsDir string) {
	log.Println("Начало миграции")

	db, err := sql.Open(driverName, connStr)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	if err := goose.SetDialect(driverName); err != nil {
		panic(err.Error())
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		//if !errors.Is(err, goose.ErrNoNextVersion) {
		panic(err.Error())
		//}
	}

	log.Println("Конец миграции")
}
