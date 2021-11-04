package mysqltestcontainer

import "testing"

func TestStart(t *testing.T) {
	db, err := Start("test", "")
	if err != nil {
		t.Errorf("Failed to start: %v", err.Error())
	}
	err = db.Ping()
	if err != nil {
		t.Errorf("Ping failed: %v", err.Error())
	}
}
