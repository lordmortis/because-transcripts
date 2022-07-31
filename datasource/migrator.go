package datasource

import (
	"BecauseLanguageBot/datasource/migrationData"
	"fmt"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"

	"BecauseLanguageBot/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func PerformMigrations(config config.DatabaseConfig, development bool) error {
	var connString = fmt.Sprintf("sqlite3://%s", config.Path)

	var m *migrate.Migrate
	var err error

	if development {
		m, err = migrate.New("file://datasource/migrations", connString)
	} else {
		s := bindata.Resource(migrationData.AssetNames(),
			func(name string) ([]byte, error) {
				return migrationData.Asset(name)
			})
		d, err := bindata.WithInstance(s)
		if err != nil {
			return err
		}
		m, err = migrate.NewWithSourceInstance("go-bindata", d, connString)
	}

	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	srcErr, dstErr := m.Close()

	if srcErr != nil {
		return srcErr
	}

	if dstErr != nil {
		return dstErr
	}

	return nil
}
