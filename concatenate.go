package concatenate

import (
	"io/ioutil"
	"strings"
)

func Strings(d string, src ...string) string {
	return strings.Join(src, d)
}

func StringsLinebreak(src ...string) string {
	return strings.Join(src, "\n")
}

func Concatenate(src ...string) (string, error) {

	// build it...
	// log.Printf("==> Files Changed % #v\n", src)
	tmp := make([]byte, 0)
	// tmp = append(tmp, []byte("/*//* GENERATED SOURCECODE*/ \n//* DO NOT EDIT BY HAND */\n\n")...)

	for _, v := range src {
		d, err := ioutil.ReadFile(v)
		if err != nil {
			return "", err
		}
		tmp = append(tmp, d...)
		tmp = append(tmp, []byte("\n\n")...)
	}

	return string(tmp), nil
}

func ConcatenateToFile(file string, src ...string) error {
	// log.Printf("--> Write target %q\n", target)
	con, err := Concatenate(src...)
	if err != nil {
		return err
	}
	// log.Println("write file", file)
	return ioutil.WriteFile(file, []byte(con), 0666)
}
