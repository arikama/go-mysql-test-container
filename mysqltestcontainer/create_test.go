package mysqltestcontainer_test

import (
	"context"
	"github.com/andylongstaffe/go-mysql-test-container/mysqltestcontainer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mySql, err := mysqltestcontainer.Create("test")

	assert.Nil(t, err)
	defer mySql.GetContainer().Terminate(context.Background())

	assert.NotNil(t, mySql.GetDb())

	dbInfo := mySql.GetDbInfo()

	assert.NotEmpty(t, dbInfo.Ip)
	assert.NotEmpty(t, dbInfo.Port)
	assert.NotEmpty(t, dbInfo.Username)
	assert.NotEmpty(t, dbInfo.Password)
	assert.NotEmpty(t, dbInfo.DbName)

	err = mySql.GetDb().Ping()

	assert.Nil(t, err)
}

func TestCreateWithMigrate(t *testing.T) {

	container, err := mysqltestcontainer.CreateWithMigrate("test", "file://./migrations")
	assert.Nil(t, err)
	defer container.GetContainer().Terminate(context.Background())

	// check schema here
	firstname := "andy"
	db := container.GetDb()
	_, err = db.Exec("insert into test(firstname) values (?)", firstname)
	assert.Nil(t, err)

	rows, err := db.Query("select firstname from test")
	assert.Nil(t, err)
	var name string
	if rows.Next() {
		assert.Nil(t, rows.Scan(&name))
	}
	assert.Equal(t, firstname, name)

}
