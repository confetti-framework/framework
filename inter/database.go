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

type TypeCast func(ct sql.ColumnType, raw []byte) interface{}
