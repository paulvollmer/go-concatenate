package concatenate

import (
	"io/ioutil"
	"testing"
)

func Test_BytesToBytes(t *testing.T) {
	// fmt.Println(string(BytesToBytes([]byte("-"), []byte("a1"), []byte("b1"))))
	if string(BytesToBytes([]byte("-"), []byte("a1"), []byte("b1"))) != string([]byte("a1-b1")) {
		t.Error("BytesToBytes return not equal, must be 'a1-b1'")
	}
}

func Test_StringsToString(t *testing.T) {
	if StringsToString("-", "a", "b") != "a-b" {
		t.Error("StringsToString return not equal, must be 'a-b'")
	}
}

func Test_FilesToBytes(t *testing.T) {
	src, err := FilesToBytes("-", "fixture/a.txt", "fixture/b.txt")
	if err != nil {
		t.Error(err)
	}
	// fmt.Printf("S %q", src)
	if string(src) != "a1\na2\n-b1\nb2\n" {
		t.Error("FilesToBytes not equal")
	}
}

func Test_FilesToFile(t *testing.T) {
	tmpTestFile := "tmp_test.txt"

	// ConcatendateToFile("test.txt", "a", "b")
	err := FilesToFile(tmpTestFile, "\n", "fixture/a.txt", "fixture/b.txt")
	if err != nil {
		t.Error(err)
	}

	// check the generated file
	d, err := ioutil.ReadFile(tmpTestFile)
	if err != nil {
		t.Error("FilesToFile written file does not exist")
	}
	// fmt.Println(string(d))
	if string(d) != "a1\na2\n\nb1\nb2\n" {
		t.Error("FilesToFile not equal")
	}
}
