package migration

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func Test(t *testing.T) {
	_, err := mysqltestcontainer.Start("test", "./resources/db/migration")
	if err != nil {
		t.Errorf("Error starting MySQL test container: %v\n", err.Error())
	}
}
