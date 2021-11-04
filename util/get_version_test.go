package util

import "testing"

func TestGetVersion(t *testing.T) {
	v := GetVersion("/Users/1/workspace/mysql-test-container/integrationtest/example/resources/db/migration/V1__create_table.sql")
	if v != 1 {
		t.Errorf("Got %v, expected %v\n", v, 1)
	}
}
