package datasource

import (
	"database/sql"
	"fmt"
	"gopkg.in/errgo.v2/errors"

	_ "github.com/mattn/go-sqlite3"

	"BecauseLanguageBot/config"
)

type DataSource struct {
	connectionString string
	connection       *sql.DB
}

func Init(config config.DatabaseConfig, development bool) (*DataSource, error) {
	source := DataSource{
		connectionString: fmt.Sprintf("sqlite3://%s", config.Path),
	}

	err := source.migrate(development)
	if err != nil {
		return nil, errors.Because(err, nil, "unable to migrate database")
	}

	source.connection, err = sql.Open("sqlite3", source.connectionString)
	if err != nil {
		return nil, errors.Because(err, nil, "unable to open database")
	}

	return &source, nil
}
