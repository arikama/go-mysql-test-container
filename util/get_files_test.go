package util

import "testing"

func Test(t *testing.T) {
	files, err := GetFiles("./../migration/migration")
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}
	expected := 5
	if len(files) != expected {
		t.Errorf("Want=%v, got=%v\n", expected, len(files))
	}
}

func TestMissing(t *testing.T) {
	_, err := GetFiles("./../migration/missing")
	if err != nil {
		expected := "open ./../migration/missing: no such file or directory"
		if err.Error() != expected {
			t.Errorf(expected, err.Error())
		}
	}
}
