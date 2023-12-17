package main

import (
	"github.com/beintil/user_service/api/base/db"
	"github.com/beintil/user_service/pkg/logger"
	lue "github.com/beintil/user_service/pkg/lu_error"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	driver, err := postgres.WithInstance(db.GetDB().DB, &postgres.Config{})
	if err != nil {
		logger.Error(lue.New("FAILED MIGRATION"+err.Error(), lue.ErrorMigrate))
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./api/base/db/migrations/files_migration",
		"postgres", driver)
	if err != nil {
		logger.Error(lue.New("FAILED MIGRATION"+err.Error(), lue.ErrorMigrate))
		return
	}
	err = m.Up()
	if err != nil {
		logger.Error(lue.New("FAILED MIGRATION UP"+err.Error(), lue.ErrorMigrate))
		return
	}
	logger.Info("MIGRATION SUCCESSFULLY")
}
