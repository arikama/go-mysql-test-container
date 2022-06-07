package mysqltestcontainer_test

import (
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func ExampleStart() {
	db, _ := mysqltestcontainer.Start("test", "./../migration")
	db.Ping()
}
