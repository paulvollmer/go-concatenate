package concatenate

import (
	"fmt"
	"testing"
)

func Test_NewSources(t *testing.T) {
	m := NewSources()
	if m == nil {
		t.Error("NewSources not type of Sources")
	}
	if len(*m) != 0 {
		t.Error("NewSources not equal")
	}
}

func Test_Sources_Add(t *testing.T) {
	testCases := []struct {
		src           []string
		expectedOK    bool
		expectedTotal int
	}{
		{
			src:           []string{""},
			expectedOK:    false,
			expectedTotal: 0,
		},
		{
			src:           []string{"src-1"},
			expectedOK:    true,
			expectedTotal: 1,
		},
		{
			src:           []string{"fixture/a.txt", "fixture/b.txt"},
			expectedOK:    true,
			expectedTotal: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q", tc.src), func(t *testing.T) {

			m := NewSources()

			// test to Add() sources and check if the result is equal
			for addI, addV := range tc.src {
				ok := m.Add(addV)
				if ok != tc.expectedOK {
					t.Errorf("Add %v failed", addI)
				}
				// if the total expected size is higher than zero, we can check the array items
				if tc.expectedTotal > 0 {
					if (*m)[addI] != addV {
						t.Errorf("Add source 0 not equal, must be %v", addV)
					}

					// // att the same source a second time
					// addOk = m.Add("a.txt")
					// if addOk == true {
					// 	t.Error("Add must be fasle")
					// }

					// if len((*m)) != 1 {
					// 	t.Error("Add total number of sets not equal")
					// }
					// if (*m)[tmpTarget][0] != "hello.txt" {
					// 	t.Error("Add set not equal, must be 'hello'")
					// }
					// if (*m)[tmpTarget][1] != "world.txt" {
					// 	t.Error("Add set not equal, must be")
					// }

					// added = m.AddSource(tmpTarget, "foo.txt")
					// if added != true {
					// 	t.Errorf("AddSource %q already exist", tmpTarget)
					// }
				}
			}

			// final check the expected total size of the array
			if len(*m) != tc.expectedTotal {
				t.Errorf("Add total sources not equal, must be %v", tc.expectedTotal)
			}

		})
	}
}

func Test_Sources_GetAllFilepaths(t *testing.T) {
	testCases := []struct {
		src      []string
		expected []string
	}{
		{
			[]string{"fixture/a.txt"},
			[]string{"fixture/a.txt"},
		},
		{
			[]string{"fixture/c/*.txt"},
			[]string{"fixture/c/c1.txt", "fixture/c/c2.txt", "fixture/c/c3.txt"},
		},
		{
			[]string{"fixture/d/**/*"},
			[]string{"fixture/d/d1/d1.json", "fixture/d/d1/d1.txt", "fixture/d/d2/d2.json", "fixture/d/d2/d2.txt"},
		},
		{
			[]string{"fixture/d/**/*.json"},
			[]string{"fixture/d/d1/d1.json", "fixture/d/d2/d2.json"},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q", tc.src), func(t *testing.T) {

			m := Sources{}
			m = make([]string, len(tc.src))
			m = tc.src

			f, err := m.GetAllFilepaths()
			if err != nil {
				t.Error(err)
			}
			if len(f) != len(tc.expected) {
				t.Errorf("GetAllFilepaths lenght not equal, must be %v", len(tc.expected))
			}
			for i := range f {
				if f[i] != tc.expected[i] {
					t.Errorf("GetFilepaths not equal at %v", i)
				}
			}
		})
	}
}

func Test_Sources_GetDirs(t *testing.T) {
	testCases := []struct {
		src      []string
		expected []string
	}{
		{
			src:      []string{"fixture/a.txt"},
			expected: []string{"fixture"},
		},
		{
			src:      []string{"fixture/a.txt", "fixture/b.txt"},
			expected: []string{"fixture"},
		},
		{
			src:      []string{"fixture/d/**/*"},
			expected: []string{"fixture/d/d1", "fixture/d/d2"},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q", tc.src), func(t *testing.T) {

			m := Sources{}
			m = make([]string, len(tc.src))
			m = tc.src

			d, err := m.GetAllDirs()
			if err != nil {
				t.Error("")
			}
			if len(d) != len(tc.expected) {
				t.Errorf("GetAllDirs lenght not equal, must be %v", len(tc.expected))
			}
			for i, v := range d {
				if v != tc.expected[i] {
					t.Errorf("GetAllDirs not equal, must be %v", tc.expected[i])
				}
			}
		})
	}
}

func Test_Sources_ExistSource(t *testing.T) {
	testCases := []struct {
		src      []string
		exist    string
		expected bool
	}{
		{
			src:      []string{"fixture/a.txt"},
			exist:    "fixture/a.txt",
			expected: true,
		},
		{
			src:      []string{"fixture/a.txt"},
			exist:    "fixture/not_exist.txt",
			expected: false,
		},
		{
			src:      []string{"fixture/c/*.txt"},
			exist:    "fixture/c/c1.txt",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q", tc.src), func(t *testing.T) {

			m := Sources{}
			m = make([]string, len(tc.src))
			m = tc.src

			if m.ExistSource(tc.exist) != tc.expected {
				t.Errorf("ExistSource %q not equal, must be %v", tc.exist, tc.expected)
			}
		})
	}
}
