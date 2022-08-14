package datasource

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gopkg.in/errgo.v2/errors"

	_ "github.com/mattn/go-sqlite3"

	"BecauseLanguageBot/config"
)

type DataSource struct {
	dbPath           string
	connectionString string
	connection       *sql.DB
}

func Init(config config.DatabaseConfig, development bool) (*DataSource, error) {
	var err error
	source := DataSource{
		dbPath:           config.Path,
		connectionString: fmt.Sprintf("sqlite3://%s", config.Path),
	}

	err = source.migrate(development)
	if err != nil {
		return nil, errors.Because(err, nil, "unable to migrate database")
	}

	source.connection, err = sql.Open("sqlite3", source.dbPath)
	if err != nil {
		return nil, errors.Because(err, nil, "unable to open database")
	}

	_, err = source.connection.Exec("SELECT true")
	if err != nil {
		return nil, errors.Because(err, nil, "unable to open database")
	}

	if development {
		boil.DebugMode = true
	}

	return &source, nil
}
