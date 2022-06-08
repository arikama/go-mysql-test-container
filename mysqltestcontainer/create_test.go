package mysqltestcontainer_test

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mySql, err := mysqltestcontainer.Create("test")

	assert.Nil(t, err)
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
