package concatenate

import "testing"

func Test_Strings(t *testing.T) {
	if Strings(" ", "a", "b") != "a b" {
		t.Error("Strings return not equal, must be 'a b'")
	}
}

func Test_StringsLinebreak(t *testing.T) {
	if StringsLinebreak("a", "b") != "a\nb" {
		t.Error("Strings return not equal, must be 'a\nb'")
	}
}

func Test_Concatenate(t *testing.T) {
	src, err := Concatenate("fixture/a.txt", "fixture/b.txt")
	if err != nil {
		t.Error(err)
	}
	// fmt.Printf("S %q", src)
	if src != "a1\na2\n\n\nb1\nb2\n\n\n" {
		t.Error("not equal")
	}
}

func Test_ConcatenateToFile(t *testing.T) {
	// ConcatendateToFile("test.txt", "a", "b")
}
