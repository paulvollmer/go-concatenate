package concatenate

import (
	"testing"
)

var inputFiles = []string{"fixture/a.txt", "fixture/b.txt"}

func Test_Manager_Process(t *testing.T) {
	m := NewManager()
	m.Set("tmp_test1.txt", inputFiles...)

	if m.TotalFilesInSet("tmp_test1.txt") != 2 {
		t.Error("TotalFilesInSet not equal")
	}
	if m.TotalFiles() != 2 {
		t.Error("TotalFiles not equal")
	}

	err := m.Process("tmp_test1.txt", 0755)
	if err != nil {
		t.Error(err)
	}
}

func Test_Manager_ProcessAll(t *testing.T) {
	m := NewManager()
	m.Set("tmp_test2.txt", inputFiles...)

	err := m.ProcessAll(0755)
	if err != nil {
		t.Error(err)
	}
}
