package database

import (
	"fmt"
	"net/url"

	"github.com/Synaxis/battlefieldHeroesBackend/backend/config"
)

// newMySQL creates a DSN string to connect to MySQL database
func newMySQL() (string, string) {
	connParams := url.Values{
		"charset":   {"utf8"},
		"parseTime": {"True"}, // https://github.com/go-sql-driver/mysql/issues/9
		"loc":       {"UTC"},  // Timezone
	}

	dsnAddr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		config.Database.DatabaseUserName,
		config.Database.DatabasePassword,
		config.Database.DatabaseHost,
		config.Database.DatabasePort,
		config.Database.DatabaseName,
		connParams.Encode(),
	)

	return "mysql", dsnAddr
}
