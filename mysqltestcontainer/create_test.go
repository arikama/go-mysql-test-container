package mysqltestcontainer_test

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func TestCreate(t *testing.T) {
	mySql, err := mysqltestcontainer.Create("test")
	if err != nil {
		t.Errorf("Create failed: error=%v", err.Error())
	}
	err = mySql.Db.Ping()
	if err != nil {
		t.Errorf("Ping failed: error=%v", err.Error())
	}
}
