package psql

import (
	"database/sql"
	"strings"

	"compass/logger"
	"compass/store/psql/migrations"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"github.com/rs/zerolog"

	bindata "github.com/mattes/migrate/source/go-bindata"
)

// migrateLogger is used to log database migrations
type MigrateLogger struct {
	logger zerolog.Logger
}

// Printf passes the log message and arguments to our logger
func (l *MigrateLogger) Printf(format string, v ...interface{}) {
	l.logger.Printf(strings.TrimSpace(format), v...)
}

// Verbose returns true so we can log what the migrator is doing
func (l *MigrateLogger) Verbose() bool {
	return true
}

// Migrate runs database migrat
func NewMigrator(db *sql.DB) (*migrate.Migrate, error) {
	log := logger.New()
	// Get migration files from go-bindata codegen
	resource := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			log.Debug().Str("migration", name).Msg("load migration")
			return migrations.Asset(name)
		})
	// Generate a source from the bindata
	source, err := bindata.WithInstance(resource)
	if err != nil {
		return nil, err
	}
	// Get database driver from connection
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	// Create migrator using our data source and db driver
	m, err := migrate.NewWithInstance("go-bindata", source, "postgres", driver)
	if err != nil {
		return nil, err
	}
	m.Log = &MigrateLogger{log}
	return m, nil
}