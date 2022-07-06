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
	Tx() *gorm.DB
	Table(name string, args ...interface{}) Database
	First(dest interface{}) interface{}
	Where(query interface{}, args ...interface{}) Database
	Exec(sql string, args ...interface{}) Database
	Raw(sql string, args ...interface{}) Database
	Error() error
	RowsAffected() int64
}

type TypeCast func(ct sql.ColumnType, raw []byte) interface{}
