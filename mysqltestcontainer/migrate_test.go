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
	container, err := mysqltestcontainer.Create("test")
	assert.Nil(t, err)
	err = container.GetDb().Ping()
	assert.Nil(t, err)
	defer container.GetContainer().Terminate(context.TODO())

	// Migrate schema
	driver, err := mysql.WithInstance(container.GetDb(), &mysql.Config{})
	assert.Nil(t, err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql", driver)
	assert.Nil(t, err)

	err = m.Up()
	assert.Nil(t, err)

	// check schema here
	db := container.GetDb()
	_, err = db.Exec("insert into test(firstname) values (?)", "andy")
	assert.Nil(t, err)

	rows, err := db.Query("select firstname from test")
	assert.Nil(t, err)
	var name string
	if rows.Next() {
		assert.Nil(t, rows.Scan(&name))
	}
	assert.Equal(t, "andy", name)

}
