package mysqltestcontainer_test

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"testing"

	"github.com/andylongstaffe/go-mysql-test-container/mysqltestcontainer"
	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	// Create container
	mySql, err := mysqltestcontainer.Create("test")
	assert.Nil(t, err)
	err = mySql.GetDb().Ping()
	assert.Nil(t, err)
	defer mySql.GetContainer().Terminate(context.TODO())

	// Migrate schema
	driver, err := mysql.WithInstance(mySql.GetDb(), &mysql.Config{})
	assert.Nil(t, err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql", driver)
	assert.Nil(t, err)

	err = m.Up()
	assert.Nil(t, err)

	// check schema here
}
