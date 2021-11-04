package mysqltestcontainer

import (
	"testing"
)

func TestExample(t *testing.T) {
	db, _ := Start("test", "./../migration/example")
	names := []string{"Clare", "Teresa", "Priscilla"}
	for _, name := range names {
		db.Exec(`INSERT INTO test (name) VALUES (?);`, name)
	}
	rows, _ := db.Query("SELECT id, name FROM test ORDER BY id ASC;")
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
