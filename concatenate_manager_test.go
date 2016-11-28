package concatenate

import (
	"testing"
)

func Test_Manager(t *testing.T) {
	m := NewManager()
	m.Set("tmp_test2.txt", "fixture/a.txt", "fixture/b.txt")
	err := m.ProcessAll()
	if err != nil {
		t.Error(err)
	}
	m.Process("tmp_test2.txt")
}
