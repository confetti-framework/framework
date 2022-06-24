package db

import (
	"context"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/framework/inter"
	"github.com/confetti-framework/framework/support"
	"gorm.io/gorm"
)

type Database struct {
	connection inter.Connection
	app        inter.App
	tx         *gorm.DB
}

func NewDatabase(app inter.App, connection inter.Connection) *Database {
	return &Database{app: app, connection: connection, tx: connection.DB()}
}

func (d Database) Connection() inter.Connection {
	return d.connection
}

func (d Database) context() (context.Context, context.CancelFunc) {
	connection := d.Connection()
	source := d.app.Make("request").(inter.Request).Source()

	ctx, cancel := context.WithTimeout(source.Context(), connection.Timeout())

	return ctx, cancel
}

func (d Database) Error() error {
	return d.tx.Error
}

func (d Database) RowsAffected() int64 {
	return d.tx.RowsAffected
}

func (d Database) Table(name string, args ...interface{}) inter.Database {
	ctx, cancel := d.context()
	defer cancel()

	d.tx = d.tx.WithContext(ctx).Table(name)

	return d
}

func (d Database) First(dest interface{}) inter.Database {
	connection := d.Connection()
	source := d.app.Make("request").(inter.Request).Source()

	ctx, cancel := context.WithTimeout(source.Context(), connection.Timeout())
	defer cancel()

	d.tx = d.tx.WithContext(ctx).First(dest)

	return d
}

func (d Database) Where(query interface{}, args ...interface{}) inter.Database {
	ctx, cancel := d.context()
	defer cancel()

	d.tx = d.tx.WithContext(ctx).Where(query, args...)

	return d
}

func (d Database) Count() int64 {
	ctx, cancel := d.context()
	defer cancel()

	var count int64
	d.tx = d.tx.WithContext(ctx).Count(&count)

	return count
}

func (d Database) Create(value interface{}) inter.Database {
	ctx, cancel := d.context()
	defer cancel()

	d.tx = d.tx.WithContext(ctx).Create(value)

	return d
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d Database) Exec(sql string, args ...interface{}) inter.Database {
	ctx, cancel := d.context()
	defer cancel()

	d.tx = d.tx.WithContext(ctx).Exec(sql, args...)

	return d
}

// Raw executes a query that returns rows or an error, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (d Database) Raw(sql string, args ...interface{}) inter.Database {
	ctx, cancel := d.context()
	defer cancel()

	d.tx = d.tx.WithContext(ctx).Raw(sql, args...)

	return d
}

func (d Database) Get() support.Collection {
	result, err := d.GetE()
	if err != nil {
		panic(err)
	}
	return result
}

func (d Database) GetE() (support.Collection, error) {
	result := support.NewCollection()

	ctx, cancel := d.context()
	defer cancel()

	d.tx = d.tx.WithContext(ctx)

	rows, err := d.tx.Rows()
	defer rows.Close()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		err := rows.Scan(columnPointers...)
		if err != nil {
			return result, errors.WithStack(err)
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		result = result.Push(m)
	}

	return result, nil
}
