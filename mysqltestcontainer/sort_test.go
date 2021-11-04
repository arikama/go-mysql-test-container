package mysqltestcontainer

import (
	"sort"
	"testing"
)

func TestByFileVersion(t *testing.T) {
	files := byFileVersion{
		"V00090_test.sql",
		"V00080_test.sql",
		"V012_test.sql",
		"V11_test.sql",
		"V1_test.sql",
		"V000_test.sql",
	}
	expected := []string{
		"V000_test.sql",
		"V1_test.sql",
		"V11_test.sql",
		"V012_test.sql",
		"V00080_test.sql",
		"V00090_test.sql",
	}
	sort.Sort(files)
	for i := 0; i < len(expected); i++ {
		if files[i] != expected[i] {
			t.Errorf("Got %v, expected %v\n", files[i], expected[i])
		}
	}
}
