package concatenate

import (
	"io/ioutil"
	"testing"
)

func Test_BytesToBytes(t *testing.T) {
	result := BytesToBytes([]byte("-"), []byte("a1"), []byte("b1"))
	if string(result) != string([]byte("a1-b1")) {
		t.Error("BytesToBytes return not equal, must be 'a1-b1'")
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

func Test_FilesToFile(t *testing.T) {
	tmpTestFile := "tmp_test.txt"

	err := FilesToFile(tmpTestFile, "\n", "fixture/a.txt", "fixture/b.txt")
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
