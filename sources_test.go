package concatenate

import (
	"testing"
)

func Test_Sources(t *testing.T) {
	m := NewSources()
	if m == nil {
		t.Error("NewSources not type of Sources")
	}

	if m.Total() != 0 {
		t.Error("Sources Total not equal")
	}

	addOk := m.Add("a.txt")
	if addOk == false {
		t.Error("Add failed")
	}
	if (*m)[0] != "a.txt" {
		t.Error("Add not set to string array")
	}
	// att the same source a second time
	addOk = m.Add("a.txt")
	if addOk == true {
		t.Error("Add must be fasle")
	}
}
