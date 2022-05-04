package inter

import (
	"database/sql"
	"github.com/confetti-framework/framework/support"
	"time"
)

type Cache interface {
	GetE(string) (support.Value, error)
}

type Connection interface {
	Open() error
	Pool() *sql.DB
	Timeout() time.Duration
}

type Database interface {
	Connection() Connection
	Exec(sql string, args ...interface{}) sql.Result
	ExecE(sql string, args ...interface{}) (sql.Result, error)
	Query(sql string, args ...interface{}) support.Collection
	QueryE(sql string, args ...interface{}) (support.Collection, error)
}

type TypeCast func(ct sql.ColumnType, raw []byte) interface{}
