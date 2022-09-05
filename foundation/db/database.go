package db

import (
	"context"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/framework/inter"
	"github.com/confetti-framework/framework/support"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
	connection inter.Connection
	app        inter.App
	tx         *gorm.DB
}

func NewDatabase(app inter.App, connection inter.Connection) inter.Database {
	return &Database{app: app, connection: connection, DB: connection.DB()}
}

func (d Database) Connection() inter.Connection {
	return d.connection
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

	d.WithContext(ctx)

	rows, err := d.Rows()
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
