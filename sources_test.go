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
		src             []string
		expectedIsError bool
		expectedTotal   int
	}{
		{
			src:             []string{""},
			expectedIsError: false,
			expectedTotal:   0,
		},
		{
			src:             []string{"src-1"},
			expectedIsError: true,
			expectedTotal:   1,
		},
		{
			src:             []string{"fixture/a.txt", "fixture/b.txt"},
			expectedIsError: true,
			expectedTotal:   2,
		},
		{
			src:             []string{"fixture/a.txt", "fixture/b.txt", "fixture/a.txt"},
			expectedIsError: false,
			expectedTotal:   2,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q", tc.src), func(t *testing.T) {

			m := NewSources()

			// test to Add() sources and check if the result is equal
			for addI, addV := range tc.src {
				err := m.Add(addV)
				if tc.expectedIsError {
					if err != nil {
						t.Errorf("Add %v missing error", addI)
					}
				}
				// if the total expected size is higher than zero, we can check the array items
				if tc.expectedTotal > 0 {
					get, err := m.Get(addI)
					if err == nil {
						if get != addV {
							t.Errorf("Add source %v not equal, must be %v", addI, addV)
						}
					}

				}
			}

			// final check the expected total size of the array
			mTotal := m.Total()
			if mTotal != tc.expectedTotal {
				t.Errorf("Add total sources (%v) not equal, must be %v", mTotal, tc.expectedTotal)
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
