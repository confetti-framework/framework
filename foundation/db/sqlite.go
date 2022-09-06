package db

import (
	"database/sql"
	"github.com/confetti-framework/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Sqlite struct {
	Dsn string

	// When the open connection limit is reached, and all connections are in-use,
	// any new database tasks that your application needs to execute will be forced
	// to wait until a connection becomes free and marked as idle. To mitigate this
	// you can set a fixed, fast, timeout when making database calls.
	// Default 10 seconds
	QueryTimeout time.Duration

	// Set the maximum lifetime of a connection to 1 hour. Setting it to 0 means
	// that there is no maximum lifetime and the connection is reused forever.
	// Default 5 minutes
	ConnMaxLifetime time.Duration

	// Set the maximum number of concurrently open connections (in-use + idle)
	// Setting this to less than 0 will mean there is no maximum limit.
	// Default 25
	MaxOpenConnections int

	// Set the maximum number of concurrently idle connections. Setting this
	// to less than 0 will mean that no idle connections are retained.
	// Default MaxOpenConnections
	MaxIdleConnections int

	// pool is a database handle representing a pool of zero or more
	// underlying connections. It's safe for concurrent use by multiple
	// goroutines.
	pool *sql.DB

	db *gorm.DB
}

func (m *Sqlite) Open() error {
	connection, err := gorm.Open(sqlite.Open(m.Dsn), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "can't open MySQL connection")
	}

	pool, err := connection.DB()

	ConfigConnection(pool, m.ConnMaxLifetime, m.MaxOpenConnections, m.MaxIdleConnections)

	m.db = connection
	m.pool = pool

	return err
}

func (m *Sqlite) Pool() *sql.DB {
	return m.pool
}

func (m *Sqlite) DB() *gorm.DB {
	return m.db
}

func (m *Sqlite) Timeout() time.Duration {
	return m.QueryTimeout
}
