package mysqltestcontainer_test

import (
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func ExampleStart() {
	result, _ := mysqltestcontainer.Start("test")
	result.Db.Ping()
}
