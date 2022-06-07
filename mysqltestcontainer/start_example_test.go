package mysqltestcontainer_test

import (
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func ExampleStart() {
	result, _ := mysqltestcontainer.Start("test", "./../migration/example")
	result.Db.Ping()
}
