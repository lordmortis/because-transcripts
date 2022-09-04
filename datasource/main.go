package datasource

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gopkg.in/errgo.v2/errors"
	"strconv"

	_ "github.com/lib/pq"

	"BecauseLanguageBot/config"
)

type DataSource struct {
	connectionString string
	connection       *sql.DB
}

func Init(config config.DatabaseConfig, development bool) (*DataSource, error) {
	var err error

	var connString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Username, config.Password,
		config.Hostname, strconv.FormatUint(uint64(config.Port), 10),
		config.Database)

	source := DataSource{
		connectionString: connString,
	}

	err = source.migrate(development)
	if err != nil {
		return nil, errors.Because(err, nil, "unable to migrate database")
	}

	source.connection, err = sql.Open("postgres", source.connectionString)
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
