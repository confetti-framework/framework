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

func (d Database) Table(name string, args ...interface{}) inter.Database {
	connection := d.Connection()
	source := d.app.Make("request").(inter.Request).Source()

	ctx, cancel := context.WithTimeout(source.Context(), connection.Timeout())
	defer cancel()

	d.tx = d.tx.WithContext(ctx).Table(name)

	return d
}

func (d Database) Where(query interface{}, args ...interface{}) inter.Database {
	connection := d.Connection()
	source := d.app.Make("request").(inter.Request).Source()

	ctx, cancel := context.WithTimeout(source.Context(), connection.Timeout())
	defer cancel()

	d.tx = d.tx.WithContext(ctx).Where(query, args...)
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

	connection := d.Connection()
	source := d.app.Make("request").(inter.Request).Source()

	ctx, cancel := context.WithTimeout(source.Context(), connection.Timeout())
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

func (d Database) Create(value interface{}) inter.Database {
	connection := d.Connection()
	source := d.app.Make("request").(inter.Request).Source()

	ctx, cancel := context.WithTimeout(source.Context(), connection.Timeout())
	defer cancel()

	d.tx = d.tx.WithContext(ctx).Create(value)

	return d
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d Database) Exec(sql string, args ...interface{}) inter.Database {
	result, err := d.ExecE(sql, args...)
	if err != nil {
		panic(err)
	}
	return result
}

// ExecE executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d Database) ExecE(sql string, args ...interface{}) (inter.Database, error) {
	connection := d.Connection()
	source := d.app.Make("request").(inter.Request).Source()

	ctx, cancel := context.WithTimeout(source.Context(), connection.Timeout())
	defer cancel()

	db := d.tx.WithContext(ctx)

	execContext := db.Exec(sql, args...)
	err := execContext.Error
	if err != nil {
		err = errors.WithMessage(errors.WithStack(err), "can't execute database query")
	}

	return d, err
}

// Raw executes a query that returns rows, typically a SELECT. The args are
// for any placeholder parameters in the query.
func (d Database) Raw(sql string, args ...interface{}) support.Collection {
	result, err := d.RawE(sql, args...)
	if err != nil {
		panic(err)
	}
	return result
}

// RawE executes a query that returns rows or an error, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (d Database) RawE(sql string, args ...interface{}) (support.Collection, error) {
	result := support.NewCollection()

	connection := d.Connection()
	source := d.app.Make("request").(inter.Request).Source()

	ctx, cancel := context.WithTimeout(source.Context(), connection.Timeout())
	defer cancel()

	db := d.tx.WithContext(ctx)

	rows, err := db.Raw(sql, args...).Rows()
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
