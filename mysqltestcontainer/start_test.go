package mysqltestcontainer_test

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func TestStart(t *testing.T) {
	result, err := mysqltestcontainer.Start("test", "")
	if err != nil {
		t.Errorf("Failed to start: %v", err.Error())
	}
	err = result.Db.Ping()
	if err != nil {
		t.Errorf("Ping failed: %v", err.Error())
	}
}

func TestStartExample(t *testing.T) {
	result, err := mysqltestcontainer.Start("test", "./../migration/example")
	if err != nil {
		t.Errorf("Got=%v\n", err.Error())
	}
	db := result.Db
	names := []string{"Clare", "Teresa", "Priscilla"}
	for _, name := range names {
		db.Exec(`INSERT INTO user (username, name) VALUES (?, ?);`, name, name)
	}
	rows, _ := db.Query("SELECT id, name FROM user ORDER BY id ASC;")
	for i := 0; i < len(names); i++ {
		rows.Next()
		var id int
		var name string
		rows.Scan(&id, &name)
		if i+1 != id {
			t.Errorf("Expected %v, got %v\n", i+1, id)
		}
		if names[i] != name {
			t.Errorf("Expected %v, got %v\n", names[i], name)
		}
	}
}

func TestStartMigration(t *testing.T) {
	_, err := mysqltestcontainer.Start("test", "./../migration/migration")
	if err != nil {
		t.Errorf("Error starting MySQL test container: %v\n", err.Error())
	}
}

func TestStartMissing(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic to happen when migration files are missing")
		}
	}()
	mysqltestcontainer.Start("test", "./../migration/missing")
}

func TestStartInvalid(t *testing.T) {
	_, err := mysqltestcontainer.Start("test", "./../migration/invalid")
	if err == nil {
		panic("This should not happen.")
	}
}
