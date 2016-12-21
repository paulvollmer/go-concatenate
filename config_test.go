package concatenate

import (
	"fmt"
	"testing"
)

func Test_NewConfig(t *testing.T) {
	cfg := NewConfig()
	if cfg == nil {
		t.Error("Config not type of Config")
	}
}

func Test_Config(t *testing.T) {
	testCases := []struct {
		data       map[string][]string
		totalFiles int
	}{
		{
			data: map[string][]string{
				"test-1": {"fixture/a.txt"},
			},
			totalFiles: 1,
		},
		{
			data: map[string][]string{
				"test-1": {"fixture/a.txt", "fixture/b.txt"},
			},
			totalFiles: 2,
		},
		// {
		// 	name:       "test-3",
		// 	sources:    []string{"fixture/c/*"},
		// 	totalFiles: 3,
		// },
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			cfg := NewConfig()
			for name, src := range tc.data {
				cfg.AddSet(name, src...)
			}

			if cfg.TotalSets() != 1 {
				t.Error("TotalSets not equal")
			}

			totalFiles, err := cfg.TotalFiles()
			if err != nil {
				t.Error(err)
			}
			if totalFiles != tc.totalFiles {
				t.Error("TotalFiles not equal")
			}

			// if cfg.TotalFilesInSet(tc.name) != tc.totalFiles {
			// 	t.Error("TotalFilesInSet not equal")
			// }
		})
	}
}
