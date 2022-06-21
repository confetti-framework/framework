package inter

import (
	"database/sql"
	"github.com/confetti-framework/framework/support"
	"gorm.io/gorm"
	"time"
)

type Cache interface {
	GetE(string) (support.Value, error)
}

type Connection interface {
	Open() error
	Pool() *sql.DB
	Dialector() gorm.Dialector
	DB() *gorm.DB
	Timeout() time.Duration
}

type Database interface {
	Connection() Connection
	Exec(sql string, args ...interface{}) *gorm.DB
	ExecE(sql string, args ...interface{}) (*gorm.DB, error)
	Query(sql string, args ...interface{}) support.Collection
	QueryE(sql string, args ...interface{}) (support.Collection, error)
}

type TypeCast func(ct sql.ColumnType, raw []byte) interface{}
