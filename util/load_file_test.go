package util_test

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/util"
)

func TestLoadFile(t *testing.T) {
	content, err := util.LoadFile("./../migration/example/V11_alter_user_columns.sql")
	if err != nil {
		panic(err)
	}
	expected := "ALTER TABLE `user` ADD COLUMN `username` VARCHAR(64) NOT NULL UNIQUE AFTER `id`;\nALTER TABLE `user` ADD COLUMN `name` VARCHAR(64) NOT NULL UNIQUE AFTER `username`;"
	if content != expected {
		t.Errorf("Got=%v, want=%v\n", content, expected)
	}
}

func TestLoadFileMissing(t *testing.T) {
	content, err := util.LoadFile("./../migration/example/missing.sql")
	if err != nil && err.Error() != "open ./../migration/example/missing.sql: no such file or directory" {
		panic(err)
	}
	expected := ""
	if content != expected {
		t.Errorf("Got=%v, want=%v\n", content, expected)
	}
}
