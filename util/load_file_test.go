package util

import "testing"

func TestLoadFile(t *testing.T) {
	content, err := LoadFile("./../migration/example/V11__alter_table.sql")
	if err != nil {
		panic(err)
	}
	expected := "ALTER TABLE `test` ADD COLUMN `name` VARCHAR(64) NOT NULL UNIQUE AFTER `id`;"
	if content != expected {
		t.Errorf("Got=%v, want=%v\n", content, expected)
	}
}

func TestLoadFileMissing(t *testing.T) {
	content, err := LoadFile("./../migration/example/missing.sql")
	if err != nil && err.Error() != "open ./../migration/example/missing.sql: no such file or directory" {
		panic(err)
	}
	expected := ""
	if content != expected {
		t.Errorf("Got=%v, want=%v\n", content, expected)
	}
}
