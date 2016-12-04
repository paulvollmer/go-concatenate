package concatenate

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_BytesToBytes(t *testing.T) {
	testCases := []struct {
		del      string
		sources  [][]byte
		expected string
	}{
		{
			del:      "",
			sources:  [][]byte{},
			expected: "",
		},
		{
			del:      "",
			sources:  [][]byte{[]byte("a1")},
			expected: "a1",
		},
		{
			del:      "-",
			sources:  [][]byte{[]byte("a1"), []byte("b1")},
			expected: "a1-b1",
		},
		{
			del:      "\n",
			sources:  [][]byte{[]byte("a1"), []byte("b1"), []byte("c1")},
			expected: "a1\nb1\nc1",
		},
		{
			del:      "\n",
			sources:  [][]byte{[]byte("a1\n"), []byte("b1"), []byte("c1")},
			expected: "a1\n\nb1\nc1",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := BytesToBytes([]byte(tc.del), tc.sources...)
			if string(result) != string([]byte(tc.expected)) {
				t.Errorf("BytesToBytes return not equal, must be %q\n", tc.expected)
			}
		})
	}
}

func Test_StringsToString(t *testing.T) {
	testCases := []struct {
		del      string
		sources  []string
		expected string
	}{
		{
			del:      "",
			sources:  []string{},
			expected: "",
		},
		{
			del:      "",
			sources:  []string{"a1"},
			expected: "a1",
		},
		{
			del:      "-",
			sources:  []string{"a1", "b1"},
			expected: "a1-b1",
		},
		{
			del:      "\n",
			sources:  []string{"a1", "b1", "c1"},
			expected: "a1\nb1\nc1",
		},
		{
			del:      "\n",
			sources:  []string{"a1\n", "b1", "c1"},
			expected: "a1\n\nb1\nc1",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := StringsToString(tc.del, tc.sources...)
			if result != tc.expected {
				t.Errorf("StringsToString return not equal, must be %q", tc.expected)
			}
		})
	}
}

func Test_FilesToBytes(t *testing.T) {
	testCases := []struct {
		del      string
		sources  []string
		expected string
	}{
		{
			del:      "-",
			sources:  []string{"fixture/a.txt", "fixture/b.txt"},
			expected: "a1\na2\n-b1\nb2\n",
		},
		{
			del:      "-",
			sources:  []string{"fixture/c/*.txt"},
			expected: "c1\n-c2\n-c3\n",
		},
		{
			del:      "-",
			sources:  []string{"fixture/d/**/*.txt"},
			expected: "d1\n-d2\n",
		},
		{
			del:      "-",
			sources:  []string{"fixture/d/**/*"},
			expected: "{}\n-d1\n-{}\n-d2\n",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result, err := FilesToBytes(tc.del, tc.sources...)
			if err != nil {
				t.Error(err)
			}
			if string(result) != tc.expected {
				t.Errorf("FilesToBytes not equal, must be %q", tc.expected)
			}
		})
	}
}

func Test_FilesToBytes_NotFound(t *testing.T) {
	_, err := FilesToBytes("-", "fixture/not_a.txt", "fixture/not_b.txt")
	if err == nil {
		t.Error("FilesToBytes missing error")
	}
}

func Test_FilesToFile(t *testing.T) {
	tmpTestFile := "tmp_test.txt"

	err := FilesToFile(tmpTestFile, 0666, "\n", "fixture/a.txt", "fixture/b.txt")
	if err != nil {
		t.Error(err)
	}

	// check the generated file
	d, err := ioutil.ReadFile(tmpTestFile)
	if err != nil {
		t.Error("FilesToFile written file does not exist")
	}
	if string(d) != "a1\na2\n\nb1\nb2\n" {
		t.Error("FilesToFile not equal")
	}
}
