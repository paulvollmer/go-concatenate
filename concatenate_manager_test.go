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

func Test_Manager(t *testing.T) {
	testCases := []struct {
		config Config
	}{
		{
			config: Config{
				"test-1": []string{"src-1"},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {

			m := NewManager()
			m.Config = tc.config

			// if m.TotalFilesInSet("tmp_test1.txt") != 2 {
			// 	t.Error("TotalFilesInSet not equal")
			// }
			// if m.TotalFiles() != 2 {
			// 	t.Error("TotalFiles not equal")
			// }
		})
	}

}

// 	err := m.Process("tmp_test1.txt", 0777)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

func Test_Manager_ProcessAll(t *testing.T) {
	// m := NewManager()
	// m.Add("tmp_test2.txt", inputFiles...)
	//
	// err := m.ProcessAll(0777)
	// if err != nil {
	// 	t.Error(err)
	// }
}
