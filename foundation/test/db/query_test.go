package db

import (
	"github.com/confetti-framework/framework/foundation"
	"github.com/confetti-framework/framework/foundation/db"
	"github.com/confetti-framework/framework/foundation/providers"
	"github.com/confetti-framework/framework/inter"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

type Product struct {
	gorm.Model
	Name  string
	Price int
}

func Test_db_migrate_create_select(t *testing.T) {
	app := setupDB()
	qb := app.Db("sqlite")

	// Migrate the schema
	qb.AutoMigrate(&Product{})

	// Create
	qb.Create(&Product{Name: "iPhone 12 Pro", Price: 1300})

	// Read
	var product Product
	qb.First(&product, 1)

	assert.Equal(t, "iPhone 12 Pro", product.Name)
	assert.Equal(t, 1300, product.Price)
}

func setupDB() inter.App {
	app := foundation.NewApp()

	provider := providers.DatabaseServiceProvider{Connections: map[string]inter.Connection{
		"sqlite": &db.Sqlite{
			Dsn: "file::memory:?cache=shared",
		},
	}}

	provider.Boot(*app.Container())

	return app
}
