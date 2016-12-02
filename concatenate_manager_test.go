package concatenate

import (
	"testing"
)

func Test_NewManager(t *testing.T) {
	m := NewManager()
	if m == nil {
		t.Error("NewManager not type of Manager")
	}
}

func Test_Manager_Add(t *testing.T) {
	m := NewManager()
	tmpTarget := "tmp_test1.txt"
	added := m.Add(tmpTarget, []string{"hello.txt", "world.txt"}...)
	if added != true {
		t.Errorf("Add %q already exist", tmpTarget)
	}
	if len((*m)) != 1 {
		t.Error("Add total number of sets not equal")
	}
	if (*m)[tmpTarget][0] != "hello.txt" {
		t.Error("Add set not equal, must be 'hello'")
	}
	if (*m)[tmpTarget][1] != "world.txt" {
		t.Error("Add set not equal, must be")
	}

	// added = m.AddSource(tmpTarget, "foo.txt")
	// if added != true {
	// 	t.Errorf("AddSource %q already exist", tmpTarget)
	// }

}

var inputFiles = []string{"fixture/a.txt", "fixture/b.txt"}

func Test_Manager_Process(t *testing.T) {
	m := NewManager()
	added := m.Add("tmp_test1.txt", inputFiles...)
	if !added {
		t.Error("Add result not equal")
	}

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

	if !m.ExistSource("fixture/a.txt") {
		t.Error("ExistSource not equal, must exist")
	}
	if m.ExistSource("fixture/not_exist.txt") {
		t.Error("ExistSource not equal, must not exist")
	}

	err := m.Process("tmp_test1.txt", 0777)
	if err != nil {
		t.Error(err)
	}
}

func Test_Manager_ProcessAll(t *testing.T) {
	m := NewManager()
	m.Add("tmp_test2.txt", inputFiles...)

	err := m.ProcessAll(0777)
	if err != nil {
		t.Error(err)
	}
}
