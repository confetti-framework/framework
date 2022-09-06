package inter

import (
	"context"
	"database/sql"
	"github.com/confetti-framework/framework/support"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	Session(config *gorm.Session) *gorm.DB
	WithContext(ctx context.Context) *gorm.DB
	Debug() (tx *gorm.DB)
	Set(key string, value interface{}) *gorm.DB
	Get() support.Collection
	GetE() (support.Collection, error)
	InstanceSet(key string, value interface{}) *gorm.DB
	InstanceGet(key string) (interface{}, bool)
	AddError(err error) error
	SetupJoinTable(model interface{}, field string, joinTable interface{}) error
	Use(plugin gorm.Plugin) error
	ToSQL(queryFn func(tx *gorm.DB) *gorm.DB) string
	Association(column string) *gorm.Association
	Create(value interface{}) (tx *gorm.DB)
	CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB)
	Save(value interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Take(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Last(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB
	FirstOrInit(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Update(column string, value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	UpdateColumn(column string, value interface{}) (tx *gorm.DB)
	UpdateColumns(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Count(count *int64) (tx *gorm.DB)
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Scan(dest interface{}) (tx *gorm.DB)
	Pluck(column string, dest interface{}) (tx *gorm.DB)
	ScanRows(rows *sql.Rows, dest interface{}) error
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
	Begin(opts ...*sql.TxOptions) *gorm.DB
	Commit() *gorm.DB
	Rollback() *gorm.DB
	SavePoint(name string) *gorm.DB
	RollbackTo(name string) *gorm.DB
	Exec(sql string, values ...interface{}) (tx *gorm.DB)
	Migrator() gorm.Migrator
	AutoMigrate(dst ...interface{}) error
	Model(value interface{}) (tx *gorm.DB)
	Clauses(conds ...clause.Expression) (tx *gorm.DB)
	Table(name string, args ...interface{}) (tx *gorm.DB)
	Distinct(args ...interface{}) (tx *gorm.DB)
	Select(query interface{}, args ...interface{}) (tx *gorm.DB)
	Omit(columns ...string) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	Not(query interface{}, args ...interface{}) (tx *gorm.DB)
	Or(query interface{}, args ...interface{}) (tx *gorm.DB)
	Joins(query string, args ...interface{}) (tx *gorm.DB)
	Group(name string) (tx *gorm.DB)
	Having(query interface{}, args ...interface{}) (tx *gorm.DB)
	Order(value interface{}) (tx *gorm.DB)
	Limit(limit int) (tx *gorm.DB)
	Offset(offset int) (tx *gorm.DB)
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)
	Attrs(attrs ...interface{}) (tx *gorm.DB)
	Assign(attrs ...interface{}) (tx *gorm.DB)
	Unscoped() (tx *gorm.DB)
	Raw(sql string, values ...interface{}) (tx *gorm.DB)
}

type TypeCast func(ct sql.ColumnType, raw []byte) interface{}
