package mysqltestcontainer_test

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func TestStart(t *testing.T) {
	result, err := mysqltestcontainer.Start("test")
	if err != nil {
		t.Errorf("Failed to start: %v", err.Error())
	}
	err = result.Db.Ping()
	if err != nil {
		t.Errorf("Ping failed: %v", err.Error())
	}
}
