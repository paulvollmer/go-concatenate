package concatenate

import (
	"testing"
)

var inputFiles = []string{"fixture/a.txt", "fixture/b.txt"}

func Test_Manager_Process(t *testing.T) {
	m := NewManager()
	m.Set("tmp_test1.txt", inputFiles...)

	err := m.Process("tmp_test1.txt")
	if err != nil {
		t.Error(err)
	}
}

func Test_Manager_ProcessAll(t *testing.T) {
	m := NewManager()
	m.Set("tmp_test2.txt", inputFiles...)

	err := m.ProcessAll()
	if err != nil {
		t.Error(err)
	}
}
