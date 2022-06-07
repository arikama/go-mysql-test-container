package util_test

import (
	"testing"

	"github.com/arikama/go-mysql-test-container/util"
)

func TestGetFiles(t *testing.T) {
	files, err := util.GetFiles("./../migration/migration")
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}
	expected := 5
	if len(files) != expected {
		t.Errorf("Want=%v, got=%v\n", expected, len(files))
	}
}

func TestGetFilesMissing(t *testing.T) {
	_, err := util.GetFiles("./../migration/missing")
	if err != nil {
		expected := "open ./../migration/missing: no such file or directory"
		if err.Error() != expected {
			t.Errorf(expected, err.Error())
		}
	}
}
