package mysqltestcontainer_test

import (
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func ExampleCreate() {
	mySql, _ := mysqltestcontainer.Create("test")
	mySql.Db.Ping()
}
