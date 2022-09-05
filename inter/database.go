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
	AutoMigrate(dst ...interface{}) error
	Create(value interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}

type TypeCast func(ct sql.ColumnType, raw []byte) interface{}
