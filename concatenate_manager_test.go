package concatenate

import (
	"fmt"
	"testing"
)

func Test_NewManager(t *testing.T) {
	m := NewManager()
	if m == nil {
		t.Error("NewManager not type of Manager")
	}
}

func Test_Manager_ReadConfig(t *testing.T) {
	m := NewManager()
	err := m.ReadConfig("./fixture/sample.json")
	if err != nil {
		t.Error(err)
	}
	if m.Config.TotalSets() != 1 {
		t.Error("Total sets not equal. must be 1")
	}

	m2 := NewManager()
	err = m2.ReadConfig("./fixture/file_not_exist.json")
	if err == nil {
		t.Error("Missing Error")
	}
}

func Test_Manager_ProcessAll(t *testing.T) {
	testCases := []struct {
		config Config
	}{
		{
			config: Config{
				"tmp_test_1.txt": []string{"./fixture/a.txt", "./fixture/b.txt"},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {

			m := NewManager()
			m.Config = tc.config
			err := m.ProcessAll(0755)
			if err != nil {
				t.Error(err)
			}

			total, err := m.Config.TotalFilesInSet("tmp_test_1.txt")
			if err != nil {
				t.Error(err)
			}
			if total != 2 {
				t.Error("TotalFilesInSet not equal")
			}
		})
	}
}
