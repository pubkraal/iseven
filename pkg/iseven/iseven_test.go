package iseven

import "testing"

func TestIsEven(t *testing.T) {
	tables := []struct {
		given  int
		expect bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{3, false},
		{7, false},
		{-1, false},
		{-3, false},
		{-2, true},
		{12033482393, false},
		{10000000002, true},
		{9007199254740991, false},
	}

	for _, row := range tables {
		res := IsEven(row.given)
		if res != row.expect {
			t.Errorf("IsEven(%v) == %v. Expected %v", row.given, res, row.expect)
		}
	}
}
