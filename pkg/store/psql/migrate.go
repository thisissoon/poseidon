package psql

import (
	"strings"

	"github.com/mattes/migrate"
	"github.com/rs/zerolog"

	bindata "github.com/mattes/migrate/source/go-bindata"

	// postgres driver
	_ "github.com/mattes/migrate/database/postgres"
)

// migrateLogger is used to log database migrations
type MigrateLogger struct {
	Logger zerolog.Logger
}

// Printf passes the log message and arguments to our logger
func (l *MigrateLogger) Printf(format string, v ...interface{}) {
	l.Logger.Printf(strings.TrimSpace(format), v...)
}

// Verbose returns true so we can log what the migrator is doing
func (l *MigrateLogger) Verbose() bool {
	return true
}

// Migrate runs database migrat
func NewMigrator(dsn DSN, log *MigrateLogger) (*migrate.Migrate, error) {
	// Get migration files from go-bindata codegen
	resource := bindata.Resource(AssetNames(),
		func(name string) ([]byte, error) {
			return Asset(name)
		})
	// Generate a source from the bindata
	source, err := bindata.WithInstance(resource)
	if err != nil {
		return nil, err
	}
	// Create migrator using our data source and db driver
	m, err := migrate.NewWithSourceInstance("go-bindata", source, dsn.String())
	if err != nil {
		return nil, err
	}
	m.Log = log
	return m, nil
}