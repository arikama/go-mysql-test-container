package mysqltestcontainer_test

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mySql, err := mysqltestcontainer.Create("test")

	assert.Nil(t, err)
	assert.NotEmpty(t, mySql.Ip)
	assert.NotEmpty(t, mySql.Port)
	assert.NotEmpty(t, mySql.Username)
	assert.NotEmpty(t, mySql.Password)
	assert.NotEmpty(t, mySql.DbName)
	assert.NotNil(t, mySql.Db)

	err = mySql.Db.Ping()

	assert.Nil(t, err)
}
