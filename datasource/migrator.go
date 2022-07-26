package datasource

import (
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"BecauseLanguageBot/datasource/migrationData"
)

func (source *DataSource) migrate(development bool) error {
	var m *migrate.Migrate
	var err error

	if development {
		m, err = migrate.New("file://datasource/migrations", source.connectionString)
	} else {
		s := bindata.Resource(migrationData.AssetNames(),
			func(name string) ([]byte, error) {
				return migrationData.Asset(name)
			})
		d, err := bindata.WithInstance(s)
		if err != nil {
			return err
		}
		m, err = migrate.NewWithSourceInstance("go-bindata", d, source.connectionString)
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
