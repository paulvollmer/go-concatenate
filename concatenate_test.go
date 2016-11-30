package concatenate

import (
	"io/ioutil"
	"testing"
)

var TestTableBytesToBytes = []struct {
	del    string
	source [][]byte
	result string
}{
	{"", [][]byte{}, ""},
	{"", [][]byte{[]byte("a1")}, "a1"},
	{"-", [][]byte{[]byte("a1"), []byte("b1")}, "a1-b1"},
	{"\n", [][]byte{[]byte("a1"), []byte("b1"), []byte("c1")}, "a1\nb1\nc1"},
	{"\n", [][]byte{[]byte("a1\n"), []byte("b1"), []byte("c1")}, "a1\n\nb1\nc1"},
}

func Test_BytesToBytes(t *testing.T) {
	for _, tt := range TestTableBytesToBytes {
		result := BytesToBytes([]byte(tt.del), tt.source...)
		if string(result) != string([]byte(tt.result)) {
			t.Errorf("BytesToBytes return not equal, must be %q\n", tt.result)
		}
	}
}

func Test_StringsToString(t *testing.T) {
	result := StringsToString("-", "a", "b")
	if result != "a-b" {
		t.Error("StringsToString return not equal, must be 'a-b'")
	}
}

func Test_FilesToBytes(t *testing.T) {
	src, err := FilesToBytes("-", "fixture/a.txt", "fixture/b.txt")
	if err != nil {
		t.Error(err)
	}
	if string(src) != "a1\na2\n-b1\nb2\n" {
		t.Error("FilesToBytes not equal")
	}
}

func Test_FilesToBytes_NotFound(t *testing.T) {
	_, err := FilesToBytes("-", "fixture/not_a.txt", "fixture/not_b.txt")
	if err == nil {
		t.Error("FilesToBytes missing error")
	}
}

func Test_FilesToBytes_Ext(t *testing.T) {
	src, err := FilesToBytes("-", "fixture/*.txt")
	if err != nil {
		t.Error(err)
	}
	if string(src) != "a1\na2\n-b1\nb2\n" {
		t.Error("FilesToBytes not equal")
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
