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

	dirs := m.GetDirs()
	if len(dirs) != 1 {
		t.Error("GetDirs length not equal")
	}
	if m.GetDirs()[0] != "." {
		t.Error("GetDirs not equal")
	}
	if m.GetDirsOfSources()[0] != "fixture" {
		t.Error("GetDirsOfSources not equal")
	}

	err := m.Process("tmp_test1.txt", 0777)
	if err != nil {
		t.Error(err)
	}
}

func Test_Manager_ProcessAll(t *testing.T) {
	m := NewManager()
	m.Set("tmp_test2.txt", inputFiles...)

	err := m.ProcessAll(0777)
	if err != nil {
		t.Error(err)
	}
}
