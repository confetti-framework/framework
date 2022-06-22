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
	DB() *gorm.DB
	Pool() *sql.DB
	Timeout() time.Duration
}

type Database interface {
	Connection() Connection
	Get() support.Collection
	GetE() (support.Collection, error)
	Table(name string, args ...interface{}) Database
	Where(query interface{}, args ...interface{}) Database
	Exec(sql string, args ...interface{}) Database
	ExecE(sql string, args ...interface{}) (Database, error)
	Raw(sql string, args ...interface{}) support.Collection
	RawE(sql string, args ...interface{}) (support.Collection, error)
}

type TypeCast func(ct sql.ColumnType, raw []byte) interface{}
